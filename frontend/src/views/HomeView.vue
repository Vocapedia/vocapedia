<template>
  <div class="space-y-16">

    <div>
      <h1 class="p-4 text-3xl font-bold mb-5">{{ $t('personal_info') }}</h1>
      <div class="container mx-auto space-y-5">
        <div v-motion-slide-visible-once-bottom class="grid grid-cols-12 gap-4">

          <router-link class="smooth-click2 col-span-12 md:col-span-12 lg:col-span-6 relative" to="streamers">
            <div class="dark:bg-zinc-900 bg-zinc-50 card relative rounded-lg">
              <div class="bg-yellow-100/5 ">

                <Dots height="64" />
                <div class=" absolute inset-0 flex justify-center items-center text-center">
                  <span class="text-5xl font-semibold">
                    {{ $t('home.welcome') }}

                  </span>
                </div>
                <!--   <div class="bg-sky-100 dark:bg-sky-700 font-semibold text-lg text-center p-2">
                  {{ $t('home.experience.title') }}
                </div> -->
              </div>
            </div>
          </router-link>

          <div class="items-center grid grid-cols-12 gap-4 col-span-12 lg:col-span-6 sm:col-span-12">
            <router-link class="h-full col-span-12 sm:col-span-6 lg:col-span-12 smooth-click2" to="/followed">
              <div
                class="h-full rounded-lg bg-zinc-100 dark:bg-zinc-800 flex justify-center items-center text-xl font-bold p-4">
                <span> {{ $t('followed') }}
                </span>

              </div>
            </router-link>

            <router-link class="h-full col-span-12 sm:col-span-6 lg:col-span-12 smooth-click2" to="/notes">
              <div
                class="h-full rounded-lg bg-zinc-100 dark:bg-zinc-800 flex justify-center items-center space-x-1 text-xl font-bold p-4">
                <span>{{ $t('notes') }}</span>
              </div>
            </router-link>
          </div>
        </div>
      </div>
    </div>
    <div>
      <div v-auto-animate class="flex items-center">
        <h1 v-if="(response.trends ?? []).length > 0" class="p-4 text-3xl font-bold">{{ $t('trends') }}</h1>
        <transition class="py-5" name="fade">
          <mdicon v-if="!response.trends" spin name="loading" />
        </transition>
      </div>
      <transition class="py-5" name="fade" mode="out-in">
        <div v-motion-slide-visible-once-top v-if="(response.trends ?? []).length > 0" class="flex justify-center">
          <div class="flex overflow-x-auto space-x-4 p-4 scrollbar-hide">
            <div v-for="item in response.trends" :key="item.id" @click="$router.push('/l/' + BigInt(item.id))"
              class=" min-w-[250px] md:min-w-[300px] smooth-click card flex justify-center items-end bg-white dark:bg-zinc-800 rounded-lg">
              <div class="space-y-2 w-full">
                <div class="text-xl p-3">{{ item.title }}</div>
                <div class="text-sm p-3 h-18 text-gray-500 dark:text-gray-400 overflow-hidden text-ellipsis">
                  {{ item.description }}
                </div>
                <div class="bg-sky-100 dark:bg-sky-700 text-sm flex items-center justify-center space-x-2 p-1">
                  <span>{{ getLangByCode(item.lang).name }}</span>
                  <mdicon name="arrow-right" size="16" />
                  <span>{{ getLangByCode(item.target_lang).name }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </transition>
    </div>
  </div>
</template>

<script setup>
import Dots from '@/components/Dots.vue'

import { ref, onMounted } from 'vue';
import { useFetch } from '@/composable/useFetch';
import { getLangByCode } from '@/utils/language/languages';
const response = ref("{}")

onMounted(async () => {
  response.value = await useFetch("/public/chapters/trends")
})
</script>
