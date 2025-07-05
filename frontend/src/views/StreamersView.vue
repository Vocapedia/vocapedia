<template>
    <div v-auto-animate>
        <div class="container mx-auto p-6 space-y-8">
            <!-- Header -->
            <div class="text-center space-y-6">
                <div
                    class="inline-flex items-center justify-center w-20 h-20 bg-gradient-to-br from-blue-500 via-indigo-500 to-purple-600 rounded-3xl mb-4 shadow-lg">
                    <mdicon name="video-account" size="40" class="text-white" />
                </div>
                <h1
                    class="text-4xl font-bold bg-gradient-to-r from-blue-600 to-purple-600 bg-clip-text text-transparent">
                    {{ $t('home.streamers') }}
                </h1>
                <p class="text-gray-600 dark:text-gray-400 max-w-2xl mx-auto text-lg">
                    {{ $t('home.streamers_description') }}
                </p>
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
                    <!-- Active Streams Tab -->
                    <div v-if="activeTab === 'active'" class="space-y-4">
                        <div v-if="activeStreams.length === 0" class="text-center py-12">
                            <div
                                class="w-16 h-16 bg-gray-100 dark:bg-gray-700 rounded-full flex items-center justify-center mx-auto mb-4">
                                <mdicon name="video-off" size="24" class="text-gray-400" />
                            </div>
                            <p class="text-gray-500 dark:text-gray-400">{{ $t('stream.noActiveStreams') }}</p>
                        </div>

                        <div v-else class="grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                            <router-link v-for="stream in activeStreams" :key="stream.room_id"
                                :to="'/stream/' + stream.room_id" class="group">
                                <PreviewCard class="smooth-click2" :form="{
                                    title: stream.title,
                                    description: stream.description,
                                    lang: stream.lang,
                                    target_lang: stream.target_lang,
                                    scheduled_at: stream.scheduled_at,
                                    duration: stream.duration,
                                    max_participants: stream.max_participants,
                                }" :username="stream.creator.username" />
                            </router-link>
                        </div>
                    </div>

                    <!-- Upcoming Streams Tab -->
                    <div v-if="activeTab === 'upcoming'" class="space-y-4">
                        <div v-if="upcomingStreams.length === 0" class="text-center py-12">
                            <div
                                class="w-16 h-16 bg-gray-100 dark:bg-gray-700 rounded-full flex items-center justify-center mx-auto mb-4">
                                <mdicon name="video-off" size="24" class="text-gray-400" />
                            </div>
                            <p class="text-gray-500 dark:text-gray-400">{{ $t('stream.noUpcomingStreams') }}</p>
                        </div>

                        <div v-else class="grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                            <div v-for="stream in upcomingStreams" :key="stream.room_id" class="group">
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

                    <!-- Recent Streams Tab -->
                    <div v-if="activeTab === 'recent'" class="space-y-4">
                        <div v-if="recentStreams.length === 0" class="text-center py-12">
                            <div
                                class="w-16 h-16 bg-gray-100 dark:bg-gray-700 rounded-full flex items-center justify-center mx-auto mb-4">
                                <mdicon name="video-off" size="24" class="text-gray-400" />
                            </div>
                            <p class="text-gray-500 dark:text-gray-400">{{ $t('stream.noRecentStreams') }}</p>
                        </div>

                        <div v-else class="grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">

                            <PreviewCard v-for="stream in recentStreams" :form="{
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
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted, computed, watch } from "vue"
import { useRoute, useRouter } from 'vue-router'
import { useFetch } from '@/composable/useFetch'
import { getUser } from '@/utils/token'
import PreviewCard from "@/components/stream/PreviewCard.vue"

const route = useRoute()
const router = useRouter()
const activeStreams = ref([])
const upcomingStreams = ref([])
const recentStreams = ref([])
const loading = ref(false)

// Tab management
const activeTab = ref(route.query.tab || 'active')
const tabs = [
    { key: 'active', label: 'Aktif Yayınlar' },
    { key: 'upcoming', label: 'Gelecek Yayınlar' },
    { key: 'recent', label: 'Son Yayınlar' }
]

// Watch for tab changes and update URL
watch(activeTab, (newTab) => {
    router.push({ query: { ...route.query, tab: newTab } })
})

// Watch for route changes and update active tab
watch(() => route.query.tab, (newTab) => {
    if (newTab && tabs.find(t => t.key === newTab)) {
        activeTab.value = newTab
    }
})

const canCreateStream = computed(() => {
    const user = getUser()
    return user && user.is_teacher
})

const formatTime = (timestamp) => {
    const date = new Date(timestamp)
    const now = new Date()
    const diff = now - date
    const minutes = Math.floor(diff / 60000)
    const hours = Math.floor(diff / 3600000)
    const days = Math.floor(diff / 86400000)

    if (days > 0) return `${days}d ago`
    if (hours > 0) return `${hours}h ago`
    if (minutes > 0) return `${minutes}m ago`
    return 'Just now'
}

const formatScheduledTime = (timestamp) => {
    const date = new Date(timestamp)
    const now = new Date()
    const diff = date - now
    const minutes = Math.floor(diff / 60000)
    const hours = Math.floor(diff / 3600000)
    const days = Math.floor(diff / 86400000)

    if (diff < 0) return 'Started'
    if (days > 0) return `in ${days}d`
    if (hours > 0) return `in ${hours}h`
    if (minutes > 0) return `in ${minutes}m`
    return 'Starting soon'
}

const loadStreams = async () => {
    loading.value = true
    try {
        // Load active streams (within 20m of schedule)
        const activeData = await useFetch('/stream/active')
        activeStreams.value = activeData || []

        // Load upcoming streams (next 12h)
        const upcomingData = await useFetch('/stream/upcoming')
        upcomingStreams.value = upcomingData || []

        // Load recent streams (ended >20m ago)
        const recentData = await useFetch('/stream/recent')
        recentStreams.value = recentData || []

    } catch (error) {
        console.error('Error loading streams:', error)
        // Mock data for testing
        activeStreams.value = [
            {
                room_id: "test-active-1",
                title: "English Conversation Practice",
                description: "Practice speaking English with native speakers",
                lang: "en",
                target_lang: "tr",
                scheduled_at: new Date(Date.now() - 5 * 60 * 1000).toISOString(),
                participants: 3,
                creator: {
                    username: "sarah_teacher",
                    name: "Sarah Johnson"
                }
            }
        ]

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

        recentStreams.value = [
            {
                room_id: "test-recent-1",
                title: "French Pronunciation Workshop",
                description: "Master French pronunciation techniques",
                lang: "fr",
                target_lang: "en",
                scheduled_at: new Date(Date.now() - 30 * 60 * 1000).toISOString(),
                participants: 5,
                creator: {
                    username: "pierre_teach",
                    name: "Pierre Dubois"
                }
            },
            {
                room_id: "test-recent-2",
                title: "German Conversation Club",
                description: "Practice German conversation in a friendly environment",
                lang: "de",
                target_lang: "en",
                scheduled_at: new Date(Date.now() - 45 * 60 * 1000).toISOString(),
                participants: 4,
                creator: {
                    username: "hans_lehrer",
                    name: "Hans Mueller"
                }
            }
        ]
    } finally {
        loading.value = false
    }
}

// Initialize
onMounted(() => {
    loadStreams()
})
</script>
