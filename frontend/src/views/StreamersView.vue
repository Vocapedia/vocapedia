<template>
    <div v-auto-animate>
        <div class="container mx-auto p-6 space-y-8">
            <!-- Header -->
            <div class="text-center space-y-6">
                <div class="inline-flex items-center justify-center w-20 h-20 rounded-3xl mb-4 card pr-2">
                    <span class="font-logo text-6xl">V</span>
                </div>
                <h1
                    class="text-4xl font-bold bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-transparent">
                    {{ $t('streamers.header') }}
                </h1>
                <p class="text-gray-600 dark:text-gray-400 max-w-2xl mx-auto text-lg">
                    {{ $t('streamers.description') }}
                </p>

                <!-- Token Balance -->
                <div
                    class="inline-flex items-center bg-blue-50 dark:bg-blue-900/20 border border-blue-200 dark:border-blue-800 rounded-lg px-4 py-2">
                    <span class="text-sm font-medium text-blue-800 dark:text-blue-200">
                        Token: {{ userTokens }}
                    </span>
                </div>
            </div>

            <!-- Action Buttons -->
            <div v-auto-animate class="flex justify-center space-x-4">
                <router-link v-if="canCreateStream" to="/create-stream"
                    class="smooth-click bg-gradient-to-r from-purple-500 via-blue-500 to-indigo-600 hover:from-purple-600 hover:via-blue-600 hover:to-indigo-700 text-white px-8 py-4 rounded-2xl font-semibold transition-all duration-300 transform hover:scale-105 shadow-lg hover:shadow-xl">
                    <div class="flex items-center space-x-3">
                        <mdicon name="video-plus" size="24" />
                        <span>{{ $t('stream.createStream') }}</span>
                    </div>
                </router-link>
            </div>

            <!-- Tabs -->
            <div v-auto-animate class="">
                <nav class="flex space-x-8 justify-center p-6 border-b border-gray-200/50 dark:border-gray-700/50">
                    <button v-for="tab in tabs" :key="tab.key" @click="activeTab = tab.key" :class="[
                        'smooth-click py-3 px-6 text-sm font-semibold border-b-3 transition-all duration-300',
                        activeTab === tab.key
                            ? 'border-sky-500 text-sky-600 dark:text-sky-400  dark:bg-sky-900/20'
                            : 'border-transparent text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200 hover:bg-gray-50 dark:hover:bg-gray-700/30'
                    ]">
                        {{ tab.label }}
                    </button>
                </nav>

                <!-- Tab Content -->
                <div v-auto-animate class="p-6">
                    <!-- Upcoming Streams -->
                    <div v-if="activeTab === 'upcoming'" class="space-y-4">
                        <div v-if="upcomingStreams.length === 0" class="text-center py-12">
                            <div
                                class="w-16 h-16 bg-gray-100 dark:bg-gray-700 rounded-full flex items-center justify-center mx-auto mb-4">
                                <mdicon name="video-off" size="24" class="text-gray-400" />
                            </div>
                            <p class="text-gray-500 dark:text-gray-400">{{ $t('stream.noUpcomingStreams') }}</p>
                        </div>

                        <div v-else class="grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                            <div v-for="stream in upcomingStreams" :key="stream.room_id" class="group cursor-pointer"
                                @click="handleStreamClick(stream)">
                                <PreviewCard :form="{
                                    title: stream.title,
                                    description: stream.description,
                                    lang: stream.lang,
                                    target_lang: stream.target_lang,
                                    scheduled_at: stream.scheduled_at,
                                    duration: stream.duration,
                                    max_participants: stream.max_participants,
                                }" :username="stream.creator.username" />
                            </div>
                        </div>
                    </div>

                    <!-- Joined Streams -->
                    <div v-if="activeTab === 'joined'" class="space-y-4">
                        <div v-if="joinedStreams.length === 0" class="text-center py-12">
                            <div
                                class="w-16 h-16 bg-gray-100 dark:bg-gray-700 rounded-full flex items-center justify-center mx-auto mb-4">
                                <mdicon name="account-check" size="24" class="text-gray-400" />
                            </div>
                            <p class="text-gray-500 dark:text-gray-400">Henüz hiçbir stream'e katılmadınız</p>
                        </div>

                        <div v-else class="grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                            <div v-for="stream in joinedStreams" :key="stream.room_id" class="relative">
                                <div class="cursor-pointer" @click="handleJoinedStreamClick(stream)">
                                    <PreviewCard :form="{
                                        title: stream.title,
                                        description: stream.description,
                                        lang: stream.lang,
                                        target_lang: stream.target_lang,
                                        scheduled_at: stream.scheduled_at,
                                        duration: stream.duration,
                                        max_participants: stream.max_participants,
                                    }" :username="stream.creator.username" />
                                </div>

                                <!-- Cancel button -->
                                <div v-if="stream.can_cancel" class="absolute top-2 right-2">
                                    <button @click.stop="handleCancelRegistration(stream)"
                                        class="bg-red-500 hover:bg-red-600 text-white p-2 rounded-full shadow-lg transition-colors">
                                        <mdicon name="close" size="16" />
                                    </button>
                                </div>

                                <!-- Status badge -->
                                <div class="absolute top-2 left-2">
                                    <span class="bg-green-500 text-white px-2 py-1 rounded-full text-xs font-medium">
                                        Kayıtlı
                                    </span>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <!-- Join Stream Modal -->
        <Popup v-model="showJoinModal">
            <template #header>
                <h3 class="text-xl font-bold text-gray-900 dark:text-white">
                    Stream'e Katıl
                </h3>
            </template>

            <template #description>
                <div class="space-y-4">
                    <div>
                        <h4 class="font-semibold text-gray-800 dark:text-gray-200">
                            {{ selectedStream?.title }}
                        </h4>
                        <p class="text-sm text-gray-600 dark:text-gray-400 mt-1">
                            {{ selectedStream?.description }}
                        </p>
                    </div>

                    <div
                        class="bg-yellow-50 dark:bg-yellow-900/20 border border-yellow-200 dark:border-yellow-800 rounded-lg p-4">
                        <div class="flex items-center">
                            <mdicon name="information" size="20" class="text-yellow-600 dark:text-yellow-400 mr-2" />
                            <span class="text-sm text-yellow-800 dark:text-yellow-200">
                                Bu stream'e katılmak için <strong>2 token</strong> gerekiyor.
                            </span>
                        </div>
                        <div class="mt-2 text-sm text-yellow-700 dark:text-yellow-300">
                            Mevcut token: <strong>{{ userTokens }}</strong>
                        </div>
                    </div>
                </div>
            </template>

            <template #buttons>
                <div class="flex space-x-3">
                    <button @click="showJoinModal = false"
                        class="flex-1 px-4 py-2 bg-gray-300 dark:bg-gray-600 text-gray-700 dark:text-gray-300 rounded-lg hover:bg-gray-400 dark:hover:bg-gray-500 transition-colors">
                        İptal
                    </button>
                    <button @click="confirmJoin" :disabled="loading || userTokens < 2"
                        class="flex-1 px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 disabled:opacity-50 disabled:cursor-not-allowed transition-colors">
                        {{ loading ? 'Katılıyor...' : (userTokens < 2 ? 'Yetersiz Token' : 'Katıl') }} </button>
                </div>
            </template>
        </Popup>
    </div>
</template>

<script setup>
import { ref, onMounted, computed, watch } from "vue"
import { useRoute, useRouter } from 'vue-router'
import { useFetch } from '@/composable/useFetch'
import { getUser } from '@/utils/token'
import PreviewCard from "@/components/stream/PreviewCard.vue"
import Popup from "@/components/shared/Popup.vue"
import { useToast } from "@/composable/useToast"

const route = useRoute()
const router = useRouter()
const toast = useToast()
const upcomingStreams = ref([])
const joinedStreams = ref([])
const loading = ref(false)
const showJoinModal = ref(false)
const selectedStream = ref(null)
const userTokens = ref(0)

// Tab management
const activeTab = ref(route.query.tab || 'upcoming')
const tabs = [
    { key: 'upcoming', label: 'Gelecek Yayınlar' },
    { key: 'joined', label: 'Katıldığım Streamler' }
]

const canCreateStream = computed(() => {
    const user = getUser()
    return user && user.is_teacher
})

watch(activeTab, (newTab) => {
    router.push({ query: { ...route.query, tab: newTab } })
})

watch(() => route.query.tab, (newTab) => {
    if (newTab && tabs.find(t => t.key === newTab)) {
        activeTab.value = newTab
    }
})

const loadUserTokens = async () => {
    const data = await useFetch('/user/tokens')
    userTokens.value = data.tokens || 0
}

const handleStreamClick = (stream) => {
    selectedStream.value = stream
    showJoinModal.value = true
}

const handleJoinedStreamClick = (stream) => {
    const now = new Date()
    const streamTime = new Date(stream.scheduled_at)
    const endTime = new Date(streamTime.getTime() + 45 * 60 * 1000) // 45 minutes after start

    if (now >= streamTime && now <= endTime) {
        router.push(`/stream/${BigInt(stream.room_id)}`)
    } else if (now > endTime) {
        toast.error('Bu stream süresi dolmuş!')
    } else {
        toast.show('Stream henüz başlamadı!')
    }
}

const handleCancelRegistration = async (stream) => {
    if (!confirm('Bu stream kaydınızı iptal etmek istediğinizden emin misiniz? 2 token iade edilecek.')) {
        return
    }

    try {
        loading.value = true
        await useFetch(`/stream/${stream.room_id}/cancel`, {
            method: 'DELETE'
        })

        toast.success('Stream kaydınız iptal edildi! 2 token iade edildi.')

        // Refresh streams and tokens
        await loadJoinedStreams()
        await loadUserTokens()
    } catch (error) {
        console.error('Error cancelling registration:', error)
        if (error.message === 'Cancellation deadline has passed') {
            toast.error('İptal etme süresi geçmiş! Stream\'den 2 saat öncesine kadar iptal edebilirsiniz.')
        } else {
            toast.error('Kayıt iptal edilirken bir hata oluştu.')
        }
    } finally {
        loading.value = false
    }
}

const confirmJoin = async () => {
    if (!selectedStream.value) return

    try {
        loading.value = true
        await useFetch(`/stream/${BigInt(selectedStream.value.room_id)}/join`, {
            method: 'POST'
        })

        toast.success('Stream\'e başarıyla katıldınız!')
        showJoinModal.value = false
        selectedStream.value = null

        // Refresh streams and tokens
        await loadStreams()
        await loadJoinedStreams()
        await loadUserTokens()
    } catch (error) {
        console.error('Error joining stream:', error)
        if (error.message === 'Insufficient tokens') {
            toast.error('Yetersiz token! Bu stream\'e katılmak için 2 token gerekiyor.')
        } else if (error.message === 'Already registered for this stream') {
            toast.error('Bu stream\'e zaten kayıtlısınız!')
        } else if (error.message === 'Stream has reached maximum participants') {
            toast.error('Bu stream maksimum katılımcı sayısına ulaşmış!')
        } else {
            toast.error('Stream\'e katılırken bir hata oluştu.')
        }
    } finally {
        loading.value = false
    }
}

const loadStreams = async () => {
    loading.value = true
    try {
        // Load upcoming streams (next 12h)
        const upcomingData = await useFetch('/stream/upcoming')
        upcomingStreams.value = upcomingData || []

    } catch (error) {
        console.error('Error loading streams:', error)
        // Mock data for testing
        upcomingStreams.value = [
            {
                room_id: "test-upcoming-1",
                title: "Turkish Grammar Lesson",
                description: "Learn Turkish grammar fundamentals",
                lang: "tr",
                target_lang: "en",
                scheduled_at: new Date(Date.now() + 2 * 60 * 60 * 1000).toISOString(),
                participants: 0,
                creator: {
                    username: "mehmet_hoca",
                    name: "Mehmet Yılmaz"
                }
            },
            {
                room_id: "test-upcoming-2",
                title: "Spanish Vocabulary Building",
                description: "Expand your Spanish vocabulary",
                lang: "es",
                target_lang: "en",
                scheduled_at: new Date(Date.now() + 30 * 60 * 1000).toISOString(),
                participants: 2,
                creator: {
                    username: "maria_prof",
                    name: "María García"
                }
            }
        ]
    } finally {
        loading.value = false
    }
}

const loadJoinedStreams = async () => {
    try {
        const joinedData = await useFetch('/stream/joined')
        joinedStreams.value = joinedData || []
    } catch (error) {
        console.error('Error loading joined streams:', error)
        joinedStreams.value = []
    }
}

// Initialize
onMounted(() => {
    loadStreams()
    loadJoinedStreams()
    loadUserTokens()
})
</script>
