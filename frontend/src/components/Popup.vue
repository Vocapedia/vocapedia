<template>
    <Teleport to="body">
        <transition name="slide-fade" mode="out-in">
            <div v-if="modelValue"
                class="fixed inset-0 flex items-center justify-center bg-black/50 backdrop-blur-sm px-5 z-1"
                @click="closeModal">
                <div class="space-y-5 bg-white dark:bg-zinc-800 p-5 rounded-lg shadow-xl w-full max-w-md" @click.stop>
                    <slot name="header" />
                    <div class="max-h-96 overflow-y-auto overflow-x-hidden scrollbar-hide">
                        <slot name="description" />
                    </div>
                    <slot name="buttons" />
                </div>
            </div>
        </transition>
    </Teleport>
</template>

<script setup>
import { defineProps, defineEmits } from 'vue';
import { useEventListener } from '@vueuse/core';

const props = defineProps({
    modelValue: Boolean,
});

const emit = defineEmits(['update:modelValue']);

const closeOnEscape = (event) => {
    if (event.key === 'Escape') emit('update:modelValue', false);
};
useEventListener('keydown', closeOnEscape);
const closeModal = () => emit('update:modelValue', false);
</script>