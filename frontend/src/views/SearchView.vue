<template>
    <div class="max-w-160 mx-auto">

        <p class="py-5">
            <router-link :to="'/search?q=' + $route.query.q">
                {{
                    $t('search_results', { query: decodeURIComponent($route.query.q) })
                }}
            </router-link>

        </p>

        <WordLists :key="$route.query.q" :uri="'/public/chapters/search?q=' + $route.query.q" />
    </div>
</template>

<script setup>
import WordLists from '@/components/WordLists.vue';
</script>