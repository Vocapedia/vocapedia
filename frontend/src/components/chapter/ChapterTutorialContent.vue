<template>
    <div class="max-w-160 mx-auto prose dark:prose-invert">
        <div v-html="tutorialContent" />
    </div>
</template>

<script setup>
const props = defineProps({
    tutorialContent: {
        type: String,
        required: true,
    }
})
</script>
