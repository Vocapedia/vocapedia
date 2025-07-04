<template>
    <div class="container mx-auto p-4 space-y-6">
        <!-- Header -->
        <div class="text-center space-y-4">
            <h1 class="text-3xl font-bold text-gray-900 dark:text-white">
                {{ $t('home.streamers') }}
            </h1>
            <p class="text-gray-600 dark:text-gray-400 max-w-2xl mx-auto">
                {{ $t('home.streamers_description') }}
            </p>
        </div>

        <!-- Active Streams (Live Now) -->
        <div class="space-y-4">
            <h2 class="text-xl font-semibold text-gray-800 dark:text-gray-200">
                {{ $t('stream.activeStreams') }}
            </h2>
            
            <div v-if="activeStreams.length === 0" class="text-center py-12">
                <div class="w-16 h-16 bg-gray-100 dark:bg-gray-700 rounded-full flex items-center justify-center mx-auto mb-4">
                    <mdicon name="video-off" size="24" class="text-gray-400" />
                </div>
                <p class="text-gray-500 dark:text-gray-400">{{ $t('stream.noActiveStreams') }}</p>
            </div>

            <div v-else class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
                <router-link 
                    v-for="stream in activeStreams" 
                    :key="stream.room_id" 
                    :to="'/stream/' + stream.room_id" 
                    class="group">
                    <div class="card smooth-click2 bg-white dark:bg-zinc-800 rounded-xl transition-all duration-300 ease-in-out transform group-hover:scale-105">
                        <div class="space-y-2 w-full">
                            <div class="flex items-center justify-between px-3 pt-3">
                                <div class="text-xl font-semibold text-gray-900 dark:text-white">
                                    {{ stream.title }}
                                </div>
                                <div class="flex items-center space-x-1">
                                    <div class="w-2 h-2 bg-green-500 rounded-full animate-pulse"></div>
                                    <span class="text-xs text-green-600 dark:text-green-400">Live</span>
                                </div>
                            </div>
                            <div class="text-sm px-3 text-gray-600 dark:text-gray-400">
                                {{ stream.description }}
                            </div>
                            <div class="flex items-center justify-between px-3 pb-3">
                                <div class="flex items-center space-x-2">
                                    <mdicon name="account" size="16" class="text-gray-500" />
                                    <span class="text-sm text-gray-600 dark:text-gray-400">{{ stream.participants }} participants</span>
                                </div>
                                <div class="flex flex-col items-end text-xs text-gray-500 dark:text-gray-400">
                                    <span>{{ stream.lang }} → {{ stream.target_lang }}</span>
                                    <span>{{ formatTime(stream.scheduled_at) }}</span>
                                </div>
                            </div>
                        </div>
                    </div>
                </router-link>
            </div>
        </div>

        <!-- Upcoming Streams (Next 12h) -->
        <div class="space-y-4">
            <h2 class="text-xl font-semibold text-gray-800 dark:text-gray-200">
                {{ $t('stream.upcomingStreams') }}
            </h2>
            
            <div v-if="upcomingStreams.length === 0" class="text-center py-12">
                <div class="w-16 h-16 bg-gray-100 dark:bg-gray-700 rounded-full flex items-center justify-center mx-auto mb-4">
                    <mdicon name="video-off" size="24" class="text-gray-400" />
                </div>
                <p class="text-gray-500 dark:text-gray-400">{{ $t('stream.noUpcomingStreams') }}</p>
            </div>

            <div v-else class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
                <router-link 
                    v-for="stream in upcomingStreams" 
                    :key="stream.room_id" 
                    :to="'/stream/' + stream.room_id" 
                    class="group">
                    <div class="card smooth-click2 bg-white dark:bg-zinc-800 rounded-xl transition-all duration-300 ease-in-out transform group-hover:scale-105">
                        <div class="space-y-2 w-full">
                            <div class="flex items-center justify-between px-3 pt-3">
                                <div class="text-xl font-semibold text-gray-900 dark:text-white">
                                    {{ stream.title }}
                                </div>
                                <div class="flex items-center space-x-1">
                                    <div class="w-2 h-2 bg-blue-500 rounded-full"></div>
                                    <span class="text-xs text-blue-600 dark:text-blue-400">Scheduled</span>
                                </div>
                            </div>
                            <div class="text-sm px-3 text-gray-600 dark:text-gray-400">
                                {{ stream.description }}
                            </div>
                            <div class="flex items-center justify-between px-3 pb-3">
                                <div class="flex items-center space-x-2">
                                    <mdicon name="account" size="16" class="text-gray-500" />
                                    <span class="text-sm text-gray-600 dark:text-gray-400">{{ stream.participants }} participants</span>
                                </div>
                                <div class="flex flex-col items-end text-xs text-gray-500 dark:text-gray-400">
                                    <span>{{ stream.lang }} → {{ stream.target_lang }}</span>
                                    <span>{{ formatScheduledTime(stream.scheduled_at) }}</span>
                                </div>
                            </div>
                        </div>
                    </div>
                </router-link>
            </div>
        </div>

        <!-- Recent Streams (Ended >20m ago) -->
        <div class="space-y-4">
            <h2 class="text-xl font-semibold text-gray-800 dark:text-gray-200">
                {{ $t('stream.recentStreams') }}
            </h2>
            
            <div v-if="recentStreams.length === 0" class="text-center py-12">
                <div class="w-16 h-16 bg-gray-100 dark:bg-gray-700 rounded-full flex items-center justify-center mx-auto mb-4">
                    <mdicon name="video-off" size="24" class="text-gray-400" />
                </div>
                <p class="text-gray-500 dark:text-gray-400">{{ $t('stream.noRecentStreams') }}</p>
            </div>

            <div v-else class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
                <div 
                    v-for="stream in recentStreams" 
                    :key="stream.room_id"
                    class="card bg-white dark:bg-zinc-800 rounded-xl p-4 space-y-3 opacity-60">
                    <div class="flex items-center justify-between">
                        <div class="text-lg font-medium text-gray-900 dark:text-white">
                            {{ stream.title }}
                        </div>
                        <div class="text-xs text-gray-500 dark:text-gray-400">
                            Ended
                        </div>
                    </div>
                    <div class="text-sm text-gray-600 dark:text-gray-400">
                        {{ stream.description }}
                    </div>
                    <div class="flex items-center justify-between text-xs text-gray-500 dark:text-gray-400">
                        <div class="flex items-center space-x-2">
                            <span>{{ stream.participants }} participants</span>
                            <span>•</span>
                            <span>{{ stream.lang }} → {{ stream.target_lang }}</span>
                        </div>
                        <span>{{ formatTime(stream.scheduled_at) }}</span>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from "vue"
import { useFetch } from '@/composable/useFetch'

const activeStreams = ref([])
const upcomingStreams = ref([])
const recentStreams = ref([])
const loading = ref(false)

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
                participants: 3
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
                participants: 0
            },
            {
                room_id: "test-upcoming-2",
                title: "Spanish Vocabulary Building",
                description: "Expand your Spanish vocabulary",
                lang: "es",
                target_lang: "en",
                scheduled_at: new Date(Date.now() + 30 * 60 * 1000).toISOString(),
                participants: 2
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
                participants: 5
            },
            {
                room_id: "test-recent-2",
                title: "German Conversation Club",
                description: "Practice German conversation in a friendly environment",
                lang: "de",
                target_lang: "en",
                scheduled_at: new Date(Date.now() - 45 * 60 * 1000).toISOString(),
                participants: 4
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
