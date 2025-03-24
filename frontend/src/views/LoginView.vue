<template>
    <div class="pt-36 flex items-center justify-center ">
        <Shine />
        <div class="z-1 w-full max-w-120">
            <div class="bg-white/90 dark:bg-zinc-800/90 p-8  rounded-lg shadow-lg">
                <div class="text-2xl text-center pb-4">
                    <RouterLink to="/" class="flex-shrink-0 py-2">
                        <div class="font-logo">
                            Vocapedia
                        </div>
                    </RouterLink>
                </div>
                <form @submit.prevent="handleSubmit">
                    <div class="mb-4">
                        <label for="email" class="block text-sm font-medium">{{ $t('login.input.label') }}</label>
                        <input type="email" id="email" v-model="email" required class="w-full p-3 border shadow-sm outline-none transition-all 
             bg-zinc-100 text-zinc-900 py-2 mt-3 border-none rounded-full
             dark:bg-zinc-900 dark:text-white" :placeholder="$t('login.input.placeholder')" />
                    </div>

                    <div class="mb-4" id="recaptcha-container"></div>

                    <div class="text-center">
                        <button type="submit" :disabled="isLoading"
                            :class="isCaptchaVerified ? 'cursor-pointer transition duration-200 hover:scale-101 active:scale-99 bg-sky-700 focus:ring-sky-700 ' : 'bg-zinc-200 dark:bg-zinc-900'"
                            class="flex justify-between items-center py-2 px-4 rounded-lg w-full focus:outline-none">
                            <span>
                                <mdicon :class="isCaptchaVerified ? 'text-white dark:text-zinc-900' : ''"
                                    v-if="isLoading" name="loading" spin />
                            </span>
                            <span :class="isCaptchaVerified ? 'text-white dark:text-zinc-900' : ''"
                                class="font-semibold text-lg">{{ $t('login.button') }}</span>
                            <span>
                                <mdicon v-if="isLoading" name="loading" spin
                                    :class="isCaptchaVerified ? 'text-sky-700' : 'text-zinc-200 dark:text-zinc-900'" />
                            </span>
                        </button>
                    </div>
                </form>
            </div>
        </div>

    </div>

</template>

<script setup>
import Shine from '@/components/Shine.vue';
import { useFetch } from '@/composable/useFetch';
import { i18n } from '@/i18n/i18n';
import { ref, onMounted } from 'vue';

const isLoading = ref(false);
const email = ref('');
const isCaptchaVerified = ref(false);
const recaptchaSiteKey = import.meta.env.VITE_RECAPTCHA_SITE_KEY;

const handleSubmit = async () => {
    if (isLoading.value) {
        return;
    }
    isLoading.value = true;
    const recaptchaResponse = grecaptcha.getResponse();
    if (recaptchaResponse) {
        await useFetch('/auth/login', { method: "POST", body: { "email": email.value }, recaptchaResponse }).then((response) => {
            console.log(response);
            if (response.is_mail_sent) {
                alert(i18n.global.t('login.success'));
            } else {
                alert(response.error);
            }
        });
    } else {
        alert('Please complete the CAPTCHA');
    }
    isLoading.value = false;
};


const recaptchaCallback = () => {
    const recaptchaResponse = grecaptcha.getResponse();
    isCaptchaVerified.value = !!recaptchaResponse;
};

onMounted(() => {
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