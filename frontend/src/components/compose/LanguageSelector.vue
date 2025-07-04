<template>
    <div class="flex gap-4 max-w-160 mx-auto p-4">
        <div class="z-1 relative w-full">
            <input 
                data-v-step="main_lang" 
                v-model="languageQuery"
                class="bg-zinc-100 dark:bg-zinc-800 w-full p-3 rounded-xl outline-none border border-zinc-200 dark:border-zinc-700 focus:border-zinc-300 dark:focus:border-zinc-600 transition"
                :placeholder="$t('compose.select.language')" 
                @focus="showLanguageDropdown = true"
                @blur="hideLanguageDropdown" 
            />
            <ul v-if="showLanguageDropdown"
                class="absolute bg-white dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 w-full rounded-xl mt-1 max-h-40 overflow-y-auto z-10">
                <li v-for="lang in filteredLanguages" :key="lang.code" 
                    @click="selectLanguage(lang.name, lang.code)"
                    class="p-2 cursor-pointer hover:bg-zinc-200 dark:hover:bg-zinc-700">
                    {{ lang.name }}
                </li>
            </ul>
        </div>

        <div class="z-1 relative w-full">
            <input 
                data-v-step="target_lang" 
                v-model="targetLanguageQuery"
                class="bg-zinc-100 dark:bg-zinc-800 w-full p-3 rounded-xl outline-none border border-zinc-200 dark:border-zinc-700 focus:border-zinc-300 dark:focus:border-zinc-600 transition"
                :placeholder="$t('compose.select.target_language')" 
                @focus="showTargetDropdown = true"
                @blur="hideTargetDropdown" 
            />
            <ul v-if="showTargetDropdown"
                class="absolute bg-white dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 w-full rounded-xl mt-1 max-h-40 overflow-y-auto z-10">
                <li v-for="lang in filteredTargetLanguages" :key="lang.code"
                    @click="selectTargetLanguage(lang.name, lang.code)"
                    class="p-2 cursor-pointer hover:bg-zinc-200 dark:hover:bg-zinc-700">
                    {{ lang.name }}
                </li>
            </ul>
        </div>
    </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import languages from "@/utils/language/languages.json"

const props = defineProps({
    languageCode: String,
    targetLanguageCode: String,
    languageQuery: String,
    targetLanguageQuery: String
})

const emit = defineEmits(['update:languageCode', 'update:targetLanguageCode', 'update:languageQuery', 'update:targetLanguageQuery'])

const showLanguageDropdown = ref(false)
const showTargetDropdown = ref(false)

const languageQuery = computed({
    get: () => props.languageQuery,
    set: (value) => emit('update:languageQuery', value)
})

const targetLanguageQuery = computed({
    get: () => props.targetLanguageQuery,
    set: (value) => emit('update:targetLanguageQuery', value)
})

const filteredLanguages = computed(() => {
    let result = languages
    result = result.filter(lang => lang.name.toLowerCase().includes(languageQuery.value.toLowerCase()))
    if (props.targetLanguageCode) {
        result = result.filter(lang => lang.code !== props.targetLanguageCode)
    }
    return result
})

const filteredTargetLanguages = computed(() => {
    let result = languages
    result = result.filter(lang => lang.name.toLowerCase().includes(targetLanguageQuery.value.toLowerCase()))
    if (props.languageCode) {
        result = result.filter(lang => lang.code !== props.languageCode)
    }
    return result
})

const selectLanguage = (lang, langCode) => {
    emit('update:languageQuery', lang)
    emit('update:languageCode', langCode)
    showLanguageDropdown.value = false
}

const selectTargetLanguage = (lang, langCode) => {
    emit('update:targetLanguageQuery', lang)
    emit('update:targetLanguageCode', langCode)
    showTargetDropdown.value = false
}

const hideLanguageDropdown = () => {
    setTimeout(() => {
        showLanguageDropdown.value = false
    }, 200)
}

const hideTargetDropdown = () => {
    setTimeout(() => {
        showTargetDropdown.value = false
    }, 200)
}
</script>
