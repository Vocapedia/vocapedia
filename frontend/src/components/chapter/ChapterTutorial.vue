<template>
    <div class="chapter-tutorial">
        <ChapterTutorialContent :tutorial-content="response.chapter.tutorial" />
    </div>
</template>

<script setup>
import ChapterTutorialContent from './ChapterTutorialContent.vue'

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
