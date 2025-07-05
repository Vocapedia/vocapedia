<template>
    <div v-auto-animate class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 dark:from-gray-900 dark:to-gray-800">
        <div class="container mx-auto px-4 py-8">
            <!-- Header -->
            <div v-auto-animate class="text-center mb-8">
                <div class="inline-flex items-center justify-center w-16 h-16 bg-gradient-to-r from-blue-500 to-purple-600 rounded-full mb-4">
                    <mdicon name="school" size="32" class="text-white" />
                </div>
                <h1 class="text-3xl font-bold text-gray-900 dark:text-white mb-2">
                    {{ $t('stream.createStream') }}
                </h1>
                <p class="text-gray-600 dark:text-gray-400 max-w-2xl mx-auto">
                    {{ $t('stream.createStreamDescription') }}
                </p>
            </div>

            <!-- Create Stream Form -->
            <div v-auto-animate class="max-w-2xl mx-auto">
                <div class="card bg-white dark:bg-zinc-800 rounded-2xl shadow-xl p-8 space-y-6">
                    <form @submit.prevent="createStream" class="space-y-6">
                        <!-- Stream Title -->
                        <div v-auto-animate class="space-y-2">
                            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">
                                {{ $t('stream.title') }}
                            </label>
                            <input
                                v-model="form.title"
                                type="text"
                                required
                                class="w-full px-4 py-3 rounded-xl border border-gray-200 dark:border-gray-600 bg-gray-50 dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200"
                                :placeholder="$t('stream.titlePlaceholder')"
                            />
                        </div>

                        <!-- Stream Description -->
                        <div v-auto-animate class="space-y-2">
                            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">
                                {{ $t('stream.description') }}
                            </label>
                            <textarea
                                v-model="form.description"
                                rows="4"
                                required
                                class="w-full px-4 py-3 rounded-xl border border-gray-200 dark:border-gray-600 bg-gray-50 dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200"
                                :placeholder="$t('stream.descriptionPlaceholder')"
                            />
                        </div>

                        <!-- Language Selection -->
                        <div v-auto-animate class="grid grid-cols-1 md:grid-cols-2 gap-4">
                            <div class="space-y-2">
                                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">
                                    {{ $t('stream.fromLanguage') }}
                                </label>
                                <select
                                    v-model="form.lang"
                                    required
                                    class="w-full px-4 py-3 rounded-xl border border-gray-200 dark:border-gray-600 bg-gray-50 dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200"
                                >
                                    <option value="">{{ $t('stream.selectLanguage') }}</option>
                                    <option v-for="lang in languages" :key="lang.code" :value="lang.code">
                                        {{ lang.name }}
                                    </option>
                                </select>
                            </div>

                            <div class="space-y-2">
                                <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">
                                    {{ $t('stream.toLanguage') }}
                                </label>
                                <select
                                    v-model="form.target_lang"
                                    required
                                    class="w-full px-4 py-3 rounded-xl border border-gray-200 dark:border-gray-600 bg-gray-50 dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200"
                                >
                                    <option value="">{{ $t('stream.selectLanguage') }}</option>
                                    <option v-for="lang in languages" :key="lang.code" :value="lang.code">
                                        {{ lang.name }}
                                    </option>
                                </select>
                            </div>
                        </div>

                        <!-- Schedule Date & Time -->
                        <div v-auto-animate class="space-y-2">
                            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">
                                {{ $t('stream.scheduledAt') }}
                            </label>
                            <input
                                v-model="form.scheduled_at"
                                type="datetime-local"
                                required
                                :min="minDateTime"
                                class="w-full px-4 py-3 rounded-xl border border-gray-200 dark:border-gray-600 bg-gray-50 dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200"
                            />
                        </div>

                        <!-- Stream Duration -->
                        <div v-auto-animate class="space-y-2">
                            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">
                                {{ $t('stream.duration') }}
                            </label>
                            <select
                                v-model="form.duration"
                                required
                                class="w-full px-4 py-3 rounded-xl border border-gray-200 dark:border-gray-600 bg-gray-50 dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200"
                            >
                                <option value="">{{ $t('stream.selectDuration') }}</option>
                                <option value="30">30 {{ $t('stream.minutes') }}</option>
                                <option value="60">1 {{ $t('stream.hour') }}</option>
                                <option value="90">1.5 {{ $t('stream.hours') }}</option>
                                <option value="120">2 {{ $t('stream.hours') }}</option>
                                <option value="180">3 {{ $t('stream.hours') }}</option>
                            </select>
                        </div>

                        <!-- Max Participants -->
                        <div v-auto-animate class="space-y-2">
                            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">
                                {{ $t('stream.maxParticipants') }}
                            </label>
                            <select
                                v-model="form.max_participants"
                                required
                                class="w-full px-4 py-3 rounded-xl border border-gray-200 dark:border-gray-600 bg-gray-50 dark:bg-gray-700 text-gray-900 dark:text-white focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all duration-200"
                            >
                                <option value="">{{ $t('stream.selectMaxParticipants') }}</option>
                                <option value="5">5 {{ $t('stream.participants') }}</option>
                                <option value="10">10 {{ $t('stream.participants') }}</option>
                                <option value="15">15 {{ $t('stream.participants') }}</option>
                                <option value="20">20 {{ $t('stream.participants') }}</option>
                                <option value="25">25 {{ $t('stream.participants') }}</option>
                            </select>
                        </div>

                        <!-- Action Buttons -->
                        <div v-auto-animate class="flex flex-col sm:flex-row gap-4 pt-4">
                            <button
                                type="button"
                                @click="$router.push('/streamers')"
                                class="flex-1 px-6 py-3 rounded-xl border border-gray-200 dark:border-gray-600 text-gray-700 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-gray-700 transition-all duration-200"
                            >
                                {{ $t('common.cancel') }}
                            </button>
                            <button
                                type="submit"
                                :disabled="loading"
                                class="flex-1 px-6 py-3 rounded-xl bg-gradient-to-r from-blue-500 to-purple-600 hover:from-blue-600 hover:to-purple-700 text-white font-medium transition-all duration-200 transform hover:scale-[1.02] disabled:opacity-50 disabled:cursor-not-allowed"
                            >
                                <div v-if="loading" class="flex items-center justify-center">
                                    <div class="animate-spin w-5 h-5 border-2 border-white border-t-transparent rounded-full mr-2"></div>
                                    {{ $t('common.creating') }}
                                </div>
                                <div v-else class="flex items-center justify-center">
                                    <mdicon name="plus" size="20" class="mr-2" />
                                    {{ $t('stream.createStream') }}
                                </div>
                            </button>
                        </div>
                    </form>
                </div>
            </div>

            <!-- Preview Card -->
            <div v-auto-animate v-if="isFormValid" class="max-w-md mx-auto mt-8">
                <div class="card bg-white dark:bg-zinc-800 rounded-2xl shadow-lg p-6">
                    <h3 class="text-lg font-semibold text-gray-900 dark:text-white mb-4">
                        {{ $t('stream.preview') }}
                    </h3>
                    <div class="space-y-3">
                        <div class="flex items-center justify-between">
                            <div class="text-base font-medium text-gray-900 dark:text-white">
                                {{ form.title }}
                            </div>
                            <div class="flex items-center space-x-1">
                                <div class="w-2 h-2 bg-blue-500 rounded-full"></div>
                                <span class="text-xs text-blue-600 dark:text-blue-400">Scheduled</span>
                            </div>
                        </div>
                        <div class="text-sm text-gray-600 dark:text-gray-400">
                            {{ form.description }}
                        </div>
                        <div class="flex items-center justify-between text-xs text-gray-500 dark:text-gray-400">
                            <div class="flex items-center space-x-2">
                                <span>{{ form.max_participants }} max</span>
                                <span>•</span>
                                <span>{{ getLanguageName(form.lang) }} → {{ getLanguageName(form.target_lang) }}</span>
                            </div>
                            <span>{{ formatDateTime(form.scheduled_at) }}</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { useFetch } from '@/composable/useFetch'
import { useToast } from '@/composable/useToast'
import { getUser } from '@/utils/token'

const router = useRouter()
const { t } = useI18n()
const toast = useToast()

const loading = ref(false)
const form = ref({
    title: '',
    description: '',
    lang: '',
    target_lang: '',
    scheduled_at: '',
    duration: '',
    max_participants: ''
})

const languages = ref([
    { code: 'en', name: 'English' },
    { code: 'tr', name: 'Türkçe' },
    { code: 'es', name: 'Español' },
    { code: 'fr', name: 'Français' },
    { code: 'de', name: 'Deutsch' },
    { code: 'it', name: 'Italiano' },
    { code: 'pt', name: 'Português' },
    { code: 'ru', name: 'Русский' },
    { code: 'ja', name: '日本語' },
    { code: 'ko', name: '한국어' },
    { code: 'zh', name: '中文' },
    { code: 'ar', name: 'العربية' }
])

// Check if user has teacher role
const checkTeacherRole = () => {
    const user = getUser()
    if (!user || user.role !== 'teacher') {
        toast.error(t('stream.teacherRoleRequired'))
        router.push('/streamers')
        return false
    }
    return true
}

// Minimum date time (current time + 30 minutes)
const minDateTime = computed(() => {
    const now = new Date()
    now.setMinutes(now.getMinutes() + 30)
    return now.toISOString().slice(0, 16)
})

// Check if form is valid for preview
const isFormValid = computed(() => {
    return form.value.title && 
           form.value.description && 
           form.value.lang && 
           form.value.target_lang && 
           form.value.scheduled_at && 
           form.value.duration && 
           form.value.max_participants
})

// Get language name by code
const getLanguageName = (code) => {
    const lang = languages.value.find(l => l.code === code)
    return lang ? lang.name : code
}

// Format date time for preview
const formatDateTime = (datetime) => {
    if (!datetime) return ''
    const date = new Date(datetime)
    return date.toLocaleDateString() + ' ' + date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
}

// Create stream
const createStream = async () => {
    if (!checkTeacherRole()) return
    
    loading.value = true
    try {
        const response = await useFetch('/stream/create', {
            method: 'POST',
            body: JSON.stringify(form.value)
        })
        
        toast.success(t('stream.createSuccess'))
        router.push(`/stream/${response.room}`)
    } catch (error) {
        console.error('Error creating stream:', error)
        toast.error(t('stream.createError'))
    } finally {
        loading.value = false
    }
}

// Initialize
onMounted(() => {
    checkTeacherRole()
    // Set default minimum date
    form.value.scheduled_at = minDateTime.value
})
</script>
