<template>
    <div v-auto-animate class="space-y-15">
        <div v-for="(w, wi) in wordBases" :key="wi">
            <div class="flex items-center justify-between">
                <div class="w-96 p-2">
                    <h6>{{ $t('type') }}</h6>
                    <div data-v-step="word_type" class="pl-5 w-96">
                        <select v-model="w.type"
                            class="w-full px-4 py-2 bg-zinc-100 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-xl focus:border-zinc-300 dark:focus:border-zinc-600">
                            <option value="noun">{{ $t('word_types.noun') }}</option>
                            <option value="verb">{{ $t('word_types.verb') }}</option>
                            <option value="adjective">{{ $t('word_types.adjective') }}</option>
                            <option value="adverb">{{ $t('word_types.adverb') }}</option>
                            <option value="pronoun">{{ $t('word_types.pronoun') }}</option>
                            <option value="preposition">{{ $t('word_types.preposition') }}</option>
                            <option value="conjunction">{{ $t('word_types.conjunction') }}</option>
                            <option value="interjection">{{ $t('word_types.interjection') }}</option>
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
                    <WordCard 
                        :word="word" 
                        :wordIndex="i"
                        :languageCode="languageCode"
                        :targetLanguageCode="targetLanguageCode"
                        :languageQuery="languageQuery"
                        :targetLanguageQuery="targetLanguageQuery"
                        @remove="removeWord(i, wi)"
                    />
                </div>
                <div>
                    <button data-v-step="new_word_button" @click="addWordToWordBase(wi)"
                        class="smooth-click p-5">
                        <mdicon name="translate" />
                    </button>
                </div>
            </div>
        </div>
        <button data-v-step="new_word_base" @click="addWordBase"
            class="smooth-click2 py-3 bg-sky-100 dark:bg-sky-700 w-full">
            {{ $t('compose.add_new_wordbase') }}
        </button>
    </div>
</template>

<script setup>
import WordCard from './WordCard.vue'

const props = defineProps({
    wordBases: Array,
    languageCode: String,
    targetLanguageCode: String,
    languageQuery: String,
    targetLanguageQuery: String
})

const emit = defineEmits(['update:wordBases'])

const addWordBase = () => {
    const newWordBases = [...props.wordBases]
    newWordBases.push({
        type: "noun",
        words: [{
            lang: props.languageCode,
        }, {
            lang: props.targetLanguageCode,
        }]
    })
    emit('update:wordBases', newWordBases)
}

const addWordToWordBase = (i) => {
    const newWordBases = [...props.wordBases]
    newWordBases[i].words.push({ lang: props.targetLanguageCode })
    emit('update:wordBases', newWordBases)
}

const removeWordBase = (i) => {
    if (props.wordBases.length <= 1) {
        return
    }
    const newWordBases = [...props.wordBases]
    newWordBases.splice(i, 1)
    emit('update:wordBases', newWordBases)
}

const removeWord = (i, wi) => {
    if (props.wordBases[wi].words.length <= 2) {
        return
    }
    const newWordBases = [...props.wordBases]
    newWordBases[wi].words.splice(i, 1)
    emit('update:wordBases', newWordBases)
}
</script>
