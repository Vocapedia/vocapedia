<template>
    <div class="max-w-4xl mx-auto px-4">
        <span class="text-2xl font-bold">
            {{ $t('games.header') }}
        </span>
        <hr class="border-t-2 border-zinc-200 dark:border-zinc-800 my-4 opacity-50">
        <div class="grid grid-cols-12 gap-4">
            <router-link v-show="wordCount > g.wordCount" v-for="g in games" :key="g.name"
                :to="'/l/' + $route.params.id + '/game' + g.location"
                class="card flex justify-center items-center col-span-12 sm:col-span-6 lg:col-span-4 rounded-xl smooth-click h-48">
                <span class="text-xl font-bold p-4"> {{ $t(g.name) }}</span>
            </router-link>
        </div>
    </div>
</template>
<script setup>
import { useFetch } from "@/composable/useFetch"
import { ref, onMounted } from "vue"
import { useRoute } from "vue-router"
const route = useRoute()
const wordCount = ref(0)
onMounted(async () => {
    await useFetch("/public/chapters/" + route.params.id).then(r => {
        wordCount.value = r.chapter.word_count
    })

})
const games = ref([
    {
        location: "/test",
        name: "test",
        wordCount: 4
    },
    {
        location: "/flip-word",
        name: "flip-word",
        wordCount: 99
    },
    {
        location: "/quick-pick",
        name: "quick-pick",
        wordCount: 99
    },
    {
        location: "/word-match",
        name: "word-match",
        wordCount: 99
    },
    {
        location: "/word-rush",
        name: "word-rush",
        wordCount: 99
    },
])
</script>