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
    <FloatDot />
  </div>

</template>
<script setup>
import FloatDot from '@/components/FloatDot.vue';
import { getUser } from '@/utils/token';

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
  }
]
</script>