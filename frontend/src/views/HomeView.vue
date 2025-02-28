<template>
  <div class="space-y-16">
    <div>
      <h1 class="p-4 text-3xl font-bold">{{ $t('trends') }}</h1>
      <div class="flex justify-center">
        <div class="flex overflow-x-auto space-x-4 p-4 scrollbar-hide">
          <div v-for="i in 10" :key="i" @click="$router.push('/l/' + i)"
            class="min-w-[250px] md:min-w-[300px] smooth-click card h-50 flex justify-center items-center">
            <p class="text-xl p-3">trend / {{ i }}</p>
          </div>
        </div>
      </div>
    </div>

    <div>
      <h1 class="p-4 text-3xl font-bold">{{ $t('personal_info') }}</h1>
      <div class="container mx-auto px-4">
        <div class="grid grid-cols-12 gap-4">

          <router-link class="card rounded-xl transition-all duration-200 hover:scale-101 col-span-12 sm:col-span-6" to="/followed">
            <div class="bg-zinc-100 dark:bg-zinc-800 flex justify-center items-center text-xl font-bold p-4">
              <mdicon name="star-outline" />
              {{ $t('followed') }}
            </div>
          </router-link>

          <router-link class="card rounded-xl transition-all duration-200 hover:scale-101 col-span-12 sm:col-span-6" to="/notes">
            <div class="bg-zinc-100 dark:bg-zinc-800 flex justify-center items-center space-x-1 text-xl font-bold p-4">
              <mdicon name="text" />
              <span>{{ $t('notes') }}</span>
            </div>
          </router-link>

        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useStorage } from '@vueuse/core';
import fake_response from "@/fake/search_list.json";

const follows = fake_response.list
const notes = useStorage('notes', [])
</script>