<template>
    <div class="text-center py-5 space-y-5">
        <input 
            data-v-step="list_name" 
            v-model="listName" 
            type="text" 
            :placeholder="$t('compose.list.title')" 
            class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all bg-white text-zinc-900 border-none max-w-160 dark:bg-zinc-800 dark:text-white" 
        />
        <textarea 
            data-v-step="list_description" 
            v-model="listDescription" 
            type="text"
            :placeholder="$t('compose.list.description')" 
            class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all bg-white text-zinc-900 border-none max-w-160 dark:bg-zinc-800 dark:text-white" 
        />
    </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
    listName: String,
    listDescription: String
})

const emit = defineEmits(['update:listName', 'update:listDescription'])

const listName = computed({
    get: () => props.listName,
    set: (value) => emit('update:listName', value)
})

const listDescription = computed({
    get: () => props.listDescription,
    set: (value) => emit('update:listDescription', value)
})
</script>
