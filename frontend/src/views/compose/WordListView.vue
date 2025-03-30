<template>
    <div v-auto-animate class="space-y-15" v-if="targetLanguageCode && languageCode">
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
                                        name: languageQuery
                                    }, {
                                        code: targetLanguageCode,
                                        name: targetLanguageQuery
                                    }]" :value="lang.code"> {{ lang.name }}</option>
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
                                <input v-model="word.description" type="text" :placeholder="$t('list_name')" class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all 
                             bg-white text-zinc-900  border-none
                             max-w-160 
                             dark:bg-zinc-800 dark:text-white " />
                            </div>
                        </div>
                        <div>
                            <h6>{{ $t('examples') }}</h6>

                            <div class="pl-5">
                                <input v-model="word.examples" type="text" :placeholder="$t('example')" class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all 
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
</template>

<script setup>

</script>