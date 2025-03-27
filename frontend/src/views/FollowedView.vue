<template>
    <div class="max-w-160 mx-auto">
        <p class="py-5">
            {{ $t('followed_results', { count: (response.list ?? []).length, }) }}
        </p>

        <transition name="fade" mode="out-in">
            <WordLists v-if="response.list" :response="response" />
            <div v-else class="loading-spinner mx-auto" />
        </transition>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import WordLists from '@/components/WordLists.vue';
import { useFetch } from '@/composable/useFetch';
const response = ref("{}")

onMounted(async () => {
    response.value = await useFetch("/chapters/favorite")
})
</script>