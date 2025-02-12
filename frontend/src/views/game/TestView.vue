<template>
    <div class="max-w-5xl mx-auto space-y-5">
        <div class="flex justify-between items-center">
            <router-link :to="'/l/' + $route.params.id"
                class="smooth-click dark:text-blue-100 text-blue-900 px-2 py-1 rounded-full">
                <small class="flex justify-between items-center space-x-2">
                    <mdicon name="arrow-left" size="15" />
                    <span class="sm:text-base font-semibold">{{ response.title }}</span>
                </small>
            </router-link>
        </div>
        <div>
            <swiper @slideChange="onSlideChange" @swiper="onSwiper" :scrollbar="{ draggable: false }" :modules="modules"
                :slides-per-view="1">
                <swiper-slide v-for="test in 10 ">
                    <div class="card p-5 rounded-xl my-5">
                        <div class="text-center text-4xl font-bold capitalize pb-10">Kelime</div>
                        <div class="grid grid-cols-12 gap-4 py-5">
                            <button @click="answer(test)"
                                class="capitalize text-center col-span-12 md:col-span-6 lg:col-span-3 card rounded-full px-4 py-2 smooth-click"
                                v-for="i in 4">
                                Cevap
                            </button>
                        </div>
                    </div>
                </swiper-slide>
            </swiper>
            <div class="flex justify-center items-center space-x-5 py-5">
                <button class="smooth-click" @click="swipers.slidePrev()">
                    <mdicon name="arrow-left" size="32" />
                </button>
                <button class="smooth-click" @click="swipers.slideNext()">
                    <mdicon name="arrow-right" size="32" />
                </button>
            </div>
        </div>
        <transition name="slide-fade" mode="out-in">
            <div v-if="canShowResult" class="text-center">
                <div class="flex justify-center items-center">
                    <button @click="triggerComposePopup = true"
                        class="flex justify-center items-center smooth-click bg-yellow-100 dark:bg-yellow-900 p-2 space-x-1 rounded-full">
                        <mdicon name="certificate-outline" size="20" />
                        <span class="sm:text-xl font-semibold">
                            {{ $t('show_result') }}
                        </span>
                    </button>
                </div>
            </div>
        </transition>
        <ShowGrade v-model="triggerComposePopup">
            <template #header>
                <h2 class="text-xl font-semibold text-red-600 dark:text-red-500">
                    Hoppp! buralara dikkat et
                </h2>
            </template>
            <template #description>
                <div>
                    <h2 class="text-5xl font-semibold" :class="progressClass">
                        %{{ progress }}
                    </h2>

                    <div class="w-full bg-gray-200 dark:bg-gray-600 rounded-full h-1">
                        <div class="progress-bar bg-blue-500 h-1 rounded-full" :style="{ width: progress + '%' }"></div>
                    </div>

                </div>

                <div>
                    <router-link v-for="item in response.list" :to="'/l/' + $route.params.id + '#' + item.word"
                        :key="item.word">
                        <div class="smooth-click p-2">
                            <h4>{{ item.word }}</h4>
                            <small>{{ item.languages[0].word }}</small>
                        </div>
                    </router-link>
                </div>
            </template>
            <template #buttons>
                <div class="flex justify-around">
                    <button @click="triggerComposePopup = false"
                        class="cursor-pointer bg-red-500 w-full px-4 py-2 text-white rounded">
                        {{ $t('close') }}
                    </button>
                </div>
            </template>
        </ShowGrade>
    </div>
</template>

<script setup>
import ShowGrade from "@/components/Popup.vue"
import { ref, computed, watch } from 'vue';
import fake_response from "@/fake/list.json";
import { Swiper, SwiperSlide } from 'swiper/vue';
import { Navigation, Pagination, Scrollbar, A11y } from 'swiper/modules';
const progress = ref(5)
const canShowResult = ref(false)
const progressClass = computed(() => {
    if (progress.value < 33) {
        return "text-red-500"

    }
    else if (progress.value < 66) {
        return "text-yellow-500"

    }
    else {
        return "text-green-500"

    }
})
function onSlideChange(newSlide) {
    if (newSlide.activeIndex == 9) {
        canShowResult.value = true
    }
}
const triggerComposePopup = ref(false)
const swipers = ref();
const result = ref([]);
const modules = ref([Navigation, Pagination, Scrollbar, A11y]);
const onSwiper = (theSwiper) => {
    swipers.value = theSwiper;
};

function answer(ev) {
    result.value[swipers.value.activeIndex] = ev
    swipers.value.slideNext()
}

const response = ref(fake_response);
</script>