<template>
    <div class="max-w-5xl mx-auto space-y-5">
        <div class="flex justify-between items-center">
            <router-link :to="'/l/' + $route.params.id"
                class="smooth-click bg-blue-100 dark:bg-blue-900 px-2 py-1 rounded-full">
                <span class="flex justify-between items-center space-x-2">
                    <mdicon name="arrow-left" />
                    <span class="text-xl font-semibold">{{ response.title }}</span>
                </span>
            </router-link>
        </div>
        <div>
            <swiper @swiper="onSwiper" :scrollbar="{ draggable: false }" :modules="modules" :slides-per-view="1">
                <swiper-slide v-for="test in 10 ">
                    <div class="card p-5 rounded-xl my-5">
                        <div class="text-center text-4xl font-bold capitalize pb-10">Kelime</div>
                        <div class="grid grid-cols-12 gap-4 py-5">
                            <button @click="answer(test)"
                                class="capitalize text-center col-span-12 md:col-span-6 lg:col-span-3 card rounded-full px-4 py-2 smooth-click"
                                v-for="i in 4">Cevap</button>
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
    </div>
</template>

<script setup>
import { ref } from 'vue';
import fake_response from "@/fake/list.json";
import { Swiper, SwiperSlide } from 'swiper/vue';
import { Navigation, Pagination, Scrollbar, A11y } from 'swiper/modules';
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