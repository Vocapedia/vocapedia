<template>
    <div v-if="isVisible" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
        <div class="bg-white dark:bg-zinc-900 rounded-lg p-6 max-w-md w-full mx-4">
            <h2 class="text-xl font-bold mb-4">{{ $t('teacher_request.title') }}</h2>
            
            <div class="space-y-4">
                <p class="text-sm text-zinc-600 dark:text-zinc-400">
                    {{ $t('teacher_request.description') }}
                </p>
                
                <div class="bg-yellow-50 dark:bg-yellow-900/20 border border-yellow-200 dark:border-yellow-700 rounded-lg p-4">
                    <div class="flex items-center space-x-2">
                        <mdicon name="information-outline" class="text-yellow-600 dark:text-yellow-400" />
                        <span class="text-sm text-yellow-700 dark:text-yellow-300">
                            {{ $t('teacher_request.info') }}
                        </span>
                    </div>
                </div>
            </div>

            <div class="flex justify-end space-x-3 mt-6">
                <button @click="closeModal" 
                    class="smooth-click px-4 py-2 text-zinc-600 dark:text-zinc-400 hover:text-zinc-800 dark:hover:text-zinc-200">
                    {{ $t('common.cancel') }}
                </button>
                <button @click="requestTeacherStatus" 
                    class="smooth-click px-4 py-2 bg-emerald-500 text-white rounded-md hover:bg-emerald-600 disabled:opacity-50"
                    :disabled="isLoading">
                    {{ isLoading ? $t('common.processing') : $t('teacher_request.request') }}
                </button>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';
import { useFetch } from '@/composable/useFetch';
import { useToast } from '@/composable/useToast';

const props = defineProps({
    isVisible: {
        type: Boolean,
        default: false
    }
});

const emit = defineEmits(['close']);

const isLoading = ref(false);
const { showToast } = useToast();

// Request teacher status
const requestTeacherStatus = async () => {
    isLoading.value = true;
    
    try {
        const response = await useFetch('/user/request-teacher', {
            method: 'POST',
        });

        if (response.status === 200) {
            showToast('Teacher request submitted successfully! We will review your request.', 'success');
            closeModal();
        } else {
            showToast('Error submitting teacher request', 'error');
        }
    } catch (error) {
        console.error('Error requesting teacher status:', error);
        showToast('Error submitting teacher request', 'error');
    } finally {
        isLoading.value = false;
    }
};

const closeModal = () => {
    emit('close');
};
</script>
