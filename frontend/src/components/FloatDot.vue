<template>
    <div :style="dotStyle"
        class="fixed w-3 h-3 bg-yellow-600/50 dark:bg-yellow-400/50 rounded-full z-[9999] pointer-events-none"></div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const x = ref(window.innerWidth / 2)
const y = ref(window.innerHeight / 2)
const targetX = ref(x.value)
const targetY = ref(y.value)
const speed = 0.03

const dotStyle = ref({})

function updateStyle() {
    dotStyle.value = {
        top: `${y.value}px`,
        left: `${x.value}px`,
    }
}

function setNewTarget() {
    targetX.value = Math.random() * (window.innerWidth - 20)
    targetY.value = Math.random() * (window.innerHeight - 20)
}

function animate() {
    const dx = targetX.value - x.value
    const dy = targetY.value - y.value
    x.value += dx * speed
    y.value += dy * speed

    if (Math.abs(dx) < 2 && Math.abs(dy) < 2) {
        setNewTarget()
    }

    updateStyle()
    requestAnimationFrame(animate)
}

onMounted(() => {
    setNewTarget()
    animate()
})
</script>