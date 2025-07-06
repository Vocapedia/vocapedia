<template>
    <div class="max-w-4xl mx-auto p-6 rounded-lg">
        <h1 class="text-3xl font-semibold text-center mb-6">{{ $t("privacy.title") }}</h1>

        <p class="text-lg mb-4">{{ $t("privacy.introduction") }}</p>

        <h2 class="text-2xl font-semibold mt-6 mb-2">{{ $t("privacy.dataCollection.title") }}</h2>
        <p class="text-base mb-4">{{ $t("privacy.dataCollection.text") }}</p>

        <h2 class="text-2xl font-semibold mt-6 mb-2">{{ $t("privacy.dataUsage.title") }}</h2>
        <p class="text-base mb-4">{{ $t("privacy.dataUsage.text") }}</p>

        <h2 class="text-2xl font-semibold mt-6 mb-2">{{ $t("privacy.thirdPartySharing.title") }}</h2>
        <p class="text-base mb-4">{{ $t("privacy.thirdPartySharing.text") }}</p>

        <h2 class="text-2xl font-semibold mt-6 mb-2">{{ $t("privacy.dataSecurity.title") }}</h2>
        <p class="text-base mb-4">{{ $t("privacy.dataSecurity.text") }}</p>

        <h2 class="text-2xl font-semibold mt-6 mb-2">{{ $t("privacy.userRights.title") }}</h2>
        <p class="text-base mb-4">{{ $t("privacy.userRights.text") }}</p>

        <h2 class="text-2xl font-semibold mt-6 mb-2">{{ $t("privacy.cookies.title") }}</h2>
        <p class="text-base mb-4">{{ $t("privacy.cookies.text") }}</p>

        <h2 class="text-2xl font-semibold mt-6 mb-2">{{ $t("privacy.dataStorage.title") }}</h2>
        <p class="text-base mb-4">{{ $t("privacy.dataStorage.text") }}</p>

        <h2 class="text-2xl font-semibold mt-6 mb-2">{{ $t("privacy.childrenPrivacy.title") }}</h2>
        <p class="text-base mb-4">{{ $t("privacy.childrenPrivacy.text") }}</p>

        <p class="text-sm text-center mt-6">{{ $t("privacy.lastUpdated") }} : {{ lastUpdatedFormatted }}</p>
    </div>
</template>
<script setup>
import dayjs from 'dayjs';
import 'dayjs/locale/de';
import 'dayjs/locale/en';
import 'dayjs/locale/es';
import 'dayjs/locale/fr';
import 'dayjs/locale/tr';
import 'dayjs/locale/zh';
import { GetLang } from '@/i18n/i18n';
const lang = GetLang()
dayjs.locale(lang);

const lastUpdated = "2025-04-06";
const lastUpdatedFormatted = dayjs(lastUpdated).format('D MMMM YYYY, dddd'); // Formatı yerelleştirilmiş olarak al

</script>