<template>
    <div class="space-y-4" v-auto-animate>
        <div v-if="isLoading" class="text-center py-8">
            <div class="dots">
                <div class="dot"></div>
                <div class="dot"></div>
                <div class="dot"></div>
            </div>
            <p class="mt-4 text-sm text-zinc-600 dark:text-zinc-400">{{ $t('shared.loading') }}</p>
        </div>

        <div v-else-if="discardedChapters.length === 0" class="flex flex-col py-12">
            <mdicon name="archive-off-outline" size="64" class="text-zinc-400 mx-auto mb-6" />
            <h3 class="text-lg font-medium text-zinc-800 dark:text-zinc-200 mb-2">{{
                $t('account.popup.discarded_chapters.no_chapters.title') }}</h3>
            <p class="text-sm text-zinc-600 dark:text-zinc-400">{{
                $t('account.popup.discarded_chapters.no_chapters.description') }}
            </p>
        </div>

        <div v-else class="space-y-4"  v-auto-animate>
            <div  v-for="chapter in discardedChapters" :key="chapter.id"
                class="bg-white dark:bg-zinc-900 rounded-xl p-6 border border-zinc-200 dark:border-zinc-700 shadow-sm hover:shadow-md transition-shadow">
                <div class="flex flex-col sm:flex-row sm:items-start sm:justify-between gap-4">
                    <div class="flex-1">
                        <h3 class="font-semibold text-xl text-zinc-900 dark:text-zinc-100 mb-2">{{ chapter.title }}</h3>
                        <p class="text-zinc-600 dark:text-zinc-400 mb-4 leading-relaxed">
                            {{ chapter.description }}
                        </p>
                        <div class="flex flex-wrap gap-4 text-sm text-zinc-500 dark:text-zinc-400">
                            <div class="flex items-center gap-1">
                                <mdicon name="calendar-remove" size="16" />
                                <span>{{ $t('account.popup.discarded_chapters.deleted_at') }}: {{
                                    formatDate(chapter.deleted_at)
                                    }}</span>
                            </div>
                            <div class="flex items-center gap-1">
                                <mdicon name="format-list-bulleted" size="16" />
                                <span>{{ $t('account.popup.discarded_chapters.words_count') }}: {{ chapter.word_count ||
                                    0 }}</span>
                            </div>
                        </div>
                    </div>
                    <div class="flex gap-3 sm:flex-col sm:w-auto">
                        <PopupButton icon="restore" theme="blue" @click="restoreChapter(BigInt(chapter.id))" />
                        <PopupButton icon="delete-forever" theme="red" @click="showDeleteConfirm(chapter)" />
                    </div>
                </div>
            </div>
        </div>
    </div>

    <!-- Delete Confirmation Modal -->
    <Popup v-model="showDeleteModal">
        <template #header>
            <div class="flex items-center gap-3">
                <mdicon name="alert-circle" size="24" class="text-rose-500" />
                <h3 class="text-xl font-semibold text-zinc-900 dark:text-zinc-100">{{
                    $t('discarded_chapters.confirm_delete_title') }}</h3>
            </div>
        </template>
        <template #description>
            <div class="space-y-4">
                <p class="text-zinc-600 dark:text-zinc-400">
                    {{ $t('discarded_chapters.confirm_delete_description') }}
                </p>
                <div class="bg-rose-50 dark:bg-rose-900/20 border border-rose-200 dark:border-rose-800 rounded-lg p-4">
                    <p class="text-sm text-rose-700 dark:text-rose-300 font-medium">
                        {{ $t('discarded_chapters.confirm_delete_warning') }}
                    </p>
                </div>
                <div v-if="chapterToDelete" class="bg-zinc-50 dark:bg-zinc-800 rounded-lg p-4">
                    <h4 class="font-medium text-zinc-900 dark:text-zinc-100 mb-1">{{ chapterToDelete.title }}</h4>
                    <p class="text-sm text-zinc-600 dark:text-zinc-400">{{ chapterToDelete.description }}</p>
                </div>
            </div>
        </template>
        <template #buttons>
            <div class="flex gap-3 justify-end">
                <PopupButton icon="delete-forever" theme="red" @click="showDeleteConfirm(chapter)" />

                <button @click="showDeleteModal = false"
                    class="smooth-click px-4 py-2 text-zinc-700 dark:text-zinc-300 hover:bg-zinc-100 dark:hover:bg-zinc-800 rounded-lg transition-colors">
                    {{ $t('common.cancel') }}
                </button>
                <button @click="confirmPermanentDelete"
                    class="smooth-click px-4 py-2 bg-rose-500 hover:bg-rose-600 text-white rounded-lg font-medium transition-colors disabled:opacity-50"
                    :disabled="actionLoading === chapterToDelete?.id">
                    {{ actionLoading === chapterToDelete?.id ? $t('shared.loading') :
                        $t('discarded_chapters.delete_permanently') }}
                </button>
            </div>
        </template>
    </Popup>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useI18n } from 'vue-i18n';
import { useFetch } from '@/composable/useFetch';
import { useToast } from '@/composable/useToast';
import Popup from '@/components/shared/Popup.vue';
import PopupButton from './shared/PopupButton.vue';

const { t } = useI18n();
const emit = defineEmits(['update:itemCount']);
const discardedChapters = ref([]);
const isLoading = ref(false);
const actionLoading = ref(null);
const showDeleteModal = ref(false);
const chapterToDelete = ref(null);
const toast = useToast();

// Load discarded chapters
const loadDiscardedChapters = async () => {
    isLoading.value = true;

    try {
        const response = await useFetch('/chapters/discarded', {
            method: 'GET',
        });

        if (response) {
            discardedChapters.value = response.chapters || [];
            emit('update:itemCount', discardedChapters.value.length);
        }
    } catch (error) {
        console.error('Error loading discarded chapters:', error);
        toast.error(t('discarded_chapters.error_loading'));
    } finally {
        isLoading.value = false;
    }
};

// Restore chapter (set deleted_at to null)
const restoreChapter = async (chapterId) => {
    actionLoading.value = chapterId;

    try {
        const response = await useFetch(`/chapters/restore/${chapterId}`, {
            method: 'POST',
        });
        if (response) {
            toast.success(response.message);
            discardedChapters.value = discardedChapters.value.filter(chapter => BigInt(chapter.id) !== chapterId);
            emit('update:itemCount', discardedChapters.value.length);
        }
    } catch (error) {
        toast.error(t(response.error));
    } finally {
        actionLoading.value = null;
    }
};

// Show delete confirmation modal
const showDeleteConfirm = (chapter) => {
    chapterToDelete.value = chapter;
    showDeleteModal.value = true;
};

// Permanently delete chapter
const confirmPermanentDelete = async () => {
    if (!chapterToDelete.value) return;

    const chapterId = BigInt(chapterToDelete.value.id);
    actionLoading.value = chapterId;

    try {
        const response = await useFetch(`/chapters/delete/${chapterId}`, {
            method: 'DELETE',
        });

        if (response) {
            toast.success(t('discarded_chapters.delete_success'));
            // Remove from discarded list
            discardedChapters.value = discardedChapters.value.filter(chapter => chapter.id !== chapterId);
            emit('update:itemCount', discardedChapters.value.length);
            showDeleteModal.value = false;
            chapterToDelete.value = null;
        }
    } catch (error) {
        console.error('Error deleting chapter:', error);
        toast.error(t('discarded_chapters.delete_error'));
    } finally {
        actionLoading.value = null;
    }
};

const formatDate = (dateString) => {
    if (!dateString) return '';
    const date = new Date(dateString);
    return date.toLocaleDateString() + ' ' + date.toLocaleTimeString();
};

onMounted(loadDiscardedChapters);
</script>

<style scoped>
.dots {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 0.5rem;
}

.dot {
    width: 0.5rem;
    height: 0.5rem;
    border-radius: 50%;
    background-color: #3b82f6;
    animation: dot-bounce 1.4s infinite ease-in-out;
}

.dot:nth-child(1) {
    animation-delay: -0.32s;
}

.dot:nth-child(2) {
    animation-delay: -0.16s;
}

@keyframes dot-bounce {

    0%,
    80%,
    100% {
        transform: scale(0);
    }

    40% {
        transform: scale(1);
    }
}
</style>
