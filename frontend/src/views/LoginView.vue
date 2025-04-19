<template>
    <div class="pt-36 flex items-center justify-center ">
        <div class="z-1 w-full max-w-120">
            <div class="bg-white/90 dark:bg-zinc-800/90 p-8 pb-4 rounded-lg shadow-lg space-y-4">
                <div class="text-2xl text-center pb-4">
                    <RouterLink to="/" class="flex-shrink-0 py-2">
                        <div class="font-logo">
                            Vocapedia
                        </div>
                    </RouterLink>
                </div>
                <form v-if="isCodeSent" @submit.prevent="handleVerifyOTP">
                    <div class="mb-4 space-y-2">
                        <label for="otp" class="block text-sm font-medium">{{ $t('login.input.otp') }}</label>
                        <div class="flex justify-center space-x-2">
                            <input v-for="(digit, index) in otp" :key="index" type="text" inputmode="numeric"
                                maxlength="1"
                                class="w-10 h-10 text-center text-xl border rounded-md outline-none border-zinc-200 bg-zinc-100 dark:border-zinc-800 dark:bg-zinc-700"
                                v-model="otp[index]" @input="handleInput(index, $event)"
                                @keydown.delete="handleBackspace(index, $event)" />
                        </div>
                    </div>
                    <div class="text-center">
                        <button type="submit" :disabled="isLoading"
                            :class="isCaptchaVerified ? 'smooth-click2 bg-sky-700 focus:ring-sky-700 ' : 'bg-zinc-200 dark:bg-zinc-900'"
                            class="flex justify-between items-center py-2 px-4 rounded-lg w-full focus:outline-none">
                            <span>
                                <mdicon :class="isCaptchaVerified ? 'text-white dark:text-zinc-900' : ''"
                                    v-if="isLoading" name="loading" spin />
                            </span>
                            <span :class="isCaptchaVerified ? 'text-white dark:text-zinc-900' : ''"
                                class="font-semibold text-lg">{{ $t('login.otp.verify_button') }}</span>
                            <span>
                                <mdicon v-if="isLoading" name="loading" spin
                                    :class="isCaptchaVerified ? 'text-sky-700' : 'text-zinc-200 dark:text-zinc-900'" />
                            </span>
                        </button>
                    </div>

                </form>
                <form v-else @submit.prevent="handleOTPSend">
                    <div class="mb-4">
                        <label for="email" class="block text-sm font-medium">{{ $t('login.input.label') }}</label>
                        <input type="email" id="email" v-model="email" required class="w-full p-3 border shadow-sm outline-none transition-all 
             bg-zinc-100 text-zinc-900 py-2 mt-3 border-none rounded-full
             dark:bg-zinc-900 dark:text-white" :placeholder="$t('login.input.placeholder')" />
                    </div>
                    <div class="mb-4" id="recaptcha-container"></div>

                    <div class="text-center">
                        <button type="submit" :disabled="isLoading"
                            :class="isCaptchaVerified ? 'smooth-click2 bg-sky-700 focus:ring-sky-700 ' : 'bg-zinc-200 dark:bg-zinc-900'"
                            class="flex justify-between items-center py-2 px-4 rounded-lg w-full focus:outline-none">
                            <span>
                                <mdicon :class="isCaptchaVerified ? 'text-white dark:text-zinc-900' : ''"
                                    v-if="isLoading" name="loading" spin />
                            </span>
                            <span :class="isCaptchaVerified ? 'text-white dark:text-zinc-900' : ''"
                                class="font-semibold text-lg">{{ $t('login.otp.send_button') }}</span>
                            <span>
                                <mdicon v-if="isLoading" name="loading" spin
                                    :class="isCaptchaVerified ? 'text-sky-700' : 'text-zinc-200 dark:text-zinc-900'" />
                            </span>
                        </button>
                    </div>
                </form>
                <div v-if="isCodeSent" class="text-center">
                    {{ remaining }}
                </div>
            </div>
        </div>
    </div>

</template>

<script setup>
import { useFetch } from '@/composable/useFetch';
import { useToast } from '@/composable/useToast';
import { i18n } from '@/i18n/i18n';
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
const otp = ref(new Array(6).fill(""));
const isCodeSent = ref(false)
const isLoading = ref(false);
const email = ref('');
const emailOTP = ref('')
const isCaptchaVerified = ref(false);
const router = useRouter()
const recaptchaSiteKey = import.meta.env.VITE_RECAPTCHA_SITE_KEY;
const toast = useToast()
import { useCountdown } from '@vueuse/core'
import { getDevice } from '@/utils/token';

const duration = 5 * 60;
const { remaining, start } = useCountdown(duration, {
    onComplete() {
        router.go(0)
    },
    onTick() {

    }
})

const handleInput = (index, event) => {
    const value = event.target.value;
    if (/^\d$/.test(value)) {
        otp.value[index] = value;
        if (index < otp.value.length - 1) {
            event.target.nextElementSibling?.focus();
        }
    } else {
        otp.value[index] = "";
    }
};

const handleBackspace = (index, event) => {
    if (event.key === "Backspace" && !otp.value[index] && index > 0) {
        otp.value[index - 1] = "";
        event.target.previousElementSibling?.focus();
    }
};

const handleOTPSend = async () => {
    if (isLoading.value) {
        return;
    }
    isLoading.value = true;
    const recaptchaResponse = grecaptcha.getResponse();
    if (recaptchaResponse) {
        await useFetch('/public/auth/send-otp', { method: "POST", body: { "email": email.value }, recaptchaResponse }).then((response) => {
            if (response.is_mail_sent) {
                isCodeSent.value = true
                emailOTP.value = email.value
                toast.show(i18n.global.t('login.otp.send_success'));
                start()
            } else {
                toast.show(response.error);
            }
        });
    } else {
        toast.show('Please complete the CAPTCHA');
    }
    isLoading.value = false;
};


const handleVerifyOTP = async () => {
    if (isLoading.value) {
        return;
    }
    isLoading.value = true;
    await useFetch('/public/auth/verify-otp', {
        method: "POST",
        body: {
            "email": emailOTP.value,
            "otp": otp.value.join(""),
            "device": getDevice()
        }
    }
    ).then((response) => {
        if (response.token) {
            toast.show(i18n.global.t('login.otp.verify_success'));
            localStorage.setItem("token", response.token)
            router.replace("/").then(() => router.go());
        } else {
            toast.show(response.error);
        }
    }).catch(e => {
        toast.show(e.error);
    });
    isLoading.value = false;
};


const recaptchaCallback = () => {
    const recaptchaResponse = grecaptcha.getResponse();
    isCaptchaVerified.value = !!recaptchaResponse;
};

onMounted(() => {
    getDevice()
    if (window.grecaptcha) {
        grecaptcha.render('recaptcha-container', {
            sitekey: recaptchaSiteKey,
            callback: recaptchaCallback,
        });
    } else {
        console.error("reCAPTCHA script is not loaded.");
    }
});
</script>