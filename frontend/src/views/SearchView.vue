<template>
    <div class="max-w-160 mx-auto">
        <p class="py-5">
            {{ $t('search_results', { query: route.query.q, count: response.list.length, }) }}
        </p>
        <div v-for="l in response.list">
            <router-link :to="'/l/' + l.id">
                <div class="card smooth-click">
                    <h1 class="text-xl font-semibold p-5">
                        {{ l.title }}
                    </h1>
                    <h4 class="px-5">
                        {{ l.description }}
                    </h4>
                    <div class="flex justify-evenly pt-5">
                        <button class="bg-yellow-50 dark:bg-yellow-900 rounded-full px-2 font-bold">
                            <span class="flex items-center space-x-1 dark:text-yellow-400 text-yellow-500">
                                <mdicon name="star-outline" class="dark:text-yellow-400 text-yellow-500" size="27" />
                                <span>{{ l.fav_count }}</span>
                            </span>
                        </button>
                        <button class="bg-blue-50 dark:bg-blue-900 rounded-full px-2 font-bold">
                            <span class="flex items-center space-x-1 dark:text-blue-400 text-blue-500">
                                <mdicon name="format-letter-case" class="dark:text-blue-400 text-blue-500" size="27" />
                                <span>{{ l.word_count }}</span>
                            </span>
                        </button>
                    </div>
                    <small class="flex justify-end px-2 py-1">
                        <span>
                            {{ l.creator.username }}
                        </span>
                    </small>
                </div>
            </router-link>
        </div>
    </div>
</template>

<script setup>
import { useRoute } from 'vue-router';
import { watch } from "vue"
import fake_response from "@/fake/search_list.json";
const response = fake_response
const route = useRoute()

watch(route, async (newQuery, oldQuery) => {
    //console.log(newQuery.query.q)
})


</script>