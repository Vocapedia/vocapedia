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
                            <input type="text" inputmode="numeric" :placeholder="$t('login.input.otp_placeholder')"
                                maxlength="6"
                                class="w-full h-10 text-center text-xl border rounded-md outline-none border-zinc-200 bg-zinc-100 dark:border-zinc-800 dark:bg-zinc-700"
                                v-model="otp" />
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
                                class="font-semibold text-lg">{{ $t('login.button.otp.verify') }}</span>
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

                    <div ref="turnstileRef"></div>

                    <div class="mb-4">
                        <div v-if="captchaLoadError" class="mt-2 text-sm text-red-600 dark:text-red-400">
                            {{ captchaLoadError }}
                            <button @click="renderTurnstile" type="button"
                                class="ml-2 text-sky-700 dark:text-sky-500 hover:underline">
                                {{ $t('login.captcha.retry_button') }}
                            </button>
                        </div>
                    </div>

                    <div class="text-center">
                        <button type="submit" :disabled="isLoading || !!captchaLoadError"
                            :class="isCaptchaVerified ? 'smooth-click2 bg-sky-700 focus:ring-sky-700 ' : 'bg-zinc-200 dark:bg-zinc-900'"
                            class="flex justify-between items-center py-2 px-4 rounded-lg w-full focus:outline-none">
                            <span>
                                <mdicon :class="isCaptchaVerified ? 'text-white dark:text-zinc-900' : ''"
                                    v-if="isLoading" name="loading" spin />
                            </span>
                            <span :class="isCaptchaVerified && !captchaLoadError ? 'text-white dark:text-zinc-900' : ''"
                                class="font-semibold text-lg">{{ $t('login.button.otp.send') }}</span>
                            <span>
                                <mdicon v-if="isLoading" name="loading" spin
                                    :class="isCaptchaVerified && !captchaLoadError ? 'text-sky-700' : 'text-zinc-200 dark:text-zinc-900'" />
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

const otp = ref('');
const isCodeSent = ref(false);
const isLoading = ref(false);
const email = ref('');
const emailOTP = ref('');
const isCaptchaVerified = ref(false);
const captchaLoadError = ref(null); // For Turnstile load error
const turnstileToken = ref(null); // Store Turnstile token
const turnstileWidgetId = ref(null); // Store Turnstile widget ID
const router = useRouter();
const turnstileSiteKey = import.meta.env.VITE_TURNSTILE_SITE_KEY;
const toast = useToast();
import { useCountdown } from '@vueuse/core';
import { getDevice } from '@/utils/token';

const duration = 5 * 60;
const { remaining, start } = useCountdown(duration, {
    onComplete() {
        router.go(0)
    },
    onTick() {

    }
})

const handleOTPSend = async () => {
    if (isLoading.value) {
        return;
    }
    isLoading.value = true;
    if (turnstileToken.value) {
        await useFetch('/public/auth/send-otp', {
            method: "POST",
            body: { "email": email.value },
            turnstileToken: turnstileToken.value
        }).then((response) => {
            if (response.is_mail_sent) {
                isCodeSent.value = true;
                emailOTP.value = email.value;
                toast.show(i18n.global.t('login.otp.send_success'));
                start();
            } else {
                if (response.error === 'INVALID_EMAIL_FORMAT') {
                    toast.show(i18n.global.t('login.otp.send_error_invalid_email'));
                } else if (response.error === 'OTP_LIMIT_EXCEEDED') {
                    toast.show(i18n.global.t('login.otp.send_error_limit_exceeded'));
                } else {
                    toast.show(i18n.global.t('login.otp.send_error_generic'));
                }
            }
        }).catch(error => {
            console.error("Send OTP Error:", error);
            if (error.name === 'AbortError') {
                toast.show(i18n.global.t('login.error.timeout'));
            } else {
                toast.show(i18n.global.t('login.otp.send_error_network'));
            }
        }).finally(() => {
            isLoading.value = false;
        });
    } else {
        toast.show(i18n.global.t('login.captcha.missing'));
        isLoading.value = false;
    }
};


const handleVerifyOTP = async () => {
    if (isLoading.value) {
        return;
    }
    isLoading.value = true;
    try {
        const response = await useFetch('/public/auth/verify-otp', {
            method: "POST",
            body: {
                "email": emailOTP.value,
                "otp": otp.value,
                "device": getDevice()
            }
        });

        if (response.token) {
            toast.show(i18n.global.t('login.button.otp.verify_success'));
            localStorage.setItem("token", response.token);
            router.replace("/").then(() => router.go());
        } else {
            // Specific error handling for verify OTP
            if (response.error === 'INVALID_OTP') { // Example error code
                toast.show(i18n.global.t('login.otp.verify_error_invalid_otp'));
            } else if (response.error === 'OTP_EXPIRED') { // Example error code
                toast.show(i18n.global.t('login.otp.verify_error_expired'));
            } else if (response.error === 'TOO_MANY_ATTEMPTS') { // Example error code
                toast.show(i18n.global.t('login.otp.verify_error_too_many_attempts'));
            } else {
                toast.show(i18n.global.t('login.otp.verify_error_generic'));
            }
        }
    } catch (e) {
        console.error("Verify OTP Error:", e);
        if (e.name === 'AbortError') {
            toast.show(i18n.global.t('login.error.timeout'));
        } else {
            toast.show(i18n.global.t('login.otp.verify_error_network'));
        }
    } finally {
        isLoading.value = false;
    }
};



const turnstileCallback = (token) => {
    turnstileToken.value = token;
    isCaptchaVerified.value = !!token;
};


onMounted(() => {
    getDevice();
    renderTurnstile();
});


const renderTurnstile = () => {
    captchaLoadError.value = null;
    isCaptchaVerified.value = false;
    turnstileToken.value = null;
    // Remove previous widget if exists
    if (turnstileWidgetId.value && window.turnstile) {
        window.turnstile.remove(turnstileWidgetId.value);
        turnstileWidgetId.value = null;
    }
    // Load Turnstile script if not loaded
    if (!window.turnstile) {
        const script = document.createElement('script');
        script.src = 'https://challenges.cloudflare.com/turnstile/v0/api.js';
        script.async = true;
        script.defer = true;
        script.onload = () => {
            mountTurnstile();
        };
        script.onerror = () => {
            captchaLoadError.value = i18n.global.t('login.captcha.load_error_retry');
        };
        document.head.appendChild(script);
    } else {
        mountTurnstile();
    }
};

const mountTurnstile = () => {
    try {
        const el = (turnstileRef.value instanceof HTMLElement) ? turnstileRef.value : turnstileRef.value?.$el;
        if (el) {
            el.innerHTML = '';
            turnstileWidgetId.value = window.turnstile.render(el, {
                sitekey: turnstileSiteKey,
                callback: turnstileCallback,
                'error-callback': () => {
                    captchaLoadError.value = i18n.global.t('login.captcha.load_error_retry');
                    isCaptchaVerified.value = false;
                },
                theme: 'auto',
            });
        }
    } catch (error) {
        captchaLoadError.value = i18n.global.t('login.captcha.load_error_retry');
        isCaptchaVerified.value = false;
    }
};

import { onUnmounted } from 'vue';
const turnstileRef = ref(null);
onUnmounted(() => {
    if (turnstileWidgetId.value && window.turnstile) {
        window.turnstile.remove(turnstileWidgetId.value);
    }
});
</script>