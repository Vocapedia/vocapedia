<template>
    <div v-auto-animate>
        <div class="container mx-auto px-4 py-8">
            <!-- Header -->
            <div v-auto-animate class="text-center mb-12">
                <div
                    class="inline-flex items-center justify-center w-20 h-20 bg-gradient-to-br from-purple-500 via-blue-500 to-indigo-600 rounded-3xl mb-6 shadow-lg">
                    <mdicon name="video-account" size="40" class="text-white" />
                </div>
                <h1
                    class="text-4xl font-bold bg-gradient-to-r from-purple-600 to-blue-600 bg-clip-text text-transparent mb-4">
                    {{ $t('stream.createStream') }}
                </h1>
                <p class="text-gray-600 dark:text-gray-400 max-w-2xl mx-auto text-lg">
                    {{ $t('stream.createStreamDescription') }}
                </p>
            </div>

            <!-- Create Stream Form -->
            <div v-auto-animate class="max-w-3xl mx-auto">
                <div
                    class="bg-white/80 dark:bg-zinc-800/80 backdrop-blur-xl rounded-3xl card p-8 border border-white/20 dark:border-gray-700/50 space-y-8">
                    <form @submit.prevent="createStream" class="space-y-8">
                        <!-- Stream Title -->
                        <div v-auto-animate class="space-y-3">
                            <label class="block text-sm font-semibold text-gray-800 dark:text-gray-200">
                                <mdicon name="format-title" size="18" class="inline mr-2 text-purple-500" />
                                {{ $t('stream.title') }}
                            </label>
                            <input v-model="form.title" type="text" required
                                class="w-full px-5 py-4 rounded-2xl border-2 border-gray-200 dark:border-gray-600 bg-white/50 dark:bg-gray-700/50 text-gray-900 dark:text-white focus:ring-4 focus:ring-purple-500/20 focus:border-purple-500 transition-all duration-300 placeholder-gray-400"
                                :placeholder="$t('stream.titlePlaceholder')" />
                        </div>

                        <!-- Stream Description -->
                        <div v-auto-animate class="space-y-3">
                            <label class="block text-sm font-semibold text-gray-800 dark:text-gray-200">
                                <mdicon name="text-long" size="18" class="inline mr-2 text-blue-500" />
                                {{ $t('stream.description') }}
                            </label>
                            <textarea v-model="form.description" rows="4" required
                                class="w-full px-5 py-4 rounded-2xl border-2 border-gray-200 dark:border-gray-600 bg-white/50 dark:bg-gray-700/50 text-gray-900 dark:text-white focus:ring-4 focus:ring-blue-500/20 focus:border-blue-500 transition-all duration-300 placeholder-gray-400 resize-none"
                                :placeholder="$t('stream.descriptionPlaceholder')" />
                        </div>

                        <!-- Language Selection -->
                        <div v-auto-animate class="grid grid-cols-1 md:grid-cols-2 gap-6">
                            <div class="space-y-3">
                                <label class="block text-sm font-semibold text-gray-800 dark:text-gray-200">
                                    <mdicon name="translate" size="18" class="inline mr-2 text-green-500" />
                                    {{ $t('stream.fromLanguage') }}
                                </label>
                                <select v-model="form.lang" required
                                    class="w-full px-5 py-4 rounded-2xl border-2 border-gray-200 dark:border-gray-600 bg-white/50 dark:bg-gray-700/50 text-gray-900 dark:text-white focus:ring-4 focus:ring-green-500/20 focus:border-green-500 transition-all duration-300">
                                    <option value="">{{ $t('stream.selectLanguage') }}</option>
                                    <option v-for="lang in userKnownLanguages" :key="lang" :value="lang">
                                        {{ getLangByCode(lang).name }}
                                    </option>

                                </select>
                                <p v-if="userKnownLanguages.length === 0"
                                    class="text-xs text-amber-600 dark:text-amber-400">
                                    <mdicon name="information" size="14" class="inline mr-1" />
                                    {{ $t('stream.noLanguagePreferences') }}
                                </p>
                            </div>

                            <div class="space-y-3">
                                <label class="block text-sm font-semibold text-gray-800 dark:text-gray-200">
                                    <mdicon name="translate-variant" size="18" class="inline mr-2 text-indigo-500" />
                                    {{ $t('stream.toLanguage') }}
                                </label>
                                <select v-model="form.target_lang" required
                                    class="w-full px-5 py-4 rounded-2xl border-2 border-gray-200 dark:border-gray-600 bg-white/50 dark:bg-gray-700/50 text-gray-900 dark:text-white focus:ring-4 focus:ring-indigo-500/20 focus:border-indigo-500 transition-all duration-300">
                                    <option value="">{{ $t('stream.selectLanguage') }}</option>
                                    <option v-for="lang in userKnownLanguages" :key="lang" :value="lang">
                                        {{ getLangByCode(lang).name }}
                                    </option>
                                </select>
                                <p v-if="userKnownLanguages.length === 0"
                                    class="text-xs text-amber-600 dark:text-amber-400">
                                    <mdicon name="information" size="14" class="inline mr-1" />
                                    {{ $t('stream.noLanguagePreferences') }}
                                </p>
                            </div>
                        </div>

                        <!-- Schedule Date & Time -->
                        <div v-auto-animate class="space-y-3">
                            <label class="block text-sm font-semibold text-gray-800 dark:text-gray-200">
                                <mdicon name="calendar-clock" size="18" class="inline mr-2 text-orange-500" />
                                {{ $t('stream.scheduledAt') }}
                            </label>
                            <input v-model="form.scheduled_at" type="datetime-local" required 
                                class="w-full px-5 py-4 rounded-2xl border-2 border-gray-200 dark:border-gray-600 bg-white/50 dark:bg-gray-700/50 text-gray-900 dark:text-white focus:ring-4 focus:ring-orange-500/20 focus:border-orange-500 transition-all duration-300" />
                        </div>

                        <!-- Stream Duration and Max Participants -->
                        <div v-auto-animate class="grid grid-cols-1 md:grid-cols-2 gap-6">
                            <div class="space-y-3">
                                <label class="block text-sm font-semibold text-gray-800 dark:text-gray-200">
                                    <mdicon name="timer" size="18" class="inline mr-2 text-red-500" />
                                    {{ $t('stream.duration') }} ({{ $t('stream.minutes') }})
                                </label>
                                <input v-model.number="form.duration" type="number" min="15" max="60" step="5" required
                                    class="w-full px-5 py-4 rounded-2xl border-2 border-gray-200 dark:border-gray-600 bg-white/50 dark:bg-gray-700/50 text-gray-900 dark:text-white focus:ring-4 focus:ring-red-500/20 focus:border-red-500 transition-all duration-300 placeholder-gray-400"
                                    :placeholder="$t('stream.durationPlaceholder')" />
                                <p class="text-xs text-gray-500 dark:text-gray-400">
                                    {{ $t('stream.durationHelper') }}
                                </p>
                            </div>

                            <div class="space-y-3">
                                <label class="block text-sm font-semibold text-gray-800 dark:text-gray-200">
                                    <mdicon name="account-group" size="18" class="inline mr-2 text-teal-500" />
                                    {{ $t('stream.maxParticipants') }}
                                </label>
                                <select v-model="form.max_participants" required
                                    class="w-full px-5 py-4 rounded-2xl border-2 border-gray-200 dark:border-gray-600 bg-white/50 dark:bg-gray-700/50 text-gray-900 dark:text-white focus:ring-4 focus:ring-teal-500/20 focus:border-teal-500 transition-all duration-300">
                                    <option value="">{{ $t('stream.selectMaxParticipants') }}</option>
                                    <option :value="1">1 {{ $t('stream.participants') }}</option>
                                    <option :value="2">2 {{ $t('stream.participants') }}</option>
                                    <option :value="3">3 {{ $t('stream.participants') }}</option>
                                    <option :value="4">4 {{ $t('stream.participants') }}</option>
                                    <option :value="5">5 {{ $t('stream.participants') }}</option>
                                </select>
                            </div>
                        </div>

                        <!-- Action Buttons -->
                        <div v-auto-animate class="flex flex-col sm:flex-row gap-4 pt-6">
                            <button type="button" @click="$router.push('/streamers')"
                                class="flex-1 px-8 py-4 rounded-2xl border-2 border-gray-200 dark:border-gray-600 text-gray-700 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-gray-700 transition-all duration-300 font-medium">
                                <mdicon name="arrow-left" size="20" class="inline mr-2" />
                                {{ $t('common.cancel') }}
                            </button>
                            <button type="submit" :disabled="loading"
                                class="flex-1 px-8 py-4 rounded-2xl bg-gradient-to-r from-purple-500 via-blue-500 to-indigo-600 hover:from-purple-600 hover:via-blue-600 hover:to-indigo-700 text-white font-semibold transition-all duration-300 transform hover:scale-[1.02] disabled:opacity-50 disabled:cursor-not-allowed shadow-lg hover:shadow-xl">
                                <div v-if="loading" class="flex items-center justify-center">
                                    <div
                                        class="animate-spin w-5 h-5 border-2 border-white border-t-transparent rounded-full mr-3">
                                    </div>
                                    {{ $t('common.creating') }}
                                </div>
                                <div v-else class="flex items-center justify-center">
                                    <mdicon name="video-plus" size="20" class="mr-3" />
                                    {{ $t('stream.createStream') }}
                                </div>
                            </button>
                        </div>
                    </form>
                </div>
            </div>
            <div v-if="isFormValid" v-auto-animate class="max-w-lg mx-auto mt-12">
                <div>
                    <div class="flex items-center mb-6 space-x-2">
                        <div class="bg-gradient-to-r from-purple-500 to-blue-500 rounded-full p-1">
                            <mdicon name="eye" class="text-white" />
                        </div>
                        <div class="text-xl font-bold text-center text-gray-900 dark:text-white">
                            {{ $t('stream.preview') }}
                        </div>
                    </div>
                    <PreviewCard :form="form" :username="getUser().username" />
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
import { getLangByCode } from '@/utils/language/languages'
import PreviewCard from '@/components/stream/PreviewCard.vue'

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
const userKnownLanguages = ref([])



// Load user language preferences
const loadUserLanguagePreferences = async () => {

    const response = await useFetch('/user/language-preferences')

    if (response && response.known_languages) {
        userKnownLanguages.value = response.known_languages || []
    }

}

// Check if user has teacher role
const checkTeacherRole = () => {
    const user = getUser()
    if (!user || !user.is_teacher) {
        toast.error(t('stream.teacherRoleRequired'))
        router.push('/streamers')
        return false
    }
    return true
}
function toUTCISOString(localDateTimeStr) {
    const localDate = new Date(localDateTimeStr)
    return localDate.toISOString()
}
const minDateTime = computed(() => {
    const now = new Date()
    now.setDate(now.getDate() + 1) // local günde 1 artır
    const year = now.getFullYear()
    const month = String(now.getMonth() + 1).padStart(2, '0')
    const day = String(now.getDate()).padStart(2, '0')
    const hours = String(now.getHours()).padStart(2, '0')
    const minutes = String(now.getMinutes()).padStart(2, '0')
    return `${year}-${month}-${day}T${hours}:${minutes}`
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

// Create stream
const createStream = async () => {
    if (!checkTeacherRole()) return

    loading.value = true
    try {
        const response = await useFetch('/stream/create', {
            method: 'POST',
            body: {
                title: form.value.title,
                description: form.value.description,
                lang: form.value.lang,
                target_lang: form.value.target_lang,
                scheduled_at: toUTCISOString(form.value.scheduled_at),
                duration: form.value.duration,
                max_participants: form.value.max_participants
            }
        })

        toast.success(t('stream.createSuccess'))
        router.push('/streamers')
    } catch (error) {
        console.error('Error creating stream:', error)
        toast.error(t('stream.createError'))
    } finally {
        loading.value = false
    }
}

// Initialize
onMounted(async () => {
    checkTeacherRole()
    await loadUserLanguagePreferences()
    form.value.scheduled_at = minDateTime.value
})
</script>
