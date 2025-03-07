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

    <div ref="captureDiv" class="p-1 animated-gradient">
      <div class="dark:bg-zinc-900 bg-zinc-50 px-5 py-2">
        <div class="flex justify-between items-center">
          <h1 class="p-4 text-3xl font-bold">{{ $t('home_chart_title') }}</h1>
          <button v-if="showShareButton" @click="shareImage"
            class="smooth-click p-2 bg-zinc-100 dark:bg-zinc-800 rounded-full">
            <mdicon name="share-variant-outline" />
          </button>
        </div>
        <div class="container mx-auto space-y-5">
          <div class="space-y-5">
            <swiper @slideChange="onSlideChange" @swiper="onSwiper" :modules="modules" :slides-per-view="1">
              <swiper-slide>
                <Bar v-if="activeIndex === 0" class="max-w-160 sm:max-w-140 w-full mx-auto" ref="bar"
                  :options="chartOptions" :data="chartData" />
                <div class="text-center font-semibold text-5xl">
                  243
                </div>
              </swiper-slide>

              <swiper-slide>
                <Line v-if="activeIndex === 1" class="max-w-160 sm:max-w-140 w-full mx-auto" ref="line"
                  :options="chartOptions" :data="chartData2" />
                <div class="text-center font-semibold text-5xl">
                  638
                </div>
              </swiper-slide>
            </swiper>
            <div v-if="showShareButton" class="flex justify-center items-center space-x-5 py-2">
              <button class="smooth-click" @click="swipers.slidePrev()">
                <mdicon name="arrow-left" size="32" />
              </button>
              <button class="smooth-click" @click="swipers.slideNext()">
                <mdicon name="arrow-right" size="32" />
              </button>
            </div>
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
import { ref, nextTick } from 'vue'
import html2canvas from 'html2canvas-pro'
import { Bar, Line } from 'vue-chartjs'
import { Chart as ChartJS, Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale, PointElement, LineElement, Filler } from 'chart.js'
import { i18n } from '@/i18n/i18n';
import { Swiper, SwiperSlide } from 'swiper/vue';
import { Navigation, Pagination, Scrollbar, A11y } from 'swiper/modules';
const modules = ref([Navigation, Pagination, Scrollbar, A11y]);
const swipers = ref();
const onSwiper = (theSwiper) => {
  swipers.value = theSwiper;
};
const activeIndex = ref(0)
const onSlideChange = (swiper) => {
  activeIndex.value = swiper.activeIndex; // Aktif slaytı güncelle
};

const captureDiv = ref(null)
const showShareButton = ref(true)
const shareImage = async () => {
  if (!navigator.canShare) {
    alert(i18n.global.t('share_not_supported'))
    return
  }
  showShareButton.value = false
  nextTick(async () => {
    const canvas = await html2canvas(captureDiv.value)
    const blob = await new Promise(resolve => canvas.toBlob(resolve, 'image/png'))
    const file = new File([blob], i18n.global.t('scoreboard') + ".png", { type: 'image/png' })

    if (navigator.canShare({ files: [file] })) {
      await navigator.share({
        title: "Puan Tablosu",
        text: "İşte puan tablom! 🚀",
        files: [file]
      })
    } else {
      alert(i18n.global.t('share_not_supported'))
    }
  })
  nextTick(() => {
    showShareButton.value = true
  })
}


ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale, PointElement, LineElement, Filler)
const chartData = {
  labels: ['Akif', 'John', 'Mike'],
  datasets: [{
    label: ' ',
    borderWidth: 2,
    borderRadius: 5,
    backgroundColor: '#0ea5e950',
    borderColor: '#0ea5e9',
    borderSkipped: false,
    barThickness: 25,
    data: [40, -20, 12],
  }]
}
const chartData2 = {
  labels: ['Ocak', 'Şubat', 'Mart', 'Nisan', 'Mayıs', 'Haziran'],
  datasets: [{
    label: ' ',
    backgroundColor: '#0ea5e950',
    borderColor: '#0ea5e9',
    data: [40, -20, 50, 60, -6, 35],
    fill: true,
  }]
}
const delayed = ref(false)
const chartOptions = {
  responsive: true,
  plugins: {
    legend: {
      position: '',
    },
  },
  animation: {
    onComplete: () => {
      delayed.value = true;
    },
    delay: (context) => {
      let delay = 0;
      if (context.type === 'data' && context.mode === 'default' && !delayed.value) {
        delay = context.dataIndex * 300 + context.datasetIndex * 100;
      }
      return delay;
    },
  },
}
</script>

<style scoped>
.animated-gradient {
  background: linear-gradient(45deg, #007BFF, #FFD700, #FF6F00);
  background-size: 200% 200%;
  animation: gradientAnimation 4s infinite alternate ease-in-out;
}

@keyframes gradientAnimation {
  0% {
    background-position: 0% 50%;
  }

  50% {
    background-position: 100% 50%;
  }

  100% {
    background-position: 0% 50%;
  }
}
</style>