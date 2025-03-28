<template>
    <div class="max-w-160 mx-auto prose dark:prose-invert">
        <div v-html="response.chapter.tutorial" />
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