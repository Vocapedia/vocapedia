<template>
    <div class="max-w-160 mx-auto space-y-10">
        <div class="space-y-5">
            <div class="flex space-x-5">
                <div v-motion-slide-visible-once-top class="flex items-center space-x-5 w-full">
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
                <SwitchThemeVue />
            </div>
            <div v-motion-slide-visible-once-left v-if="vocatoken" class="flex w-full items-center space-x-5">
                <button @click="isVocatokenShowing = !isVocatokenShowing"
                    class="smooth-click rounded-full bg-sky-100 dark:bg-sky-700 p-2">
                    <mdicon v-if="isVocatokenShowing" name="eye-off-outline" />
                    <mdicon v-else name="eye-outline" />
                </button>
                <input placeholder="VocaToken" disabled :value="'VOCATOKEN-' + vocatoken"
                    :type="isVocatokenShowing ? '' : 'password'"
                    class="w-full px-4 py-2 bg-zinc-100 dark:bg-zinc-800 border border-zinc-200 dark:border-zinc-700 rounded-xl focus:border-zinc-300 dark:focus:border-zinc-600" />
                <button @click="CopyVocatoken" class="smooth-click rounded-full bg-sky-100 dark:bg-sky-700 p-2">
                    <mdicon name="content-copy" />
                </button>
                <button @click="UpdateVocatoken" class="smooth-click rounded-full bg-sky-100 dark:bg-sky-700 p-2">
                    <mdicon name="reload" />
                </button>
            </div>
        </div>
        <div v-motion-slide-visible-once-bottom v-if="(tokenInfo ?? []).length > 0"
            class="overflow-x-auto rounded-xl space-y-5">
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
                    <tr v-for="(tkn, index) in tokenInfo" :key="index"
                        class="transition border-t dark:border-zinc-800 border-zinc-200 hover:bg-zinc-100 dark:hover:bg-zinc-800">
                        <td class="w-45 px-4 py-2">{{ dayjs(tkn.created_at).format("YYYY MM DD HH:mm:ss") }}</td>
                        <td class="w-45 px-4 py-2">{{ dayjs(tkn.updated_at).format("YYYY MM DD HH:mm:ss") }}</td>
                        <td class="w-36 px-4 py-2">
                            <p v-for="d in tkn.device">{{ d }}</p>
                        </td>
                        <td class="px-4 py-2 text-center">
                            <button v-if="tkn.token != token" @click="deleteToken(BigInt(tkn.id))"
                                class="smooth-click bg-red-100 hover:bg-red-300 dark:bg-red-950 dark:hover:bg-red-700  p-1.5 rounded-full">
                                <mdicon name="close" />
                            </button>
                            <button v-else
                                class="smooth-click bg-green-100 hover:bg-green-300 dark:bg-green-950 dark:hover:bg-green-700 p-1.5 rounded-full">
                                <mdicon name="check" />
                            </button>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>

        <div v-motion-slide-visible-once-right v-if="(tokenInfo ?? []).length > 0" class="flex justify-between">
            <div class="flex items-center justify-end space-x-5">
                <span>{{ $t('settings.delete_account') }}</span>
                <button @click="logout" class="smooth-click rounded-full bg-red-100 dark:bg-red-700 p-2">
                    <mdicon name="account-remove-outline" />
                </button>
            </div>
            <div class="flex items-center justify-end space-x-5">
                <span>{{ $t('settings.logout') }}</span>
                <button @click="logout" class="smooth-click rounded-full bg-red-100 dark:bg-red-700 p-2">
                    <mdicon name="logout" />
                </button>
            </div>
        </div>
    </div>
</template>

<script setup>
import SwitchThemeVue from "@/components/SwitchTheme.vue";
import { useFetch } from '@/composable/useFetch';
import { ref, onMounted } from 'vue';
import { ChangeLang, GetLang, i18n } from '../i18n/i18n';
import { useRouter } from "vue-router";
import dayjs from "dayjs";
import { getLangByCode } from '@/utils/language/languages';
import { useToast } from '@/composable/useToast';
const router = useRouter()
const token = localStorage.getItem("token")
const modules = import.meta.glob("../i18n/*.json", { eager: true });

const isVocatokenShowing = ref(false)
const toast = useToast();
const languages = Object.entries(modules).map(([path, module]) => {
    const code = path.replace("../i18n/", "").replace(".json", "");
    const language = getLangByCode(code);

    return { code, name: language ? language.name : code }; // Eğer dil ismi varsa onu al, yoksa kodu döndür
});

const selectedLanguage = ref(GetLang().slice(0, 2));
function SaveLang() {
    ChangeLang(selectedLanguage.value.toString())
    router.go()
}
async function deleteToken(tokenID) {
    await useFetch("/user/token/" + tokenID, {
        method: "DELETE"
    }).then(
        tokenInfo.value = tokenInfo.value.filter(e => e.id != tokenID)
    )
}
const tokenInfo = ref([]);

const fetchTokenInfo = async () => {
    await useFetch('/user/token').then(r => {
        tokenInfo.value = r.tokens
    });
    await useFetch('/user/vocatoken', {
        method: "GET"
    }).then(r => {
        vocatoken.value = r.vocatoken
    });

};

onMounted(() => {
    if (token) {
        fetchTokenInfo();
    }
});
const vocatoken = ref("");
async function UpdateVocatoken() {
    await useFetch('/user/vocatoken', {
        method: "PUT"
    }).then(r => {
        vocatoken.value = r.vocatoken
    });
}
async function CopyVocatoken() {
    await navigator.clipboard.writeText("VOCATOKEN-" + vocatoken.value);
    toast.show(i18n.global.t('settings.vocatoken_copied'))
}
async function logout() {
    await useFetch("/auth/logout", {
        method: "DELETE"
    })
    localStorage.removeItem("token");
    router.replace("/login");
}
</script>