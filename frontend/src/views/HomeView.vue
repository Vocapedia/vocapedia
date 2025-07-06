<template>
  <div class="px-2 max-w-4xl mx-auto space-y-10">
    <div v-motion-slide-visible-once-top class="space-y-2">
      <div class="text-5xl font-semibold">
        {{ $t('home.welcome', { username: getUser().username ? `, ${getUser().username}` : '' }) }}
      </div>

      <div class="text-lg font-light">
        {{ $t('home.description') }}
      </div>
    </div>
    <div v-motion-slide-visible-once-bottom class="space-y-4">
      <h2 class="text-2xl font-semibold">
        {{ $t('home.quick_access.title') }}
      </h2>
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
        <router-link :to="card.uri" v-for="card in cards" :key="card.name"
          class="p-4 bg-white dark:bg-zinc-800 rounded-lg">
          <div>
            <h5 class="flex items-center justify-between space-x-2">
              <span class="text-lg font-semibold">
                {{ $t("home.quick_access.cards." + card.name) }}
              </span>
              <mdicon :name="card.icon" :class="card.icon_class" />

            </h5>

            <p class="text-sm text-gray-500 dark:text-gray-400">
              {{ $t("home.quick_access.cards." + card.description) }}
            </p>
          </div>
        </router-link>

      </div>
    </div>
    <hr class="border-t-2 border-zinc-200 dark:border-zinc-700 opacity-50">
    <div @click="() => { !getUser() ? $router.push('/login') : '' }">
      <div v-motion-slide-visible-once-left class="rounded-lg">
        <div class="pb-2">
          <div class="flex items-center justify-between">
            <div>
              <h3 class="text-2xl font-medium">{{ $t('home.daily_streak.title') }}</h3>
              <p class="text-sm text-gray-500">{{ $t('home.daily_streak.description') }}</p>
            </div>
            <div
              class="flex items-center gap-1 rounded-full bg-orange-100 dark:bg-orange-500 px-2 py-1 text-orange-600 dark:text-white">
              <mdicon name="fire" />
              <span class="text-sm font-medium">{{ streakCount }}</span>
            </div>
          </div>
        </div>
        <div class="p-6">
          <div class="flex justify-between pt-2">
            <div v-for="(day, i) in weekdays" :key="i" class="flex flex-col items-center gap-2">
              <span class="text-xs text-gray-500">{{ dayjs(day).format("ddd") }}</span>
              <div class="flex h-8 w-8 items-center justify-center rounded-full"
                :class="completedDays[i] ? 'bg-orange-500 text-white' : 'border border-dashed border-gray-300'">
                <mdicon name="fire" v-if="completedDays[i]" />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <FloatDot />

  </div>

</template>
<script setup>
import FloatDot from '@/components/FloatDot.vue';
import { useFetch } from '@/composable/useFetch';
import { GetLang } from '@/i18n/i18n';
import { getUser } from '@/utils/token';
import { onMounted, ref } from 'vue';
import confetti from 'canvas-confetti'
import dayjs from 'dayjs';
import 'dayjs/locale/de';
import 'dayjs/locale/en';
import 'dayjs/locale/es';
import 'dayjs/locale/fr';
import 'dayjs/locale/tr';
import 'dayjs/locale/zh';
const lang = GetLang()
dayjs.locale(lang);

const weekdays = ref([])
const completedDays = ref([])
const rewarded = ref(false)
const streakCount = ref(0)
onMounted(async () => {
  if (!getUser()) {
    return
  }
  const response = await useFetch('/user/streak')
  streakCount.value = response.streak?.count || 0;
  const mainDate = new Date(response.streak.lastDate);

  for (let i = streakCount.value - 1; i >= 0; i--) {
    const date = new Date(mainDate);
    date.setDate(date.getDate() - i);
    weekdays.value.push(date)
  }

  for (let i = 1; i < (8 - streakCount.value); i++) {
    const date = new Date(mainDate);
    date.setDate(date.getDate() + i);
    weekdays.value.push(date)
  }

  for (let i = streakCount.value; i > 0; i--) {
    completedDays.value[i - 1] = true
  }
  if (response.streak?.rewarded) {
    confetti({
      particleCount: 200,
      angle: 90,
      spread: 150,
      origin: { x: 0.5, y: 0.5 }
    })
  }
  rewarded.value = response.streak?.rewarded || false

})

const cards = [
  {
    uri: '/trends',
    name: 'trends.title',
    icon: 'trending-up',
    description: 'trends.description',
    icon_class: 'text-green-600 dark:text-green-500'
  }, {
    uri: '/notes',
    name: 'notes.title',
    icon: 'note-text-outline',
    description: 'notes.description',
    icon_class: 'text-yellow-600 dark:text-yellow-500'

  }, {
    uri: '/followed',
    name: 'followed_results.title',
    description: 'followed_results.description',
    icon: 'format-list-text',
    icon_class: 'text-sky-600 dark:text-sky-500'
  },
  {
    uri: '/streamers',
    name: 'streamers.title',
    description: 'streamers.description',
    icon: 'volume-high',
    icon_class: 'text-orange-600 dark:text-orange-500'
  },
  {
    uri: '/pricing',
    name: 'pricing.title',
    description: 'pricing.description',
    icon: 'currency-usd',
    icon_class: 'text-purple-600 dark:text-purple-500'
  }
]
</script>