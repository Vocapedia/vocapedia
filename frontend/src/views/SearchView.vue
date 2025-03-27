<template>
    <div class="max-w-160 mx-auto">

        <p class="py-5">
            {{ $t('search_results', { query: route.query.q, count: (response.list ?? []).length, }) }}
        </p>

        <transition name="fade" mode="out-in">
            <WordLists v-if="response.list" :response="response" />
            <div v-else class="loading-spinner mx-auto" />
        </transition>
    </div>
</template>

<script setup>

import { useRoute } from 'vue-router';
import { watch, ref, onMounted } from "vue"
import WordLists from '@/components/WordLists.vue';
const response = ref("{}")
const route = useRoute()
import { useFetch } from "@/composable/useFetch";


onMounted(async () => {
    response.value = await useFetch("/public/chapters/search?q=" + route.query.q)
})

watch(route, async (newQuery, oldQuery) => {
    response.value = await useFetch("/public/chapters/search?q=" + newQuery.query.q)
})

</script>