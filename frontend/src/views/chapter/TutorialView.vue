<template>
    <div class="max-w-160 mx-auto">
        <div>{{ $t('empty') }}</div>
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