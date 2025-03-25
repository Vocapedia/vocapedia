<template>
    <transition name="fade" mode="out-in">
        <div v-if="response.chapter">
            <small class="pb-5 flex justify-center">
                {{
                    $t('list_helper', {
                        lang: $t(response.chapter.lang),
                        target_lang: $t(response.chapter.target_lang)
                    })
                }}
            </small>
            <div class="space-y-4 text-center">
                <h1 class="font-bold text-4xl">
                    {{ response.chapter.title }}
                </h1>
                <div>{{ response.chapter.description }}</div>
            </div>

            <div class="max-w-160 w-full mx-auto flex justify-around items-center py-5">
                <button
                    class="flex items-center smooth-click bg-yellow-50 dark:bg-yellow-900 rounded-full px-2 font-bold">
                    <mdicon name="star-outline" class="dark:text-yellow-400 text-yellow-500" size="32" />
                    <span v-splitting>4198</span>
                </button>
                <router-link :to="'/l/' + $route.params.id + '/games'" class="smooth-click">
                    <mdicon name="gamepad-variant-outline" size="32" />
                </router-link>
            </div>

            <div class="max-w-160 w-full mx-auto">
                <div class="space-x-2">
                    <button @click="changeVariant('word-list')"
                        :class="$route.query.variant == 'tutorial' ? 'text-zinc-400 dark:text-zinc-600' : ''"
                        class="cursor-pointer text-lg font-semibold">
                        {{ $t('word-list') }}
                    </button>
                    <button @click="changeVariant('tutorial')"
                        :class="$route.query.variant == 'tutorial' ? '' : 'text-zinc-400 dark:text-zinc-600'"
                        class="cursor-pointer text-lg font-semibold">
                        {{ $t('tutorial') }}
                    </button>
                </div>
                <hr class="border-t-2 border-zinc-200 dark:border-zinc-800 my-4 opacity-50">
            </div>
            <transition name="fade" mode="out-in">
                <component :is="currentComponent" :response="response" />
            </transition>
        </div>
        <div v-else class="loading-spinner mx-auto" />

    </transition>

</template>

<script setup>
import TutorialView from "./TutorialView.vue"
import WordListView from "./WordListView.vue"
import { ref, watch, shallowRef, onMounted } from "vue"
import { useRoute, useRouter } from "vue-router"
const route = useRoute()
const router = useRouter()

import { useFetch } from '@/composable/useFetch';
const response = ref("{}")

onMounted(async () => {
    response.value = await useFetch("/chapters/" + route.params.id)
})


function changeVariant(params) {
    router.push({ query: { ...route.query, variant: params } });
}
const currentComponent = shallowRef(WordListView)
watch(route, (newV) => {
    if (newV.query.variant == 'tutorial') {
        currentComponent.value = TutorialView
    } else {
        currentComponent.value = WordListView
    }
})

onMounted(() => {
    if (route.query.variant == 'tutorial') {
        currentComponent.value = TutorialView
    } else {
        currentComponent.value = WordListView
    }
});
</script>