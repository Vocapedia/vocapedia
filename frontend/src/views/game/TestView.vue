<template>
    <div class="max-w-5xl mx-auto space-y-5">
        <div>
            <swiper @swiper="onSwiper" :scrollbar="{ draggable: false }" :modules="modules" :slides-per-view="1">
                <swiper-slide v-for="test in response">
                    <div class="card p-5 rounded-xl my-5">
                        <!--   <div class="text-end">
                            <button @click="speak(test.word)">
                                <mdicon name="volume-high"/>
                            </button>
                        </div> -->
                        <div class="text-center text-4xl font-bold capitalize pb-2">{{ test.word }}</div>
                        <div class="text-center text-sm capitalize pb-10">{{ test.description }} </div>
                        <div class="grid grid-cols-12 gap-4 py-5">
                            <button @click="answer(i)"
                                class="capitalize text-center col-span-12 md:col-span-6 lg:col-span-3 card rounded-full px-4 py-2 smooth-click"
                                v-for="i in test.choices">
                                {{ i }}
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
                    <button @click="ShowPoint"
                        class="flex justify-center smooth-click bg-sky-100 dark:bg-sky-700 px-3 py-2 space-x-1 rounded-full">
                        <span class="sm:text-xl font-semibold">
                            {{ $t('show_result') }}
                        </span>
                    </button>
                </div>
            </div>
        </transition>
        <ShowGrade v-model="triggerComposePopup">
            <template #header>
                <h2 v-if="progress != 100" class="text-xl font-semibold" :class="progressClass">
                    Hoppp! buralara dikkat et
                </h2>
            </template>
            <template #description>
                <div class="space-y-5">
                    <h2 class="text-5xl font-semibold" :class="progressClass">
                        %{{ progress.toFixed(1) }}
                    </h2>

                    <div class="w-full bg-gray-200 dark:bg-gray-600 rounded-full h-1">
                        <div class="progress-bar bg-blue-500 h-1 rounded-full" :style="{ width: progress + '%' }"></div>
                    </div>

                </div>

                <div>
                    <router-link v-for="item in wrongAnswers"
                        :to="'/l/' + $route.params.id + '#' + item.word.replace(/\s+/g, '-')" :key="item.word">
                        <div class="smooth-click2 p-2">
                            <h4>{{ item.word }}</h4>
                            <small>{{ item.choices[item.correctIndex] }}</small>
                        </div>
                    </router-link>
                </div>
            </template>
            <template #buttons>
                <div class="flex justify-around">
                    <button @click="triggerComposePopup = false" :class="buttonClass"
                        class="smooth-click2 cursor-pointer w-full px-4 py-2 font-semibold text-xl rounded">
                        {{ $t('close') }}
                    </button>
                </div>
            </template>
        </ShowGrade>
    </div>
</template>

<script setup>
import ShowGrade from "@/components/shared/Popup.vue"
import { ref, computed, onMounted, watch, nextTick } from 'vue';
import { Swiper, SwiperSlide } from 'swiper/vue';
import { Navigation, Pagination, Scrollbar, A11y } from 'swiper/modules';
import { useFetch } from "@/composable/useFetch";
import { useRoute, useRouter } from "vue-router";
import { useToast } from "@/composable/useToast";
import { i18n } from "@/i18n/i18n";
import { speak } from "@/utils/textToSpeech";

const toast = useToast()
const route = useRoute()
const router = useRouter()

const progress = ref(5)
const canShowResult = ref(false)
const wrongAnswers = ref([])
const triggerComposePopup = ref(false)
const swipers = ref();
const result = ref([]);
const modules = ref([Navigation, Pagination, Scrollbar, A11y]);
const response = ref([]);

onMounted(async () => {
    var listLang
    var listTargetLang
    var wordCount = 0
    await useFetch("/public/chapters/" + route.params.id).then(r => {
        listLang = r.chapter.lang
        listTargetLang = r.chapter.target_lang
        wordCount = r.chapter.word_count
    })
    if (wordCount < 4) {
        toast.show(i18n.global.t('games.test.error.not_enough_wrong'))
        router.replace('/l/' + route.params.id + '/games')
        return
    }
    await useFetch(`/public/chapters/game-format/${route.params.id}?lang=${listLang}&target_lang=${listTargetLang}`).then(r => {
        response.value = r.questions
    })
})

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

const buttonClass = computed(() => {
    if (progress.value < 33) {
        return "bg-red-100 text-red-900 dark:bg-red-900 dark:text-red-100"

    }
    else if (progress.value < 66) {
        return "bg-yellow-100 text-yellow-900 dark:bg-yellow-900 dark:text-yellow-100"

    }
    else {
        return "bg-green-100 text-green-900 dark:bg-green-900 dark:text-green-100"

    }
})




watch(result.value, (n, o) => {
    if (result.value.length == response.value.length) {
        canShowResult.value = true
    }
})

function ShowPoint() {
    var correctCount = 0
    wrongAnswers.value = []
    response.value.forEach((element, i) => {
        if (element.choices[element.correctIndex] == result.value[i]) {
            correctCount += 1
        } else {
            wrongAnswers.value.push(element)
        }
    });
    progress.value = (correctCount / response.value.length) * 100
    nextTick(() => {
        triggerComposePopup.value = true
    })
}

function onSwiper(theSwiper) {
    swipers.value = theSwiper;
};

function answer(ev) {
    result.value[swipers.value.activeIndex] = ev
    swipers.value.slideNext()
}
</script>