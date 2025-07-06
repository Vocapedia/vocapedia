<template>
    <div class="p-4 max-w-md mx-auto space-y-10">
        <div v-auto-animate v-if="currentWord">
            <p class="text-lg mb-2">
                {{ $t('game-flipword.question', {
                    from: sourceLang.name,
                    word: sourceWord?.word,
                    to: targetLang.name
                }) }}
            </p>

            <p class="text-sm text-gray-600 mb-2">
                {{ $t('game-flipword.time_left') }}: {{ timeLeft }}s
            </p>

            <input v-model="userGuess" @keyup.enter="checkGuess" type="text" class="w-full p-2 border rounded mb-2"
                :placeholder="$t('game-flipword.placeholder')" :disabled="isTimeUp" v-auto-animate />

            <button @click="checkGuess" class="bg-blue-500 text-white px-4 py-2 rounded" :disabled="isTimeUp">
                {{ $t('game-flipword.submit') }}
            </button>

            <p v-if="feedback" class="mt-2 font-medium">
                {{ feedback }}
            </p>

            <p v-if="isTimeUp" class="mt-2 text-sm text-gray-500">
                {{ timeUpMessage }} {{ timeUpRemaining }}s
            </p>
        </div>

        <div v-auto-animate v-else>
            <p>{{ $t('game-flipword.no_more_words') }}</p>
        </div>
    </div>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue'
import { useCountdown } from '@vueuse/core'
import { i18n } from '@/i18n/i18n';
import { getLangByCode } from '@/utils/language/languages';

const sourceLang = getLangByCode('en')
const targetLang = getLangByCode('tr')

const wordIndex = ref(0)
const userGuess = ref('')
const feedback = ref('')

const wordBases = ref([])

const currentWordBase = computed(() => wordBases.value[wordIndex.value])
const { remaining: timeUpRemaining, start: timeUpstart } = useCountdown(5, {
    onComplete() {
        nextWord()
    },
})

const sourceWord = computed(() =>
    currentWordBase.value?.words.find(w => w.lang === sourceLang.code)
)
const targetWord = computed(() =>
    currentWordBase.value?.words.find(w => w.lang === targetLang.code)
)

const currentWord = computed(() => sourceWord.value && targetWord.value)

const { remaining: timeLeft, start: startTimer, stop: stopTimer } = useCountdown(10, {
    onComplete() {
        showTimeUpMessage()
    }
})

watch(wordIndex, () => {
    feedback.value = ''
    startTimer()
})

async function fetchWords() {
    wordBases.value = []
    startTimer()
}

function checkGuess() {
    if (!targetWord.value || isTimeUp.value) return

    const guess = userGuess.value.trim().toLowerCase()
    const correct = targetWord.value.word.toLowerCase()

    if (guess === correct) {
        feedback.value = '✅ ' + i18n.global.t('game-flipword.correct')
        stopTimer()
        setTimeout(nextWord, 1000)
    } else {
        feedback.value = `${userGuess.value} ❌`
    }
}

function nextWord() {
    userGuess.value = ''
    feedback.value = ''
    wordIndex.value++
}

const isTimeUp = computed(() => timeLeft.value <= 0)

function showTimeUpMessage() {
    feedback.value = `${i18n.global.t('game-flipword.time_up')} ${targetWord.value?.word}`
    timeUpMessage.value = `${i18n.global.t('game-flipword.time_up')} ${targetWord.value?.word}. | `
    timeUpstart()
}

const timeUpMessage = ref('')

onMounted(() => {
    fetchWords()
})
</script>
