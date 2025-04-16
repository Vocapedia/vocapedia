<template>
    <div v-auto-animate class="mb-10 space-y-5 max-w-3xl mx-auto">
        <div class="text-center bg-zinc-800 p-5 rounded-lg border border-zinc-700">
            {{ $t('compose.warn') }}
        </div>
        <div v-if="!isLanguagesSelected" class="flex gap-4 max-w-160 mx-auto p-4">
            <div class="z-1 relative w-full">
                <input v-model="languageQuery"
                    class="bg-zinc-100 dark:bg-zinc-800 w-full p-3 rounded-xl outline-none border border-zinc-200 dark:border-zinc-700 focus:border-zinc-300 dark:focus:border-zinc-600 transition"
                    :placeholder="$t('compose.select.language')" @focus="showLanguageDropdown = true"
                    @blur="hideLanguageDropdown" />
                <ul v-if="showLanguageDropdown"
                    class="absolute bg-white dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 w-full rounded-xl mt-1 max-h-40 overflow-y-auto">
                    <li v-for="lang in filteredLanguages" :key="lang" @click="selectLanguage(lang.name, lang.code)"
                        class="p-2 cursor-pointer hover:bg-zinc-200 dark:hover:bg-zinc-700">
                        {{ lang.name }}
                    </li>
                </ul>
            </div>

            <div class="z-1 relative w-full">
                <input v-model="targetLanguageQuery"
                    class="bg-zinc-100 dark:bg-zinc-800 w-full p-3 rounded-xl outline-none border border-zinc-200 dark:border-zinc-700 focus:border-zinc-300 dark:focus:border-zinc-600 transition"
                    :placeholder="$t('compose.select.target_language')" @focus="showTargetDropdown = true"
                    @blur="hideTargetDropdown" />
                <ul v-if="showTargetDropdown"
                    class="absolute bg-white dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 w-full rounded-xl mt-1 max-h-40 overflow-y-auto">
                    <li v-for="lang in filteredTargetLanguages" :key="lang"
                        @click="selectTargetLanguage(lang.name, lang.code)"
                        class="p-2 cursor-pointer hover:bg-zinc-200 dark:hover:bg-zinc-700">
                        {{ lang.name }}
                    </li>
                </ul>
            </div>
        </div>
        <div v-if="targetLanguageCode && languageCode" class="text-center py-5 space-y-5">
            <input v-model="listName" type="text" :placeholder="$t('compose.list.title')" class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all 
             bg-white text-zinc-900  border-none
             max-w-160 
             dark:bg-zinc-800 dark:text-white " />
            <textarea v-model="listDescription" type="text" :placeholder="$t('compose.list.description')" class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all 
             bg-white text-zinc-900  border-none
             max-w-160 
             dark:bg-zinc-800 dark:text-white " />
        </div>
        <div v-if="targetLanguageCode && languageCode" class="max-w-160 w-full mx-auto">
            <div class="flex items-center justify-between space-x-5">
                <div class="flex items-center space-x-2">
                    <router-link :to="$route.path + '?variant=word-list'"
                        :class="$route.query.variant == 'tutorial' ? 'text-zinc-400 dark:text-zinc-600' : ''"
                        class="truncate cursor-pointer text-lg font-semibold">
                        {{ $t('word-list') }}
                    </router-link>
                    <router-link :to="$route.path + '?variant=tutorial'"
                        :class="$route.query.variant == 'tutorial' ? '' : 'text-zinc-400 dark:text-zinc-600'"
                        class="truncate cursor-pointer text-lg font-semibold">
                        {{ $t('tutorial') }}
                    </router-link>
                </div>
                <button @click="compose" class="smooth-click rounded-full bg-sky-100 dark:bg-sky-700 px-2 py-1">{{
                    $t('save') }}
                </button>
            </div>
            <hr class="border-t-2 border-zinc-200 dark:border-zinc-800 my-4 opacity-50">
        </div>

        <div v-if="targetLanguageCode && languageCode">


            <div v-if="$route.query.variant === '' || $route.query.variant === 'tutorial'">
                <TextEditor v-model="editorContent" />
            </div>
            <div v-else>
                <div v-auto-animate class="space-y-15">
                    <div v-for="(w, wi) in wordBases" :key="wi">
                        <div class="flex items-center justify-between">

                            <div class="w-96 p-2">
                                <h6>{{ $t('type') }}</h6>
                                <div class="pl-5 w-96 ">
                                    <select v-model="w.type"
                                        class="w-full px-4 py-2 bg-zinc-100 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-xl focus:border-zinc-300 dark:focus:border-zinc-600">
                                        <option value="noun">{{ $t('word_types.noun') }}</option>
                                        <option value="verb">{{ $t('word_types.verb') }}</option>
                                        <option value="adjective">{{ $t('word_types.adjective') }}</option>
                                        <option value="adverb">{{ $t('word_types.adverb') }}</option>
                                        <option value="pronoun">{{ $t('word_types.pronoun') }}</option>
                                        <option value="preposition">{{ $t('word_types.preposition') }}</option>
                                        <option value="conjunction">{{ $t('word_types.conjunction') }}</option>
                                        <option value="interjection">{{ $t('word_types.interjection') }}
                                        </option>
                                    </select>
                                </div>
                            </div>
                            <div>
                                <button @click="removeWordBase(wi)"
                                    class="smooth-click p-2 bg-red-100 dark:bg-red-950 hover:bg-red-200 dark:hover:bg-red-700 rounded-full">
                                    <mdicon name="close" />
                                </button>
                            </div>
                        </div>
                        <div v-auto-animate class="flex items-center overflow-x-auto scrollbar-hide space-x-4 p-4">
                            <div v-for="(word, i) in w.words" :key="i" class="flex">
                                <div class="card p-5 w-96 space-y-5">
                                    <div class="flex items-center justify-end">
                                        <button @click="removeWord(i, wi)"
                                            class="smooth-click p-1 bg-red-100 dark:bg-red-950 hover:bg-red-200 dark:hover:bg-red-700 rounded-full">
                                            <mdicon name="close" />
                                        </button>
                                    </div>
                                    <div>
                                        <h6>{{ $t('language') }}</h6>
                                        <div class="pl-5">
                                            <select v-model="word.lang"
                                                class="w-full px-4 py-2 bg-zinc-100 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-xl focus:border-zinc-300 dark:focus:border-zinc-600">
                                                <option v-for="lang in [{
                                                    code: languageCode,
                                                    name: languageQuery.toString()
                                                }, {
                                                    code: targetLanguageCode,
                                                    name: targetLanguageQuery.toString()
                                                }]" :value="lang.code"> {{ lang.name.valueOf() }}</option>
                                            </select>
                                        </div>
                                    </div>
                                    <div>
                                        <h6>{{ $t('compose.input.description') }}</h6>
                                        <div class="pl-5">
                                            <input v-model="word.word" type="text"
                                                :placeholder="$t('compose.input.word')" class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all 
                             bg-white text-zinc-900  border-none
                             max-w-160 
                             dark:bg-zinc-800 dark:text-white " />

                                        </div>
                                    </div>
                                    <div>
                                        <h6>{{ $t('compose.input.description') }}</h6>
                                        <div class="pl-5">
                                            <input v-model="word.description" type="text"
                                                :placeholder="$t('compose.input.description')" class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all 
                             bg-white text-zinc-900  border-none
                             max-w-160 
                             dark:bg-zinc-800 dark:text-white " />
                                        </div>
                                    </div>
                                    <div>
                                        <h6>{{ $t('examples') }}</h6>

                                        <div class="pl-5">
                                            <input v-model="word.examples" type="text"
                                                :placeholder="$t('compose.input.example')" class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all 
                           bg-white text-zinc-900  border-none
                           max-w-160 
                           dark:bg-zinc-800 dark:text-white " />
                                        </div>
                                    </div>
                                </div>

                            </div>
                            <div>
                                <button @click="addWordToWordBase(wi)" class="smooth-click p-5">
                                    <mdicon name="translate" />
                                </button>
                            </div>
                        </div>
                    </div>
                    <button @click="addWordBase" class="smooth-click2 py-3 bg-sky-100 dark:bg-sky-700 w-full">
                        {{ $t('compose.add_new_wordbase') }}
                    </button>
                </div>
            </div>
        </div>
    </div>

</template>
<script setup>
import { ref, reactive, computed, watch, onMounted } from "vue"
import languages from "@/utils/language/languages.json"
import TextEditor from "@/components/TextEditor.vue"
import { useFetch } from "@/composable/useFetch"
import { useRouter, useRoute } from "vue-router"
import { getLangByCode } from "@/utils/language/languages"
const route = useRoute()
const router = useRouter()
const listName = ref("")
const listDescription = ref("")
const wordBases = reactive([])
const editorContent = ref("<p></p>");
const languageQuery = ref('');
const targetLanguageQuery = ref('');
const languageCode = ref('');
const targetLanguageCode = ref('');
const showLanguageDropdown = ref(false);
const showTargetDropdown = ref(false);
onMounted(async () => {
    await useFetch("/user/check")
    if (route.params.id) {
        await useFetch("/public/chapters/" + route.params.id).then(r => {
            var chapter = r.chapter
            listName.value = chapter.title
            listDescription.value = chapter.description
            editorContent.value = chapter.tutorial
            wordBases.push(...chapter.word_bases)
            languageCode.value = chapter.lang
            languageQuery.value = getLangByCode(chapter.lang).name
            targetLanguageCode.value = chapter.target_lang
            targetLanguageQuery.value = getLangByCode(chapter.target_lang).name
        })
    }
})
async function compose() {
    var method = "POST"
    if (route.params.id) {
        method = "PUT"
    }
    await useFetch("/chapters/compose", {
        method: method,
        body: {
            chapter_id: route.params.id,
            title: listName.value,
            description: listDescription.value,
            lang: languageCode.value,
            target_lang: targetLanguageCode.value,
            tutorial: editorContent.value,
            word_bases: wordBases
        }
    }).then(r => {
        router.replace('/l/' + r.chapter_id)
    })
}
function addWordBase() {
    wordBases.push({
        type: "noun",
        words: [{
            lang: languageCode.value,
        }, {
            lang: targetLanguageCode.value,
        }]
    })
}

function addWordToWordBase(i) {
    wordBases[i].words.push({ lang: targetLanguageCode.value })
}

function removeWordBase(i) {
    if (wordBases.length <= 1) {
        return
    }
    wordBases.splice(i, 1)
}
function removeWord(i, wi) {
    if (wordBases[wi].words.length <= 2) {
        return
    }
    wordBases[wi].words.splice(i, 1)
}

const isLanguagesSelected = ref(false)

watch(
    () => [languageCode.value, targetLanguageCode.value], ([newLang, newTargetLang], [oldLang, oldTargetLang]) => {
        if (newLang && newTargetLang) {
            isLanguagesSelected.value = true
        }
    }
);

watch(isLanguagesSelected, (n, o) => {
    if (isLanguagesSelected.value && wordBases.length == 0) {
        addWordBase()
    }
})


const filteredLanguages = computed(() => {
    var result = languages
    result = result.filter(lang => lang.name.toLowerCase().includes(languageQuery.value.toLowerCase()))
    if (targetLanguageCode.value != '') {
        result = result.filter(lang => !lang.code.toLowerCase().includes(targetLanguageCode.value.toLowerCase()))
    }
    return result
});

const filteredTargetLanguages = computed(() => {
    var result = languages
    result = result.filter(lang => lang.name.toLowerCase().includes(targetLanguageQuery.value.toLowerCase()))
    if (languageCode.value != '') {
        result = result.filter(lang => !lang.code.toLowerCase().includes(languageCode.value.toLowerCase()))
    }
    return result
});

const selectLanguage = (lang, langCode) => {
    languageQuery.value = lang;
    languageCode.value = langCode
    showLanguageDropdown.value = false;
};

const selectTargetLanguage = (lang, langCode) => {
    targetLanguageQuery.value = lang;
    targetLanguageCode.value = langCode
    showTargetDropdown.value = false;
};

const hideLanguageDropdown = () => {
    setTimeout(() => {
        showLanguageDropdown.value = false;
    }, 200);
};

const hideTargetDropdown = () => {
    setTimeout(() => {
        showTargetDropdown.value = false;
    }, 200);
};
</script>