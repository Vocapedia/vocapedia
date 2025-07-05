<template>
    <div class="space-y-4">
        <div class="bg-yellow-50 dark:bg-yellow-900/20 border border-yellow-200 dark:border-yellow-700 rounded-lg p-4">
            <div class="flex items-center space-x-2">
                <mdicon name="information-outline" class="text-yellow-600 dark:text-yellow-400" />
                <span class="text-sm text-yellow-700 dark:text-yellow-300">
                    {{ $t('teacher_request.info') }}
                </span>
            </div>
        </div>

        <p class="text-sm text-zinc-600 dark:text-zinc-400">
            {{ $t('teacher_request.description') }}
        </p>

        <div class="space-y-2">
            <label for="teacher-description" class="block text-sm font-medium text-zinc-700 dark:text-zinc-300">
                {{ $t('teacher_request.description_label') }}
            </label>
            <textarea 
                id="teacher-description"
                v-model="teacherDescription"
                :placeholder="$t('teacher_request.description_placeholder')"
                rows="4"
                class="w-full px-3 py-2 border border-zinc-300 dark:border-zinc-600 rounded-lg bg-white dark:bg-zinc-800 text-zinc-900 dark:text-zinc-100 placeholder-zinc-500 dark:placeholder-zinc-400 focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-400 focus:border-transparent"
            ></textarea>
            <p class="text-xs text-zinc-500 dark:text-zinc-400">
                {{ $t('teacher_request.description_help') }}
            </p>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';
import { useI18n } from 'vue-i18n';
import { useFetch } from '@/composable/useFetch';
import { useToast } from '@/composable/useToast';

const { t } = useI18n();
const emit = defineEmits(['update:loading']);
const toast = useToast();
const teacherDescription = ref('');

// Request teacher status
const requestTeacherStatus = async () => {
    if (!teacherDescription.value.trim()) {
        toast.error(t('teacher_request.validation.description_required'));
        return;
    }

    emit('update:loading', true);
    
    try {
        const response = await useFetch('/user/request-teacher', {
            method: 'POST',
            body: {
                description: teacherDescription.value
            }
        });
        
        if (response) {
            toast.success(t('teacher_request.success.submitted'));
            teacherDescription.value = '';
        }
    } catch (error) {
        console.error('Error requesting teacher status:', error);
        toast.error(t('teacher_request.error.submit_failed'));
    } finally {
        emit('update:loading', false);
    }
};

// Expose request function to parent
defineExpose({
    requestTeacherStatus
});
</script>
