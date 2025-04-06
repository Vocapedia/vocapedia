<template>
    <div class="space-y-16">
        <section class="hero bg-blue-500 dark:bg-sky-900 text-white text-center py-16">
            <div class="container mx-auto">
                <h1 class="text-4xl font-bold hero-text">
                    {{ $t('landing.heroTitle') }}
                </h1>
                <p class="text-xl mt-4">{{ $t('landing.heroSubtitle') }}</p>
            </div>
        </section>

        <section class="features">
            <div class="container mx-auto">
                <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-5 px-5">
                    <div v-for="i in 3"
                        class="feature-card bg-white dark:bg-zinc-900 p-6 rounded-lg card hover:bg-zinc-100 dark:hover:bg-zinc-800 transition">
                        <h3 class="text-xl font-semibold">{{ $t(`landing.feature${i}Title`) }}</h3>
                        <p class="mt-4">{{ $t(`landing.feature${i}Description`) }}</p>
                    </div>
                </div>
            </div>
        </section>

        <section class="chart-section">
            <div class="container mx-auto text-center">
                <h2 class="text-3xl font-semibold">{{ $t('landing.chartTitle') }}</h2>
                <div class="chart-container w-full h-96 flex justify-center mt-8">
                    <Line :options="chartOptions" :data="chartData" ref="chartRef" />

                </div>
                <div class="flex justify-evenly">
                    <div class="flex items-center space-x-2">
                        <mdicon name="circle" style="color: rgb(75, 192, 192);" />
                        <span>{{ $t('landing.vocapedia_users') }}</span>
                    </div>
                    <div class="flex items-center space-x-2">
                        <mdicon name="circle" style="color: rgb(255, 99, 132);" />
                        <span>{{ $t('landing.traditional_methods') }}</span>
                    </div>
                </div>
            </div>
        </section>

        <section class="swiper-section">
            <swiper loop :scrollbar="{ draggable: true }" :autoplay="{
                delay: 5000,
                disableOnInteraction: false,
                pauseOnMouseEnter: false,
                stopOnLastSlide: false,
                waitForTransition: true
            }" :modules="modules" :pagination="{
                dynamicBullets: true,
            }" navigation :slides-per-view="1">
                <swiper-slide>
                    <div class="card rounded-xl m-5 mx-16">
                        <section class="transition hover:bg-zinc-100 dark:hover:bg-zinc-800 py-16">
                            <div class="container mx-auto text-center">
                                <h2 class="text-3xl font-semibold">{{ $t('landing.securityTitle') }}</h2>
                                <p class="mt-4 p-5">{{ $t('landing.securityDescription') }}</p>
                            </div>
                        </section>
                    </div>
                </swiper-slide>

                <swiper-slide>
                    <div class="card rounded-xl m-5 mx-16">
                        <section class="transition hover:bg-zinc-100 dark:hover:bg-zinc-800 py-16">
                            <div class="container mx-auto text-center">
                                <h2 class="text-3xl font-semibold">{{ $t('landing.chromeExtensionTitle') }}</h2>
                                <p class="mt-4 p-5">{{ $t('landing.chromeExtensionDescription') }}</p>
                            </div>
                        </section>
                    </div>
                </swiper-slide>

                <swiper-slide>
                    <div class="card rounded-xl m-5 mx-16">
                        <section class="transition hover:bg-zinc-100 dark:hover:bg-zinc-800 py-16">
                            <div class="container mx-auto text-center">
                                <h2 class="text-3xl font-semibold">{{ $t('landing.whyLearnTitle') }}</h2>
                                <p class="mt-4 p-5">{{ $t('landing.whyLearnDescription') }}</p>
                            </div>
                        </section>
                    </div>
                </swiper-slide>
            </swiper>
        </section>

        <footer class="py-4">
            <div class="space-x-2 flex">
                <router-link to="/privacy-policy" class="underline">
                    Privacy Policy
                </router-link>
            </div>
        </footer>
    </div>
</template>
<script setup>
import { onMounted, ref } from 'vue'
import { gsap } from 'gsap'
import { Line } from 'vue-chartjs'
import { Chart as ChartJS, LinearScale, LineElement, CategoryScale, LineController, PointElement, Filler } from 'chart.js'
import { Swiper, SwiperSlide } from 'swiper/vue';
import { Autoplay, Parallax, EffectFade, Navigation, Thumbs, Scrollbar } from 'swiper/modules';

ChartJS.register(LinearScale, LineElement, CategoryScale, LineController, PointElement, Filler)
const modules = [Autoplay, Parallax, EffectFade, Navigation, Thumbs, Scrollbar]
const chartRef = ref(null)

onMounted(() => {
    const heroText = document.querySelector('.hero-text');
    const featureSection = document.querySelector('.features');
    const chartContainer = document.querySelector('.chart-container');
    const swiperSection = document.querySelector('.swiper-section');

    const observerOptions = {
        root: null,
        threshold: 0.75,
    };

    const animateOnVisibility = (target, animation) => {
        const observer = new IntersectionObserver((entries) => {
            entries.forEach(entry => {
                if (entry.isIntersecting) {
                    animation(entry.target);
                    observer.unobserve(entry.target);
                }
            });
        }, observerOptions);

        observer.observe(target);
    };

    animateOnVisibility(heroText, (target) => {
        gsap.from(target, { opacity: 0, x: -100, duration: 1 });
        gsap.to(target, { opacity: 1, x: 0, duration: 1, delay: 0.5 });
    });

    animateOnVisibility(featureSection, (target) => {
        gsap.from(target, { opacity: 0, y: 50, stagger: 0.3, duration: 1 });
        gsap.to(target, { opacity: 1, y: 0, stagger: 0.3, duration: 1 });
    });

    animateOnVisibility(swiperSection, (target) => {
        gsap.from(target, { opacity: 0, y: 50, stagger: 0.3, duration: 1 });
        gsap.to(target, { opacity: 1, y: 0, stagger: 0.3, duration: 1 });
    });

    animateOnVisibility(chartContainer, (target) => {
        gsap.from(target, { opacity: 0, y: 50, stagger: 0.3, duration: 1 });
        gsap.to(target, { opacity: 1, y: 0, stagger: 0.3, duration: 1 });
    });
});
const chartData = {
    labels: ['', '', '', '', '', ''],
    datasets: [
        {
            data: [10, 25, 50, 80, 120, 160],
            borderColor: 'rgb(75, 192, 192)',
            backgroundColor: 'rgba(75, 192, 192, 0.2)',
            fill: true,
            tension: 0.4,
            pointRadius: 0,
            cubicInterpolationMode: 'monotone',
        },
        {
            data: [5, 10, 20, 30, 45, 60],
            borderColor: 'rgb(255, 99, 132)',
            backgroundColor: 'rgba(255, 99, 132, 0.2)',
            fill: true,
            tension: 0.4,
            pointRadius: 0.4,
            cubicInterpolationMode: 'monotone',
        }
    ]
}

const chartOptions = {
    responsive: true,
    maintainAspectRatio: true,
    animation: true,
    scales: {
        x: {
            type: 'category',
            ticks: {
                display: true,
            },
            grid: {
                display: false,
            },
        },
        y: {
            type: 'linear',
            beginAtZero: true,
            ticks: {
                display: false,
            },
            grid: {
                display: false,
            },
        },
    },
}
</script>
<style scoped>
.hero-text {
    opacity: 0;
}

.features {
    opacity: 0;
    transform: translateY(50px);
}

.chart-container {
    opacity: 0;
    transform: translateY(50px);
}

.swiper-section {
    opacity: 0;
    transform: translateY(50px);
}
</style>
