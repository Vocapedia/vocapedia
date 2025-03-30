<template>
    <div v-auto-animate class="mb-10 space-y-5">
        <div class="flex py-5 justify-center">
            <input v-model="listName" type="text" :placeholder="$t('list_name')" class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all 
             bg-white text-zinc-900  border-none
             max-w-160 
             dark:bg-zinc-800 dark:text-white " />
        </div>
        <div class="flex gap-4 max-w-160 mx-auto p-4">
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
                        <div class="w-96 p-2">
                            <h6>{{ $t('type') }}</h6>
                            <div class="pl-5 w-96 ">
                                <select
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
                        <div v-auto-animate class="flex items-center overflow-x-auto scrollbar-hide space-x-4 p-4 ">
                            <div v-for="(word, i) in w.words" :key="i" class="flex">
                                <div class="card p-5 w-96 space-y-5">
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
                                        <h6>{{ $t('header') }}</h6>
                                        <div class="pl-5">
                                            <input v-model="word.word" type="text" :placeholder="$t('list_name')" class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all 
                             bg-white text-zinc-900  border-none
                             max-w-160 
                             dark:bg-zinc-800 dark:text-white " />

                                        </div>
                                    </div>
                                    <div>
                                        <h6>{{ $t('description') }}</h6>
                                        <div class="pl-5">
                                            <input v-model="word.description" type="text" :placeholder="$t('list_name')"
                                                class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all 
                             bg-white text-zinc-900  border-none
                             max-w-160 
                             dark:bg-zinc-800 dark:text-white " />
                                        </div>
                                    </div>
                                    <div>
                                        <h6>{{ $t('examples') }}</h6>

                                        <div class="pl-5">
                                            <input v-model="word.examples" type="text" :placeholder="$t('example')"
                                                class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all 
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
import { ref, reactive, computed, watch } from "vue"
import languages from "@/utils/language/languages.json"
import TextEditor from "@/components/TextEditor.vue"
const listName = ref()
const wordBases = reactive([])
const isFirstWordBaseAdded = ref(false)
const editorContent = ref("<p></p>");

function addWordBase() {
    wordBases.push({
        type: "",
        words: [{
            lang: languageCode.value,
        }, {
            lang: targetLanguageCode.value,
        }]
    })
}
function addWordToWordBase(i) {
    wordBases[i].words.push({})
}
const languageQuery = ref('');
const targetLanguageQuery = ref('');
const languageCode = ref('');
const targetLanguageCode = ref('');
const showLanguageDropdown = ref(false);
const showTargetDropdown = ref(false);

watch(
    () => [languageCode.value, targetLanguageCode.value], ([newLang, newTargetLang], [oldLang, oldTargetLang]) => {
        if (newLang && newTargetLang) {
            isFirstWordBaseAdded.value = true
        }
    }
);
watch(isFirstWordBaseAdded, (n, o) => {
    if (isFirstWordBaseAdded.value) {
        addWordBase()

    }
})


const filteredLanguages = computed(() => {
    return languages.filter(lang => lang.name.toLowerCase().includes(languageQuery.value.toLowerCase()));
});

const filteredTargetLanguages = computed(() => {
    return languages.filter(lang => lang.name.toLowerCase().includes(targetLanguageQuery.value.toLowerCase()));
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