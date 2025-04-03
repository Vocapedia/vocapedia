<template>
    <div class="max-w-160 mx-auto space-y-10">
        <div>
            <div class="flex items-center space-x-5">
                <div class="flex items-center space-x-2 w-full">
                    <select v-model="selectedLanguage"
                        class="w-full px-4 py-2 bg-zinc-100 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-xl focus:border-zinc-300 dark:focus:border-zinc-600">
                        <option v-for="lang in languages" :value="lang.code"> {{ lang.name.valueOf() }}</option>
                    </select>
                </div>
                <div>
                    <button @click="SaveLang" class="smooth-click rounded-full bg-sky-100 dark:bg-sky-700 p-2">
                        <mdicon name="content-save" />
                    </button>
                </div>
            </div>
        </div>

        <div class="overflow-x-auto rounded-xl">
            <table class="w-full border border-zinc-200 dark:border-zinc-800 bg-white dark:bg-zinc-900 shadow-md">
                <thead class="bg-sky-100 dark:bg-sky-700  text-gray-900 dark:text-white">
                    <tr>
                        <th class="w-10 px-4 py-2 text-left">{{ $t('settings.created_at') }}</th>
                        <th class="px-4 py-2 text-left">{{ $t('settings.updated_at') }}</th>
                        <th class="px-4 py-2 text-left">{{ $t('settings.device_info') }}</th>
                        <th class="px-4 py-2 text-left"></th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(token, index) in tokenInfo" :key="index"
                        class="transition border-t dark:border-zinc-800 border-zinc-200 hover:bg-zinc-100 dark:hover:bg-zinc-800">
                        <td class="w-45 px-4 py-2">{{ dayjs(token.created_at).format("YYYY MM DD HH:mm:ss") }}</td>
                        <td class="w-45 px-4 py-2">{{ dayjs(token.updated_at).format("YYYY-MM-DD HH:mm:ss") }}</td>
                        <td class="w-36 px-4 py-2">
                            <p v-for="d in token.device">{{ d }}</p>
                        </td>
                        <td class="px-4 py-2 text-center">
                            <button @click="deleteToken(BigInt(token.id))"
                                class="smooth-click bg-red-100 hover:bg-red-300 dark:bg-red-950 dark:hover:bg-red-700  p-1.5 rounded-full">
                                <mdicon name="close" />
                            </button>
                        </td>

                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>

<script setup>
import { useFetch } from '@/composable/useFetch';
import { ref, onMounted } from 'vue';
import { ChangeLang } from '../i18n/i18n';
import { useRouter } from "vue-router";
import dayjs from "dayjs";
const router = useRouter()

// Available languages (example)
const languages = [
    { code: 'en', name: 'English' },
    { code: 'tr', name: 'Türkçe' },
    { code: 'es', name: 'Español' },
    { code: 'fr', name: 'Français' }
];

// Currently selected language (default: English)
const selectedLanguage = ref('en');
function SaveLang() {
    ChangeLang(selectedLanguage.value.toString())
    router.go()
}
function deleteToken(tokenID) {
    console.log(BigInt(tokenID))
}
// Token Info
const tokenInfo = ref([]);

// Fetch current token information and selected language
const fetchTokenInfo = async () => {
    await useFetch('/auth/token').then(r => {
        tokenInfo.value = r.tokens
    });
};

onMounted(() => {
    fetchTokenInfo();
});
</script>