<template>


    <div class="space-y-5 card p-3 rounded-xl">
        <div>
            <div class="flex-1">
                <div class="text-lg font-semibold text-gray-900 dark:text-white mb-2">
                    {{ form.title }}
                </div>
                <div class="text-sm text-gray-600 dark:text-gray-400 leading-relaxed">
                    {{ form.description }}
                </div>
            </div>
        </div>
        
        <div class="space-y-3">
            <div class="flex items-center justify-between text-sm">
                <div class="flex items-center space-x-2 text-gray-600 dark:text-gray-400">
                    <mdicon name="account-group" size="16" />
                    <span>{{ form.max_participants }} {{ $t('stream.maxStudents') }}</span>
                </div>
                <div class="flex items-center space-x-2 text-gray-600 dark:text-gray-400">
                    <mdicon name="timer" size="16" />
                    <span>{{ form.duration }}{{ $t('stream.min') }}</span>
                </div>
            </div>

            <div class="flex items-center justify-between text-sm">
                <div class="flex items-center space-x-2 text-gray-600 dark:text-gray-400">
                    <mdicon name="translate" size="16" />
                    <span>
                        {{ getLangByCode(form.lang).name }} → {{ getLangByCode(form.target_lang).name }}
                    </span>

                </div>
                <div class="text-xs text-gray-500 dark:text-gray-400">
                    {{ dayjs(form.scheduled_at).format('MMM D, YYYY H:mm') }}
                </div>
            </div>
            <div class="flex items-center justify-end text-sm">
                <router-link :to="'/' + username" class="flex items-center space-x-2 text-gray-600 dark:text-gray-400">
                    <mdicon name="school" size="16" />
                    <span>{{ username }}</span>
                </router-link>
            </div>
        </div>
    </div>
</template>

<script setup>
import { GetLang } from '@/i18n/i18n';
import { getLangByCode } from '@/utils/language/languages';
import dayjs from 'dayjs';
import 'dayjs/locale/de';
import 'dayjs/locale/en';
import 'dayjs/locale/es';
import 'dayjs/locale/fr';
import 'dayjs/locale/tr';
import 'dayjs/locale/zh';
const lang = GetLang()
dayjs.locale(lang);

defineProps({
    form: {
        title: String,
        description: String,
        lang: String,
        target_lang: String,
        scheduled_at: String,
        duration: Number,
        max_participants: Number
    },
    username: {
        type: String,
        default: ''
    }
})
</script>
