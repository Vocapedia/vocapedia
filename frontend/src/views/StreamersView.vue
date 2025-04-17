<template>
    <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 p-4">
        <router-link v-for="(s, i) in streamers" :key="i" :to="'/stream/' + s.id" class="group">
            <div
                class="card smooth-click2 bg-white dark:bg-zinc-800 rounded-xl transition-all duration-300 ease-in-out transform group-hover:scale-105">
                <div class="space-y-2 w-full">
                    <div class="text-xl px-3 pt-1 font-semibold text-gray-900 dark:text-white">{{ s.name }}</div>
                    <div class="text-sm px-3 h-18 text-gray-500 dark:text-gray-400 overflow-hidden text-ellipsis">
                        {{ s.description }}
                    </div>
                    <div
                        class="rounded-b-xl bg-sky-100 dark:bg-sky-700 text-sm flex items-center justify-center space-x-2 p-1">
                        <span class="text-gray-700 dark:text-gray-200">{{ getLangByCode(s.lang).name }}</span>
                        <mdicon name="arrow-right" size="16" class="text-gray-600 dark:text-gray-400" />
                        <span class="text-gray-700 dark:text-gray-200">{{ getLangByCode(s.target_lang).name }}</span>
                    </div>
                </div>
            </div>
        </router-link>
    </div>
</template>



<script setup>
import { getLangByCode } from "@/utils/language/languages";
import { ref } from "vue"
const streamers = ref([
    {
        id: "982347293847293847",
        name: "John",
        lang: "en",
        target_lang: "tr",
        description: "Hi! Fantastic!"
    },
    {
        id: "849238472384723984",
        name: "Akif",
        lang: "tr",
        target_lang: "en",
        description: "Hi! Kokakolastik!"
    },
    {
        id: "728349823749823749",
        name: "Marie",
        lang: "fr",
        target_lang: "en",
        description: "Bonjour! incroyable"
    },
    {
        id: "638472938472938472",
        name: "Carlos",
        lang: "es",
        target_lang: "en",
        description: "¡Hola! increíble"
    },
    {
        id: "517283947239847239",
        name: "Stefan",
        lang: "de",
        target_lang: "en",
        description: "Hallo! fantastisch"
    },
    {
        id: "405938472938472938",
        name: "Luca",
        lang: "it",
        target_lang: "en",
        description: "Ciao! fantastico"
    },
    {
        id: "394857239847239847",
        name: "Mia",
        lang: "pl",
        target_lang: "en",
        description: "Cześć! niesamowite"
    }
])


</script>