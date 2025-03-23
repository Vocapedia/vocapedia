<template>
    <div class="h-64 relative w-full  overflow-hidden">
        <div class="absolute top-0 left-0 w-full h-full">
            <div v-for="(dot, index) in dots" :key="index"
                class="absolute bg-yellow-400 rounded-full w-1 h-1 shadow-5xl" :style="getDotStyle(dot)"
                :class="dot.animationClass"></div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';

const dots = ref([]);
const props = defineProps(
    {
        count: {
            type: Number,
            default: 10,
        },
        height: {
            type: String,
            default: "64",
        },
    },
);

const generateRandomDot = () => {
    const top = Math.random() * 100;
    const left = Math.random() * 100;
    const duration = Math.random() * (8 - 5) + 5;
    const animationClass = `animate-dot-${Math.floor(Math.random() * 4)}`;

    return {
        top,
        left,
        duration,
        animationClass,
    };
};

onMounted(() => {
    for (let i = 0; i < props.count; i++) {
        dots.value.push(generateRandomDot());
    }
});

const getDotStyle = (dot) => {
    return {
        top: `${dot.top}%`,
        left: `${dot.left}%`,
        animationDuration: `${dot.duration}s`,
        boxShadow: `0 0 ${Math.random() * 10 + 5}px rgba(255, 255, 0)`,
    };
};

const style = document.createElement('style');
document.head.appendChild(style);

const keyframes = computed(() => {
    let styles = '';
    dots.value.forEach((dot, index) => {
        styles += `
        @keyframes moveDot-${index} {
          0% { transform: translate(0, 0); }
          25% { transform: translate(${Math.random() * 60}px, ${Math.random() * 60}px); }
          50% { transform: translate(${Math.random() * -60}px, ${Math.random() * 60}px); }
          75% { transform: translate(${Math.random() * 60}px, ${Math.random() * -60}px); }
          100% { transform: translate(0, 0); }
        }
  
        .animate-dot-${index} {
          animation: moveDot-${index} ${dot.duration}s infinite cubic-bezier(0.25, 0.8, 0.25, 1);
        }
      `;
    });
    return styles;
});

onMounted(() => {
    style.innerHTML = keyframes.value;
});
</script>

<style scoped></style>
