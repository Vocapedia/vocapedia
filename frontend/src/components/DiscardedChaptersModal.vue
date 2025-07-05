<template>
    <div v-if="isVisible" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
        <div class="bg-white dark:bg-zinc-900 rounded-lg p-6 max-w-2xl w-full mx-4 max-h-[80vh] overflow-y-auto">
            <h2 class="text-xl font-bold mb-4">{{ $t('discarded_chapters.title') }}</h2>
            
            <DiscardedChaptersForm @update:itemCount="handleItemCountUpdate" />

            <div class="flex justify-end space-x-3 mt-6">
                <button @click="closeModal" 
                    class="smooth-click px-4 py-2 text-zinc-600 dark:text-zinc-400 hover:text-zinc-800 dark:hover:text-zinc-200">
                    {{ $t('common.close') }}
                </button>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, watch } from 'vue';
import DiscardedChaptersForm from '@/components/DiscardedChaptersForm.vue';

const props = defineProps({
    isVisible: {
        type: Boolean,
        default: false
    }
});

const emit = defineEmits(['close', 'update:itemCount']);

const handleItemCountUpdate = (count) => {
    emit('update:itemCount', count);
    // Close modal if no items left
    if (count === 0) {
        setTimeout(() => {
            closeModal();
        }, 1000);
    }
};

const closeModal = () => {
    emit('close');
};
</script>
