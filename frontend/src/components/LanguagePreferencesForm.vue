<template>
    <div class="space-y-4">
        <!-- Known Languages -->
        <div>
            <label class="block text-sm font-medium mb-2">{{ $t('language_preferences.known_languages') }}</label>
            <div class="space-y-2">
                <div v-for="(lang, index) in knownLanguages" :key="'known-' + index" class="flex items-center space-x-2">
                    <select v-model="knownLanguages[index]"
                        class="flex-1 px-3 py-2 bg-zinc-100 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-md focus:border-zinc-300 dark:focus:border-zinc-600">
                        <option value="">{{ $t('language_preferences.select_language') }}</option>
                        <option v-for="language in languages" :key="language.code" :value="language.code">
                            {{ language.name }}
                        </option>
                    </select>
                    <button @click="removeKnownLanguage(index)" 
                        class="p-2 text-red-500 hover:text-red-700">
                        <mdicon name="minus-circle" />
                    </button>
                </div>
                <button @click="addKnownLanguage" 
                    class="w-full px-3 py-2 border-2 border-dashed border-zinc-300 dark:border-zinc-600 rounded-md text-zinc-500 dark:text-zinc-400 hover:border-zinc-400 dark:hover:border-zinc-500">
                    + {{ $t('language_preferences.add_known_language') }}
                </button>
            </div>
        </div>

        <!-- Target Languages -->
        <div>
            <label class="block text-sm font-medium mb-2">{{ $t('language_preferences.target_languages') }}</label>
            <div class="space-y-2">
                <div v-for="(lang, index) in targetLanguages" :key="'target-' + index" class="flex items-center space-x-2">
                    <select v-model="targetLanguages[index]"
                        class="flex-1 px-3 py-2 bg-zinc-100 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-md focus:border-zinc-300 dark:focus:border-zinc-600">
                        <option value="">{{ $t('language_preferences.select_language') }}</option>
                        <option v-for="language in languages" :key="language.code" :value="language.code">
                            {{ language.name }}
                        </option>
                    </select>
                    <button @click="removeTargetLanguage(index)" 
                        class="p-2 text-red-500 hover:text-red-700">
                        <mdicon name="minus-circle" />
                    </button>
                </div>
                <button @click="addTargetLanguage" 
                    class="w-full px-3 py-2 border-2 border-dashed border-zinc-300 dark:border-zinc-600 rounded-md text-zinc-500 dark:text-zinc-400 hover:border-zinc-400 dark:hover:border-zinc-500">
                    + {{ $t('language_preferences.add_target_language') }}
                </button>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useI18n } from 'vue-i18n';
import { useFetch } from '@/composable/useFetch';
import { useToast } from '@/composable/useToast';
import languages from '@/utils/language/languages.json';

const { t } = useI18n();
const emit = defineEmits(['update:loading']);

const knownLanguages = ref(['']);
const targetLanguages = ref(['']);
const isLoading = ref(false);
const toast = useToast();

// Load current language preferences
const loadLanguagePreferences = async () => {
    try {
        const response = await useFetch('/user/language-preferences', {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            }
        });
        
        if (response) {
            knownLanguages.value = response.known_languages?.length > 0 ? response.known_languages : [''];
            targetLanguages.value = response.target_languages?.length > 0 ? response.target_languages : [''];
        }
    } catch (error) {
        console.error('Error loading language preferences:', error);
    }
};

// Save language preferences
const saveLanguagePreferences = async () => {
    // Filter out empty values
    const filteredKnown = knownLanguages.value.filter(lang => lang !== '');
    const filteredTarget = targetLanguages.value.filter(lang => lang !== '');

    if (filteredKnown.length === 0 || filteredTarget.length === 0) {
        toast.error(t('language_preferences.validation.select_at_least_one'));
        return false;
    }

    // Check for duplicates within known languages
    const knownDuplicates = filteredKnown.filter((lang, index) => filteredKnown.indexOf(lang) !== index);
    if (knownDuplicates.length > 0) {
        toast.error(t('language_preferences.validation.no_duplicates_known'));
        return false;
    }

    // Check for duplicates within target languages
    const targetDuplicates = filteredTarget.filter((lang, index) => filteredTarget.indexOf(lang) !== index);
    if (targetDuplicates.length > 0) {
        toast.error(t('language_preferences.validation.no_duplicates_target'));
        return false;
    }

    // Check for overlapping languages between known and target
    const hasOverlap = filteredKnown.some(lang => filteredTarget.includes(lang));
    if (hasOverlap) {
        toast.error(t('language_preferences.validation.no_overlap'));
        return false;
    }

    emit('update:loading', true);
    
    try {
        const response = await useFetch('/user/language-preferences', {
            method: 'PUT',
            body: {
                known_languages: filteredKnown,
                target_languages: filteredTarget
            }
        });

        if (response) {
            toast.success(t('language_preferences.success.updated'));
            return true;
        } else {
            toast.error(t('language_preferences.error.update_failed'));
            return false;
        }
    } catch (error) {
        console.error('Error saving language preferences:', error);
        toast.error(t('language_preferences.error.update_failed'));
        return false;
    } finally {
        emit('update:loading', false);
    }
};

// Add/Remove language functions
const addKnownLanguage = () => {
    knownLanguages.value.push('');
};

const removeKnownLanguage = (index) => {
    knownLanguages.value.splice(index, 1);
    if (knownLanguages.value.length === 0) {
        knownLanguages.value.push('');
    }
};

const addTargetLanguage = () => {
    targetLanguages.value.push('');
};

const removeTargetLanguage = (index) => {
    targetLanguages.value.splice(index, 1);
    if (targetLanguages.value.length === 0) {
        targetLanguages.value.push('');
    }
};

onMounted(loadLanguagePreferences);

// Expose save function to parent
defineExpose({
    saveLanguagePreferences
});
</script>
