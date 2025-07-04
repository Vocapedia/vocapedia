<template>
    <div class="mb-10 space-y-5 max-w-3xl mx-auto">
        <div class="text-center bg-white dark:bg-zinc-800 p-5 rounded-lg border dark:border-zinc-700 border-zinc-200">
            {{ $t('compose.warn') }}
        </div>

        <!-- Dil Seçimi -->
        <LanguageSelector v-if="!isLanguagesSelected" v-model:languageCode="languageCode"
            v-model:targetLanguageCode="targetLanguageCode" v-model:languageQuery="languageQuery"
            v-model:targetLanguageQuery="targetLanguageQuery" />

        <!-- Liste Bilgileri -->
        <ListInfo v-if="targetLanguageCode && languageCode" v-model:listName="listName"
            v-model:listDescription="listDescription" />

        <!-- İçerik Türü Seçimi ve Kaydet Butonu -->
        <div v-if="targetLanguageCode && languageCode" class="max-w-160 w-full mx-auto">
            <div class="flex items-center justify-between space-x-5">
                <div class="flex items-center space-x-2">
                    <component data-v-step="word_list" :is="'span'"
                        @click="router.replace({ path: $route.path, query: { ...$route.query, variant: 'word-list' } })"
                        :class="($route.query.variant === 'tutorial') ? 'text-zinc-400 dark:text-zinc-600' : ''"
                        class="truncate cursor-pointer text-lg font-semibold">
                        {{ $t('word-list') }}
                    </component>
                    <component data-v-step="tutorial" :is="'span'"
                        @click="router.replace({ path: $route.path, query: { ...$route.query, variant: 'tutorial' } })"
                        :class="($route.query.variant === 'tutorial') ? '' : 'text-zinc-400 dark:text-zinc-600'"
                        class="truncate cursor-pointer text-lg font-semibold">
                        {{ $t('tutorial') }}
                    </component>
                </div>
                <button data-v-step="save_button" @click="compose" v-auto-animate :disabled="isSaving"
                    class="smooth-click rounded-full bg-sky-100 dark:bg-sky-700 px-2 py-1">
                    <span v-if="isSaving">
                        <mdicon spin name="loading" />
                    </span>
                    <span v-else>
                        {{ $t('save') }}
                    </span>
                </button>
            </div>
            <hr class="border-t-2 border-zinc-200 dark:border-zinc-800 my-4 opacity-50">
        </div>

        <!-- İçerik -->
        <div v-if="targetLanguageCode && languageCode">
            <div v-if="$route.query.variant === '' || $route.query.variant === 'tutorial'">
                <TextEditor v-model="editorContent" />
            </div>
            <div v-else>
                <WordBaseList v-model:wordBases="wordBases" :languageCode="languageCode"
                    :targetLanguageCode="targetLanguageCode" :languageQuery="languageQuery"
                    :targetLanguageQuery="targetLanguageQuery" />
            </div>
        </div>
        <div class="tour-button fixed bottom-5 right-5 flex items-center space-x-2"
            v-if="!isLanguagesSelected && isLanguageTourStarted">
            <span class="hint">Reload Tour</span>
            <button @click="restartTours('language_tour')"
                class="smooth-click bg-sky-100 dark:bg-sky-700 p-3 rounded-full shadow-lg">
                <mdicon name="reload" />
            </button>
        </div>

        <div class="tour-button fixed bottom-5 right-5 flex items-center space-x-2"
            v-else-if="targetLanguageCode && languageCode && isWordBaseTourStarted">
            <span class="hint">Reload Tour</span>
            <button @click="restartTours('word_base_tour')"
                class="smooth-click bg-sky-100 dark:bg-sky-700 p-3 rounded-full shadow-lg">
                <mdicon name="reload" />
            </button>
        </div>


        <!-- Tours -->
        <v-tour name="language_tour" :steps="stepsLangs" :isOpen="isTourOpen" />
        <v-tour name="word_base_tour" :steps="stepsWordBase" :isOpen="isTourOpen" />
    </div>
</template>
<script setup>
import { ref, watch, onMounted } from "vue"
import TextEditor from "@/components/TextEditor.vue"
import LanguageSelector from "@/components/compose/LanguageSelector.vue"
import ListInfo from "@/components/compose/ListInfo.vue"
import WordBaseList from "@/components/compose/WordBaseList.vue"
import { useFetch } from "@/composable/useFetch"
import { useTourManager } from "@/composable/useComposeTour"
import { useRouter, useRoute } from "vue-router"
import { getLangByCode } from "@/utils/language/languages"
import { useStorage } from "@vueuse/core"

const route = useRoute()
const router = useRouter()
const isLanguageTourStarted = useStorage('language_tour')
const isWordBaseTourStarted = useStorage('word_base_tour')

const { isTourOpen, stepsLangs, stepsWordBase, startTour, restartTours } = useTourManager()


const listName = ref("")
const listDescription = ref("")
const wordBases = ref([])
const editorContent = ref("<p></p>")
const languageQuery = ref('')
const targetLanguageQuery = ref('')
const languageCode = ref('')
const targetLanguageCode = ref('')
const isSaving = ref(false)
const isLanguagesSelected = ref(false)


watch(
    () => [languageCode.value, targetLanguageCode.value],
    ([newLang, newTargetLang]) => {
        if (newLang && newTargetLang) {
            isLanguagesSelected.value = true
        }
    }
)

watch(isLanguagesSelected, (newValue) => {
    if (newValue && wordBases.value.length === 0) {
        addWordBase()
        startTour("word_base_tour")
    }
})


async function compose() {
    isSaving.value = true
    const method = route.params.id ? "PUT" : "POST"

    try {
        const result = await useFetch("/chapters/compose", {
            method: method,
            body: {
                chapter_id: route.params.id,
                title: listName.value,
                description: listDescription.value,
                lang: languageCode.value,
                target_lang: targetLanguageCode.value,
                tutorial: editorContent.value,
                word_bases: wordBases.value
            }
        })
        router.replace('/l/' + result.chapter_id)
    } catch (error) {
        console.error('Error saving chapter:', error)
    } finally {
        isSaving.value = false
    }
}

function addWordBase() {
    wordBases.value.push({
        type: "noun",
        words: [{
            lang: languageCode.value,
        }, {
            lang: targetLanguageCode.value,
        }]
    })
}


onMounted(async () => {
    await useFetch("/user/check")

    if (route.params.id) {
        try {
            const response = await useFetch("/public/chapters/" + route.params.id)
            const chapter = response.chapter

            listName.value = chapter.title
            listDescription.value = chapter.description
            editorContent.value = chapter.tutorial
            wordBases.value.push(
                ...chapter.word_bases.map(wb => ({
                    ...wb,
                    words: wb.words.map(word => ({
                        ...word,
                        examples: word.examples && word.examples.length > 0 ? word.examples[0] : ""
                    }))
                }))
            )

            languageCode.value = chapter.lang
            languageQuery.value = getLangByCode(chapter.lang).name
            targetLanguageCode.value = chapter.target_lang
            targetLanguageQuery.value = getLangByCode(chapter.target_lang).name
        } catch (error) {
            console.error('Error loading chapter:', error)
        }
    }

    startTour("language_tour")
})
</script>

<style scoped>
.hint {
    opacity: 0;
    visibility: hidden;
    height: 0;
    overflow: hidden;
    transition: opacity 0.3s ease, visibility 0s linear 0.3s, height 0.3s ease;
}

.tour-button:hover .hint {
    opacity: 1;
    visibility: visible;
    height: auto;
    transition: opacity 0.3s ease, visibility 0s linear 0s, height 0.3s ease;
}
</style>