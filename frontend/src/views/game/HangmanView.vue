<template>
    <div class="flex flex-col items-center p-8 space-y-8">
        <h1 class="text-3xl font-bold text-blue-500">{{ $t('game-hangman.title') }}</h1>
        <p v-motion-fade v-if="message" class="text-red-500">{{ message }}</p>
        <div class="text-4xl font-mono">
            <p v-auto-animate class="text-center">
                <transition-group name="slide-fade" tag="span">
                    <span v-auto-animate v-for="(letter, index) in displayWord" :key="index + letter"
                        class="mx-1 letter animate__animated animate__fadeInUp">
                        {{ letter }}
                    </span>
                </transition-group>
            </p>
        </div>
        <div>{{ displayDescription }}</div>
        <div v-auto-animate class="flex space-x-2">
            <input v-model="currentGuess" @keyup.enter="guessLetter" type="text" maxlength="1"
                class="p-2 w-full text-xl border border-gray-300 rounded-md" :disabled="gameOver"
                :placeholder="gameOver ? wordToGuess : $t('game-hangman.enterGuess')" />
            <div v-if="!gameOver" class="flex">
                <button @click="guessLetter"
                    class="smooth-click bg-green-100 hover:bg-green-300 dark:bg-green-950 dark:hover:bg-green-700 p-2.5 rounded-full">
                    <mdicon name="check" />
                </button>
            </div>
        </div>
        <div>
            <p class="text-lg font-semibold">{{ $t('game-hangman.attemptsLeft') }}:
                <span class="text-red-500">{{ attemptsLeft }}</span>
            </p>
        </div>
        <div v-motion-fade v-if="attemptsLeft < 6">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-48 w-48" viewBox="0 0 100 100">
                <rect x="10" y="10" width="10" height="80" :class="strokeColor" />
                <line x1="10" y1="10" x2="50" y2="10" :class="strokeColor" />
                <line x1="50" y1="10" x2="50" y2="30" :class="strokeColor" />
                <line x1="10" y1="10" x2="10" y2="90" :class="strokeColor" />
                <circle v-motion-roll-bottom v-if="attemptsLeft <= 5" cx="50" cy="40" r="10" :class="circleColor" />
                <line v-motion-roll-top v-if="attemptsLeft <= 4" x1="50" y1="50" x2="50" y2="70" :class="strokeColor" />
                <line v-motion-roll-right v-if="attemptsLeft <= 3" x1="50" y1="55" x2="30" y2="60"
                    :class="strokeColor" />
                <line v-motion-roll-left v-if="attemptsLeft <= 2" x1="50" y1="55" x2="70" y2="60"
                    :class="strokeColor" />
                <line v-motion-roll-right v-if="attemptsLeft <= 1" x1="50" y1="70" x2="30" y2="80"
                    :class="strokeColor" />
                <line v-motion-roll-left v-if="attemptsLeft <= 0" x1="50" y1="70" x2="70" y2="80"
                    :class="strokeColor" />
            </svg>
        </div>
        <div v-if="gameOver" class="text-center mt-6">
            <h2 class="text-2xl font-bold">{{ gameResult }}</h2>
            <button @click="restartGame" class="mt-4 bg-blue-500 text-white py-2 px-6 rounded-md hover:bg-blue-600">{{
                $t('game-hangman.restart') }}</button>
        </div>
    </div>
</template>
<script setup>
import { ref, computed, onMounted } from 'vue';
import { useDark } from "@vueuse/core";
import { useI18n } from 'vue-i18n';
import { useFetch } from '@/composable/useFetch';
import { useRoute } from 'vue-router';
const route = useRoute()
const { t } = useI18n();

const isDark = useDark();
const wordToGuess = ref("");
const displayWord = ref(Array(wordToGuess.value.length).fill("_"));
const displayDescription = ref("");
const currentGuess = ref("");
const attemptsLeft = ref(6);
const guessedLetters = ref([]);
const message = ref("");
const gameOver = ref(false);
const gameResult = ref("");

onMounted(async () => {
    await restartGame()
})

const guessLetter = () => {
    if (gameOver.value) return;

    const guess = currentGuess.value.toLowerCase();

    if (!guess || guessedLetters.value.includes(guess)) {
        message.value = t('game-hangman.alreadyGuessed');
        return;
    }

    guessedLetters.value.push(guess);

    if (wordToGuess.value.toLowerCase().includes(guess)) {
        updateDisplayWord(guess);
        if (!displayWord.value.includes("_")) {
            gameOver.value = true;
            gameResult.value = t('game-hangman.youWon');
        }
    } else {
        attemptsLeft.value--;
        if (attemptsLeft.value === 0) {
            gameOver.value = true;
            gameResult.value = `${t('game-hangman.youLost')} ${t('game-hangman.theWordWas')}: ${wordToGuess.value}`;
        }
    }

    currentGuess.value = "";
    message.value = "";
};

const updateDisplayWord = (letter) => {
    for (let i = 0; i < wordToGuess.value.length; i++) {
        if (wordToGuess.value[i].toLowerCase() === letter) {
            displayWord.value[i] = letter;
        }
    }
};

const restartGame = async () => {
    await useFetch("/public/chapters/hangman/" + route.params.id).then((r) => {
        wordToGuess.value = r.word
        displayWord.value = Array(wordToGuess.value?.length).fill("_")
        displayDescription.value = r.description
    })
    currentGuess.value = "";
    attemptsLeft.value = 6;
    guessedLetters.value = [];
    message.value = "";
    gameOver.value = false;
    gameResult.value = "";
};

const strokeColor = computed(() => (isDark.value ? 'stroke-white' : 'stroke-black'));
const circleColor = computed(() => (isDark.value ? 'fill-white stroke-white' : 'fill-none stroke-black'));
</script>
<style scoped>
.error {
    color: red;
}

body.dark .stroke-white {
    stroke: white;
}

body.dark .fill-white {
    fill: white;
}

.letter {
    display: inline-block;
}
</style>