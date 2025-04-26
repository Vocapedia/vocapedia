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
        {{ $t('home.quick_access') }}
      </h2>
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
        <router-link :to="card.uri" v-for="card in cards" :key="card.name"
          class="p-4 bg-white dark:bg-zinc-800 rounded-lg">
          <div>
            <h5 class="flex items-center justify-between space-x-2">
              <span class="text-lg font-semibold">
                {{ $t(card.name) }}
              </span>
              <mdicon :name="card.icon" :class="card.icon_class" />

            </h5>

            <p class="text-sm text-gray-500 dark:text-gray-400">
              {{ $t(card.description) }}
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
              <h3 class="text-2xl font-medium">{{ $t('home.daily_streak') }}</h3>
              <p class="text-sm text-gray-500">{{ $t('home.daily_streak_description') }}</p>
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
              <span class="text-xs text-gray-500">{{ day }}</span>
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


const weekdays = Array.from({ length: 7 }, (_, i) =>
  new Intl.DateTimeFormat(GetLang(), { weekday: 'short' }).format(new Date(1970, 0, 5 + i))
);
const completedDays = ref([])
const rewarded = ref(false)
const streakCount = ref(0)
onMounted(async () => {
  if (!getUser()) {
    return
  }
  const response = await useFetch('/user/streak')
  streakCount.value = response.streak?.count || 0;
  if (response.streak?.rewarded) {
    confetti({
      particleCount: 200,
      angle: 90,
      spread: 150,
      origin: { x: 0.5, y: 0.5 }
    })
  }
  completedDays.value = Array.from({ length: 7 }, (_, i) => i < streakCount.value);
  rewarded.value = response.streak?.rewarded || false

})

const cards = [
  {
    uri: '/trends',
    name: 'trends',
    icon: 'trending-up',
    description: 'home.trends_description',
    icon_class: 'text-green-600 dark:text-green-500'
  }, {
    uri: '/notes',
    name: 'notes',
    icon: 'note-text-outline',
    description: 'home.notes_description',
    icon_class: 'text-yellow-600 dark:text-yellow-500'

  }, {
    uri: '/followed',
    name: 'followed_results',
    description: 'home.followed_results_description',
    icon: 'format-list-text',
    icon_class: 'text-sky-600 dark:text-sky-500'
  },
  {
    uri: '/streamers',
    name: 'home.streamers',
    description: 'home.streamers_description',
    icon: 'volume-high',
    icon_class: 'text-orange-600 dark:text-orange-500'
  }
]
</script>