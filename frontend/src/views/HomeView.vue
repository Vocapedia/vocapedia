<template>
  <div class="space-y-16 pb-16">
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
      <h1 class="p-4 text-3xl font-bold">{{ $t('home_chart_title') }}</h1>
      <div class="container mx-auto space-y-5">

        <div class="space-y-5">
          <Bar class="max-w-160 w-full mx-auto" id="my-chart-id" :options="chartOptions"
            :title="{ color: isDark ? 'red' : 'blue' }" :data="chartData" />
          <div class="text-center font-semibold text-5xl">
            9535
          </div>
        </div>
      </div>
    </div>


    <div>
      <h1 class="p-4 text-3xl font-bold">{{ $t('personal_info') }}</h1>
      <div class="container mx-auto space-y-5">
        <div class="grid grid-cols-12 gap-4">
          <router-link class="col-span-12 sm:col-span-6 transition-all duration-200 hover:scale-101" to="/followed">
            <div class="bg-zinc-100 dark:bg-zinc-800 flex justify-center items-center text-xl font-bold p-4">
              <mdicon name="star-outline" />
              <span> {{ $t('followed') }}
              </span>

            </div>
          </router-link>

          <router-link class="col-span-12 sm:col-span-6 transition-all duration-200 hover:scale-101" to="/notes">
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
import { useDark } from "@vueuse/core";

const isDark = useDark();

import { Bar } from 'vue-chartjs'
import { Chart as ChartJS, Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale } from 'chart.js'
import { i18n } from '@/i18n/i18n';
ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale)
const chartData = {
  labels: ['January', 'February', 'March'],
  datasets: [{
    label: ' ',
    borderWidth: 2,
    borderRadius: 5,
    backgroundColor: '#0ea5e950',
    borderColor: '#0ea5e9',
    borderSkipped: false,
    barThickness: 25,
    data: [40, -20, 12],
    color: '#0ea5e9',
  }]
}
const chartOptions = {
  responsive: true,
  plugins: {
    legend: {
      position: '',
    },
  },
}

const follows = fake_response.list
const notes = useStorage('notes', [])
</script>