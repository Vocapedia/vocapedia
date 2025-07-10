<template>
    <div class="chapter-tutorial">
        <div class="max-w-160 mx-auto prose dark:prose-invert">
            <div v-html="response.chapter.tutorial" />
        </div>
    </div>
</template>

<script setup>

const props = defineProps({
    response: {
        type: Object,
        required: true,
    }
})
</script>

<style scoped>
.chapter-tutorial {
    max-width: 160rem;
    margin: 0 auto;
}
</style>
