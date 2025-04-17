<template>
    <div v-auto-animate class="container mx-auto space-y-5">
        <div v-auto-animate class="flex items-center justify-center">
            <div v-show="isMeeting" id="meet" style="width: 100%; height: 700px;" />
        </div>
        <div class="text-center">
            <button class="smooth-click card px-3 py-1 rounded-full" @click="isMeeting = !isMeeting">
                {{ isMeeting ? 'Kapat' : 'Başlat' }}
            </button>
        </div>
    </div>
</template>

<script setup>
import { GetLang } from '@/i18n/i18n';
import { getUser } from '@/utils/token';
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router';
const isMeeting = ref(false)
const route = useRoute()
onMounted(() => {
    console.log(getUser())
    const domain = 'meet.jit.si';
    const options = {
        roomName: 'vocapedia_' + route.params.id,
        parentNode: document.querySelector('#meet'),
        lang: GetLang(),
        userInfo: {
            displayName: getUser().username
        },
        configOverwrite: {
            prejoinPageEnabled: false,
            disableInviteFunctions: true,
            startWithAudioMuted: false,
            startWithVideoMuted: false,
            enableWelcomePage: false,
            authentication: false,
            disableModerator: true,
            prejoinPageEnabled: false,
            startWithAudioMuted: false,
            startWithVideoMuted: false,
            enableLobby: false,
            disableModeratorIndicator: true
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
            ]
        }
    };
    const api = new JitsiMeetExternalAPI(domain, options);
});
</script>
