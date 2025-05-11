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
import { useFetch } from '@/composable/useFetch';
import { GetLang } from '@/i18n/i18n';
import { getUser } from '@/utils/token';
import { onMounted, watch, ref } from 'vue'
import { useRoute } from 'vue-router';
const isMeeting = ref(false)
const route = useRoute()
const api = ref(null);
const password = ref("")
onMounted(async () => {
    await useFetch('/stream/' + route.params.id)
        .then((response) => {
            console.log(response);
            password.value = response.password

        })
        .catch((error) => {
            console.error(error);
        });
});
watch(isMeeting, (val) => {
    if (val) {
        const domain = 'meet.jit.si';
        const options = {
            roomName: route.params.id,
            parentNode: document.querySelector('#meet'),
            lang: GetLang(),
            userInfo: {
                displayName: getUser().username,
            },
            configOverwrite: {
                prejoinPageEnabled: false,
                disableInviteFunctions: true,
                startWithAudioMuted: false,
                startWithVideoMuted: false,
                enableWelcomePage: false,
                authentication: false,
                disableModerator: true,
                enableLobby: false,
                disableModeratorIndicator: true,
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

        api.value = new JitsiMeetExternalAPI(domain, options);
        api.value.addEventListener('videoConferenceJoined', async () => {
            await api.value.executeCommand('password', password)
        })
        api.value.addEventListener('passwordRequired', async () => {
            await api.value.executeCommand('password', password)
        })

    } else {
        if (api.value) {
            api.value.dispose();
            api.value = null;
        }
    }
});
</script>
