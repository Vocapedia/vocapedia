<template>
    <div>
        <div class="flex py-5 justify-center">
            <input v-model="listName" type="text" :placeholder="$t('list_name')" class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all 
             bg-white text-zinc-900  border-none
             max-w-160 
             dark:bg-zinc-800 dark:text-white " />
        </div>
        <div v-auto-animate class="space-y-2">
            <div class="flex justify-center">
                <div v-auto-animate class="flex items-center overflow-x-auto space-x-4 p-4 scrollbar-hide">
                    <div v-for="i in 2" :key="i">
                        <div class="card p-5 w-96 space-y-3">
                            <div class="flex justify-around items-center">
                                <div>
                                    <h6>{{ $t('language') }}</h6>
                                    <div class="pl-5">
                                        <select
                                            class="block w-full px-4 py-2 text-gray-700 bg-white border border-gray-300 rounded-lg shadow-sm focus:outline-none ">
                                            <option value="en">{{ $t('english') }}</option>
                                            <option value="tr">{{ $t('turkish') }}</option>
                                        </select>
                                    </div>
                                </div>
                                <div>
                                    <h6>{{ $t('type') }}</h6>
                                    <div class="pl-5">
                                        <select
                                            class="block w-full px-4 py-2 text-gray-700 bg-white border border-gray-300 rounded-lg shadow-sm focus:outline-none ">
                                            <option value="noun">{{ $t('noun') }}</option>
                                            <option value="verb">{{ $t('verb') }}</option>
                                            <option value="adjective">{{ $t('adjective') }}</option>
                                            <option value="adverb">{{ $t('adverb') }}</option>
                                            <option value="pronoun">{{ $t('pronoun') }}</option>
                                            <option value="preposition">{{ $t('preposition') }}</option>
                                            <option value="conjunction">{{ $t('conjunction') }}</option>
                                            <option value="interjection">{{ $t('interjection') }}</option>
                                        </select>
                                    </div>
                                </div>
                            </div>
                            <div>
                                <h6>{{ $t('header') }}</h6>
                                <div class="pl-5">
                                    <input v-model="inputLists[0]" type="text" :placeholder="$t('list_name')" class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all 
                             bg-white text-zinc-900  border-none
                             max-w-160 
                             dark:bg-zinc-800 dark:text-white " />

                                </div>
                            </div>
                            <div>
                                <h6>{{ $t('description') }}</h6>
                                <div class="pl-5">
                                    <input v-model="inputLists[0]" type="text" :placeholder="$t('list_name')" class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all 
                             bg-white text-zinc-900  border-none
                             max-w-160 
                             dark:bg-zinc-800 dark:text-white " />
                                </div>
                            </div>
                            <div>
                                <h6>{{ $t('examples') }}</h6>

                                <div class="pl-5">
                                    <input v-model="listName" type="text" :placeholder="$t('example')" class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all 
                           bg-white text-zinc-900  border-none
                           max-w-160 
                           dark:bg-zinc-800 dark:text-white " />
                                </div>
                            </div>
                        </div>

                    </div>
                    <div>
                        <button @click="addLanguage(0)" class="smooth-click p-5">
                            <mdicon name="translate" />
                        </button>
                    </div>
                </div>
            </div>

            <div v-for="i in inputLists.length">
                <hr class="border-t-2 border-zinc-200 dark:border-zinc-800 my-4 opacity-50">
                <input v-model="inputLists[i]" type="text" :placeholder="$t('list_name')" class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all 
             bg-white text-zinc-900  border-none
             max-w-160 
             dark:bg-zinc-800 dark:text-white " />
            </div>
        </div>
    </div>
</template>
<script setup>
import { ref, reactive } from "vue"
const listName = ref()
const inputLists = reactive([]);
function addLanguage(i) {
    inputLists[i].languages.push({})
}
</script>