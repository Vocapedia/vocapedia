<template>
    <div v-auto-animate class="container mx-auto space-y-5 p-4">
        <!-- Stream Info Card -->
        <div class="card bg-white dark:bg-zinc-800 rounded-xl p-6 shadow-lg">
            <div class="text-center space-y-4">
                <h1 class="text-2xl font-bold text-gray-900 dark:text-white">
                    {{ $t('stream.title') }}
                </h1>
                <div class="flex flex-col sm:flex-row items-center justify-center gap-4">
                    <div class="flex items-center space-x-2">
                        <span class="text-sm text-gray-500 dark:text-gray-400">{{ $t('stream.roomId') }}:</span>
                        <code class="bg-gray-100 dark:bg-gray-700 px-2 py-1 rounded text-sm font-mono">
                            {{ route.params.id }}
                        </code>
                        <button @click="copyRoomId" class="smooth-click p-1 rounded hover:bg-gray-100 dark:hover:bg-gray-700">
                            <mdicon name="content-copy" size="16" />
                        </button>
                    </div>
                    <div class="flex items-center space-x-2">
                        <span class="text-sm text-gray-500 dark:text-gray-400">{{ $t('stream.status') }}:</span>
                        <span class="px-2 py-1 rounded-full text-xs font-medium"
                              :class="isMeeting ? 'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-200' : 'bg-gray-100 text-gray-800 dark:bg-gray-700 dark:text-gray-200'">
                            {{ isMeeting ? $t('stream.active') : $t('stream.inactive') }}
                        </span>
                    </div>
                </div>
            </div>
        </div>

        <!-- Meeting Container -->
        <div v-auto-animate class="flex items-center justify-center">
            <div v-show="isMeeting" id="meet" class="w-full rounded-xl overflow-hidden shadow-2xl" style="height: 700px;" />
            <div v-show="!isMeeting" class="card bg-gradient-to-br from-blue-50 to-indigo-50 dark:from-blue-900/20 dark:to-indigo-900/20 rounded-xl p-12 text-center">
                <div class="space-y-4">
                    <div class="w-20 h-20 bg-blue-100 dark:bg-blue-900/50 rounded-full flex items-center justify-center mx-auto">
                        <mdicon name="video" size="32" class="text-blue-600 dark:text-blue-400" />
                    </div>
                    <h3 class="text-xl font-semibold text-gray-900 dark:text-white">{{ $t('stream.waitingToStart') }}</h3>
                    <p class="text-gray-600 dark:text-gray-400 max-w-md mx-auto">{{ $t('stream.waitingDescription') }}</p>
                </div>
            </div>
        </div>

        <!-- Control Panel -->
        <div class="flex flex-col sm:flex-row items-center justify-center gap-4">
            <button 
                class="smooth-click card px-6 py-3 rounded-xl font-medium transition-all duration-200 transform hover:scale-105"
                :class="isMeeting ? 'bg-red-500 hover:bg-red-600 text-white' : 'bg-green-500 hover:bg-green-600 text-white'"
                @click="toggleMeeting"
                :disabled="loading">
                <div class="flex items-center space-x-2">
                    <mdicon :name="isMeeting ? 'stop' : 'play'" size="20" />
                    <span>{{ isMeeting ? $t('stream.stop') : $t('stream.start') }}</span>
                </div>
            </button>
            
        </div>

        <!-- Loading State -->
        <div v-if="loading" class="fixed inset-0 bg-black/50 backdrop-blur-sm flex items-center justify-center z-50">
            <div class="card bg-white dark:bg-zinc-800 rounded-xl p-6 text-center">
                <div class="animate-spin w-8 h-8 border-4 border-blue-500 border-t-transparent rounded-full mx-auto mb-4"></div>
                <p class="text-gray-600 dark:text-gray-400">{{ $t('stream.loading') }}</p>
            </div>
        </div>
    </div>
</template>

<script setup>
import { useFetch } from '@/composable/useFetch';
import { useToast } from '@/composable/useToast';
import { GetLang } from '@/i18n/i18n';
import { getUser } from '@/utils/token';
import { onMounted, onUnmounted, watch, ref } from 'vue'
import { useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'

const isMeeting = ref(false)
const loading = ref(false)
const route = useRoute()
const toast = useToast()
const { t } = useI18n()
const api = ref(null)
const password = ref("")
// Whether room info loaded successfully
const roomExists = ref(false)


// Copy room ID to clipboard
const copyRoomId = async () => {
    try {
        await navigator.clipboard.writeText(route.params.id)
        toast.success(t('stream.copySuccess'))
    } catch (error) {
        toast.error(t('stream.copyError'))
    }
}


// Toggle meeting state
const toggleMeeting = () => {
    if (loading.value || !roomExists.value) return
    isMeeting.value = !isMeeting.value
}

// Initialize room data
onMounted(async () => {
    // Ensure a room ID is present, otherwise do nothing
    if (!route.params.id) {
        return
    }

    loading.value = true
    try {
        const response = await useFetch('/stream/' + route.params.id)
        password.value = response.password
        roomExists.value = true
    } catch (error) {
        console.error('Error fetching stream data:', error)
        toast.error(t('stream.loadError'))
        // Still allow the stream to work with a default password
        password.value = "vocapedia"
    } finally {
        loading.value = false
    }
})

// Watch for meeting state changes
watch(isMeeting, (val) => {
    if (val) {
        startJitsiMeeting()
    } else {
        stopJitsiMeeting()
    }
})

// Start Jitsi meeting
const startJitsiMeeting = () => {
    if (!window.JitsiMeetExternalAPI) {
        toast.error(t('stream.startError'))
        return
    }

    const domain = 'meet.jit.si'
    const user = getUser()
    
    const options = {
        roomName: route.params.id,
        parentNode: document.querySelector('#meet'),
        lang: GetLang(),
        userInfo: {
            displayName: user?.username || 'Anonymous',
        },
        configOverwrite: {
            prejoinPageEnabled: false,
            disableInviteFunctions: true,
            startWithAudioMuted: false,
            startWithVideoMuted: false,
            enableWelcomePage: false,
            authentication: false,
            disableModerator: false, // Allow both users to have equal rights
            enableLobby: false,
            disableModeratorIndicator: true,
            defaultLocalDisplayName: user?.username || 'Anonymous',
            defaultRemoteDisplayName: 'Participant',
            // Disable features that create hierarchy
            disableDeepLinking: true,
            disableShortcuts: false,
            enableClosePage: false,
        },
        interfaceConfigOverwrite: {
            DEFAULT_REMOTE_DISPLAY_NAME: 'Participant',
            SHOW_JITSI_WATERMARK: false,
            SHOW_WATERMARK_FOR_GUESTS: false,
            SHOW_BRAND_WATERMARK: false,
            SHOW_POWERED_BY: false,
            TOOLBAR_BUTTONS: [
                'microphone', 'camera', 'closedcaptions', 'desktop', 'fullscreen',
                'fodeviceselection', 'hangup', 'profile', 'chat', 'recording',
                'livestreaming', 'etherpad', 'sharedvideo', 'settings', 'raisehand',
                'videoquality', 'filmstrip', 'invite', 'feedback', 'stats', 'shortcuts',
                'tileview', 'videobackgroundblur', 'download', 'help', 'mute-everyone',
                'e2ee'
            ],
            MOBILE_APP_PROMO: false,
            HIDE_INVITE_MORE_HEADER: true,
        }
    }

    try {
        api.value = new JitsiMeetExternalAPI(domain, options)
        
        // Event listeners
        api.value.addEventListener('videoConferenceJoined', async () => {
            console.log('User joined the conference')
            if (password.value && password.value !== "vocapedia") {
                await api.value.executeCommand('password', password.value)
            }
        })
        
        api.value.addEventListener('passwordRequired', async () => {
            console.log('Password required')
            if (password.value) {
                await api.value.executeCommand('password', password.value)
            }
        })
        
        api.value.addEventListener('videoConferenceLeft', () => {
            console.log('User left the conference')
            isMeeting.value = false
        })
        
        api.value.addEventListener('readyToClose', () => {
            console.log('Ready to close')
            isMeeting.value = false
        })
        
    } catch (error) {
        console.error('Error starting Jitsi meeting:', error)
        toast.error(t('stream.startError'))
        isMeeting.value = false
    }
}

// Stop Jitsi meeting
const stopJitsiMeeting = () => {
    if (api.value) {
        try {
            api.value.dispose()
            api.value = null
        } catch (error) {
            console.error('Error stopping Jitsi meeting:', error)
        }
    }
}

// Cleanup on unmount
onUnmounted(() => {
    stopJitsiMeeting()
})
</script>
