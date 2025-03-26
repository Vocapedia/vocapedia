<template>
  <div class="space-y-16 pb-16">
    <div>
      <div class="flex items-center">
        <h1 class="p-4 text-3xl font-bold">{{ $t('trends') }}</h1>
        <transition name="fade" mode="out-in">
          <mdicon v-if="!response.trends" spin name="loading" />
        </transition>
      </div>
      <div v-auto-animate>
        <div v-if="response.trends" class="flex justify-center">
          <div class="flex overflow-x-auto space-x-4 p-4 scrollbar-hide">
            <div v-for="item in response.trends" :key="item.id" @click="$router.push('/l/' + BigInt(item.id))"
              class=" min-w-[250px] md:min-w-[300px] smooth-click card flex justify-center items-end bg-white dark:bg-zinc-800 rounded-lg">
              <div class="space-y-2 w-full">
                <div class="text-xl p-3">{{ item.title }}</div>
                <div class="text-sm p-3 h-18 text-gray-500 dark:text-gray-400 overflow-hidden text-ellipsis">
                  {{ item.description }}
                </div>
                <div class="bg-sky-100 dark:bg-sky-700 text-sm flex items-center justify-center space-x-2">
                  <span>{{ item.lang }}</span>
                  <mdicon name="arrow-right" size="16" />
                  <span>{{ item.target_lang }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div>
      <h1 class="p-4 text-3xl font-bold mb-5">{{ $t('personal_info') }}</h1>
      <div class="container mx-auto space-y-5">
        <div class="grid grid-cols-12 gap-4">
          <div class="col-span-12 md:col-span-12 lg:col-span-6  relative">

            <div class="dark:bg-zinc-900 bg-zinc-50 card relative rounded-lg">
              <div class="bg-yellow-100/5 ">

                <Dots height="64" />
                <div class="absolute inset-0 flex justify-center items-center text-center">
                  <span v-splitting class="text-5xl font-semibold">
                    654
                  </span>
                </div>

                <div class="bg-sky-100 dark:bg-sky-700 font-semibold text-lg text-center p-2">
                  {{ $t('home.experience.title') }}
                </div>
              </div>
            </div>
          </div>
          <div class="items-center grid grid-cols-12 gap-4 col-span-12 lg:col-span-6 sm:col-span-12">
            <router-link
              class="h-full col-span-12 sm:col-span-6 lg:col-span-12 transition-all duration-200 hover:scale-101"
              to="/followed">
              <div
                class="h-full rounded-lg bg-zinc-100 dark:bg-zinc-800 flex justify-center items-center text-xl font-bold p-4">
                <mdicon name="star-outline" />
                <span> {{ $t('followed') }}
                </span>

              </div>
            </router-link>

            <router-link
              class="h-full col-span-12 sm:col-span-6 lg:col-span-12 transition-all duration-200 hover:scale-101"
              to="/notes">
              <div
                class="h-full rounded-lg bg-zinc-100 dark:bg-zinc-800 flex justify-center items-center space-x-1 text-xl font-bold p-4">
                <mdicon name="text" />
                <span>{{ $t('notes') }}</span>
              </div>
            </router-link>

          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import Dots from '@/components/Dots.vue'

import { ref, onMounted } from 'vue';
import { useFetch } from '@/composable/useFetch';
const response = ref("{}")

onMounted(async () => {
  response.value = await useFetch("/chapters/trends")
})
</script>
