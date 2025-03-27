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
                <button @click="response.chapter.is_favorited ? deleteFavChapter() : favoriteChapter()"
                    class="space-x-2 flex items-center smooth-click bg-yellow-50 dark:bg-yellow-900 rounded-full py-px pl-2 pr-4 font-bold">
                    <mdicon v-if="response.chapter.is_favorited" name="star"
                        class="dark:text-yellow-400 text-yellow-500 " size="32" />
                    <mdicon v-else name="star-outline" class="dark:text-yellow-400 text-yellow-500 " size="32" />
                    <span>{{ response.chapter.fav_count }}</span>
                </button>
                <router-link :to="'/l/' + $route.params.id + '/games'" class="smooth-click">
                    <mdicon name="gamepad-variant-outline" size="32" />
                </router-link>
            </div>

            <div class="max-w-160 w-full mx-auto">
                <div class="flex items-center justify-between">
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
                    <router-link :to="'/' + response.chapter.creator.username">
                        {{ response.chapter.creator.username }}
                    </router-link>
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
import { useToast } from "@/composable/useToast"

const response = ref({})
const toast = useToast()

onMounted(async () => {
    response.value = await useFetch("/public/chapters/" + route.params.id)
})

async function favoriteChapter() {
    await useFetch("/chapters/favorite?chapter_id=" + route.params.id, {
        method: "POST",
    }).then(r => {
        response.value.chapter.is_favorited = 1
        response.value.chapter.fav_count += 1
        toast.show(r.message)
    })
        .catch(e => {
            toast.show(e.error)
        })
}
async function deleteFavChapter() {
    await useFetch("/chapters/favorite?chapter_id=" + route.params.id, {
        method: "DELETE",
    }).then(r => {
        response.value.chapter.is_favorited = 0
        response.value.chapter.fav_count -= 1
        toast.show(r.message)
    })
        .catch(e => {
            toast.show(e.error)
        })
}

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