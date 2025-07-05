<template>
    <div v-if="isVisible" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
        <div class="bg-white dark:bg-zinc-900 rounded-lg p-6 max-w-md w-full mx-4">
            <h2 class="text-xl font-bold mb-4">{{ $t('language_preferences.title') }}</h2>
            
            <div class="space-y-4">
                <!-- Native Language -->
                <div>
                    <label class="block text-sm font-medium mb-2">{{ $t('language_preferences.native_language') }}</label>
                    <select v-model="selectedNativeLanguage"
                        class="w-full px-3 py-2 bg-zinc-100 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-md focus:border-zinc-300 dark:focus:border-zinc-600">
                        <option value="">{{ $t('language_preferences.select_language') }}</option>
                        <option v-for="lang in languages" :key="lang.code" :value="lang.code">
                            {{ lang.name }}
                        </option>
                    </select>
                </div>

                <!-- Target Language -->
                <div>
                    <label class="block text-sm font-medium mb-2">{{ $t('language_preferences.target_language') }}</label>
                    <select v-model="selectedTargetLanguage"
                        class="w-full px-3 py-2 bg-zinc-100 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-md focus:border-zinc-300 dark:focus:border-zinc-600">
                        <option value="">{{ $t('language_preferences.select_language') }}</option>
                        <option v-for="lang in languages" :key="lang.code" :value="lang.code">
                            {{ lang.name }}
                        </option>
                    </select>
                </div>
            </div>

            <div class="flex justify-end space-x-3 mt-6">
                <button @click="closeModal" 
                    class="smooth-click px-4 py-2 text-zinc-600 dark:text-zinc-400 hover:text-zinc-800 dark:hover:text-zinc-200">
                    {{ $t('common.cancel') }}
                </button>
                <button @click="saveLanguagePreferences" 
                    class="smooth-click px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 disabled:opacity-50"
                    :disabled="isLoading">
                    {{ isLoading ? $t('common.saving') : $t('common.save') }}
                </button>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, watch, onMounted } from 'vue';
import { useFetch } from '@/composable/useFetch';
import { useToast } from '@/composable/useToast';
import languages from '@/utils/language/languages.json';

const props = defineProps({
    isVisible: {
        type: Boolean,
        default: false
    }
});

const emit = defineEmits(['close']);

const selectedNativeLanguage = ref('');
const selectedTargetLanguage = ref('');
const isLoading = ref(false);

const { showToast } = useToast();

// Load current language preferences
const loadLanguagePreferences = async () => {
    try {
        const response = await useFetch('/user/language-preferences', {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${localStorage.getItem('token')}`
            }
        });
        
        if (response.data) {
            selectedNativeLanguage.value = response.data.lang || '';
            selectedTargetLanguage.value = response.data.target_lang || '';
        }
    } catch (error) {
        console.error('Error loading language preferences:', error);
    }
};

// Save language preferences
const saveLanguagePreferences = async () => {
    if (!selectedNativeLanguage.value || !selectedTargetLanguage.value) {
        showToast('Please select both languages', 'error');
        return;
    }

    if (selectedNativeLanguage.value === selectedTargetLanguage.value) {
        showToast('Native and target languages cannot be the same', 'error');
        return;
    }

    isLoading.value = true;
    
    try {
        const response = await useFetch('/user/language-preferences', {
            method: 'PUT',
            headers: {
                'Authorization': `Bearer ${localStorage.getItem('token')}`,
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                lang: selectedNativeLanguage.value,
                target_lang: selectedTargetLanguage.value
            })
        });

        if (response.status === 200) {
            showToast('Language preferences updated successfully', 'success');
            closeModal();
        } else {
            showToast('Error updating language preferences', 'error');
        }
    } catch (error) {
        console.error('Error saving language preferences:', error);
        showToast('Error updating language preferences', 'error');
    } finally {
        isLoading.value = false;
    }
};

const closeModal = () => {
    emit('close');
};

// Load preferences when modal becomes visible
watch(() => props.isVisible, (visible) => {
    if (visible) {
        loadLanguagePreferences();
    }
});
</script>
