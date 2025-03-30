<template>
    <div class="max-w-160 mx-auto">

        <div class="p-6 transition-all">
            <div class="text-center">
                <small>{{ $t('note_description') }}</small>
            </div>
            <div class="flex mt-4 gap-2">
                <input v-model="newNote" @keyup.enter="addNote" :placeholder="$t('add_new_note')" class="w-full p-3 border shadow-sm outline-none transition-all 
             bg-white text-zinc-900 border-none rounded-full
             dark:bg-zinc-800 dark:text-white" />
                <button @click="addNote"
                    class="smooth-click rounded-full bg-blue-500 text-white px-3  hover:bg-blue-600 transition-all">
                    <mdicon name="plus" />
                </button>
            </div>
            <transition name="slide-fade" mode="out-in">
                <draggable v-if="notes && notes.length > 0" v-model="notes" tag="div" item-key="id"
                    class="mt-6 space-y-3" v-auto-animate>
                    <template #item="{ element }" :key="element.id">
                        <div
                            class="p-4 bg-white dark:bg-zinc-800 rounded-lg shadow-md flex justify-between items-center smooth-click2 transition-transform cursor-move">
                            <div>
                                <p class="text-zinc-800 dark:text-white text-lg">
                                    {{ element.content }}
                                </p>
                                <small class="text-zinc-500 dark:text-zinc-400">
                                    {{ element.date }}
                                </small>
                            </div>
                            <button @click="removeNote(element.id)"
                                class="smooth-click text-red-500 hover:text-red-700 transition-all">
                                <mdicon name="close" />
                            </button>
                        </div>
                    </template>
                </draggable>
                <div v-else class="text-center py-10">{{ $t('empty') }}</div>
            </transition>
        </div>

    </div>
</template>

<script setup>
import { ref } from "vue";
import { useStorage } from "@vueuse/core";
import draggable from 'vuedraggable';

const notes = useStorage("notes", []);
const newNote = ref("");

const addNote = () => {
    if (!newNote.value.trim()) return;
    notes.value.unshift({
        id: Date.now(),
        content: newNote.value,
        date: new Date().toLocaleString(),
    });
    newNote.value = "";

};


const removeNote = (id) => {
    notes.value = notes.value.filter((note) => note.id !== id);
};
</script>