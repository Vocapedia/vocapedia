<template>
    <div class="space-y-2">
        <div v-for="l in response" :key="l.id">
            <router-link :to="'/l/' + BigInt(l.id)">
                <div class="card smooth-click2 duration-200">
                    <h1 class="text-xl font-semibold p-5">
                        {{ l.title }}
                    </h1>
                    <h4 class="px-5">
                        {{ l.description }}
                    </h4>
                    <div class="flex justify-evenly pt-5">
                        <button class="bg-yellow-50 dark:bg-yellow-900 rounded-full px-2 font-bold">
                            <span class="flex items-center space-x-1 dark:text-yellow-400 text-yellow-500">
                                <mdicon v-if="l.is_favorited" name="star" class="dark:text-yellow-400 text-yellow-500"
                                    size="27" />
                                <mdicon v-else name="star-outline" class="dark:text-yellow-400 text-yellow-500"
                                    size="27" />
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
                    <small class="flex justify-between px-2 py-1">
                        <span class="flex items-center space-x-2">
                            <span>
                                {{ getLangByCode(l.lang).name }}
                            </span>
                            <mdicon size="16" name="arrow-right" />
                            <span>
                                {{ getLangByCode(l.target_lang).name }}
                            </span>
                        </span>
                        <span>
                            {{ l.creator.username }}
                        </span>
                    </small>
                </div>
            </router-link>
        </div>
        <InfiniteLoading @infinite="load">
            <template #complete>
                <div class="text-center">.</div>
            </template>
            <template #error="{ retry }">
                <div class="text-center">
                    <button class="smooth-click bg-red-200 dark:bg-red-700 rounded-full p-2" @click="retry">
                        <mdicon name="reload" />
                    </button>
                </div>
            </template>
            <template #spinner>
                <div v-motion-fade class="flex justify-center p-2">
                    <mdicon class="text-center" spin name="loading"/>
                </div>
            </template>

        </InfiniteLoading>
    </div>
</template>

<script setup>
import { useFetch } from '@/composable/useFetch';
import { getLangByCode } from '@/utils/language/languages';
import { ref } from 'vue';
import InfiniteLoading from "v3-infinite-loading";
import "v3-infinite-loading/lib/style.css";

const props = defineProps({
    uri: {
        type: String,
        default: ""
    },
});

const page = ref(1);
const response = ref([]);

const load = async $state => {
    try {
        const res = await fetchData(page.value)
        response.value.push(...res);
        if (res.length < 10) {
            $state.complete()
        }
        else {
            $state.loaded();
        }
        page.value++;
    } catch (error) {
        $state.error();
    }
};

const fetchData = async (page) => {
    const urlWithPage = addQueryParamToUri(props.uri, 'page', page);
    const result = await useFetch(urlWithPage);
    return result.list;
};

const addQueryParamToUri = (uri, key, value) => {
    const separator = uri.includes('?') ? '&' : '?';
    return `${uri}${separator}${key}=${value}`;
};
</script>
