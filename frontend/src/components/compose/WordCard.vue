<template>
    <div class="card p-5 w-96 space-y-5">
        <div class="flex items-center justify-end">
            <button @click="$emit('remove')"
                class="smooth-click p-1 bg-red-100 dark:bg-red-950 hover:bg-red-200 dark:hover:bg-red-700 rounded-full">
                <mdicon name="close" />
            </button>
        </div>
        <div>
            <h6>{{ $t('language') }}</h6>
            <div class="pl-5">
                <select
                    :data-v-step="wordIndex === 0 ? 'word_lang1' : (wordIndex === 1 ? 'word_lang2' : '')"
                    v-model="wordLang"
                    class="w-full px-4 py-2 bg-zinc-100 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-xl focus:border-zinc-300 dark:focus:border-zinc-600">
                    <option v-for="lang in availableLanguages" :key="lang.code" :value="lang.code">
                        {{ lang.name }}
                    </option>
                </select>
            </div>
        </div>
        <div>
            <h6>{{ $t('compose.input.word') }}</h6>
            <div class="pl-5">
                <input v-model="wordText" type="text"
                    :placeholder="$t('compose.input.word')" 
                    class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all bg-white text-zinc-900 border-none max-w-160 dark:bg-zinc-800 dark:text-white" />
            </div>
        </div>
        <div>
            <h6>{{ $t('compose.input.description') }}</h6>
            <div class="pl-5">
                <input v-model="wordDescription" type="text"
                    :placeholder="$t('compose.input.description')" 
                    class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all bg-white text-zinc-900 border-none max-w-160 dark:bg-zinc-800 dark:text-white" />
            </div>
        </div>
        <div>
            <h6>{{ $t('examples') }}</h6>
            <div class="pl-5">
                <input v-model="wordExamples" type="text"
                    :placeholder="$t('compose.input.example')" 
                    class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all bg-white text-zinc-900 border-none max-w-160 dark:bg-zinc-800 dark:text-white" />
            </div>
        </div>
    </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
    word: Object,
    wordIndex: Number,
    languageCode: String,
    targetLanguageCode: String,
    languageQuery: String,
    targetLanguageQuery: String
})

const emit = defineEmits(['remove'])

const availableLanguages = computed(() => [
    {
        code: props.languageCode,
        name: props.languageQuery
    },
    {
        code: props.targetLanguageCode,
        name: props.targetLanguageQuery
    }
])

const wordLang = computed({
    get: () => props.word.lang,
    set: (value) => {
        props.word.lang = value
    }
})

const wordText = computed({
    get: () => props.word.word,
    set: (value) => {
        props.word.word = value
    }
})

const wordDescription = computed({
    get: () => props.word.description,
    set: (value) => {
        props.word.description = value
    }
})

const wordExamples = computed({
    get: () => props.word.examples,
    set: (value) => {
        props.word.examples = value
    }
})
</script>
