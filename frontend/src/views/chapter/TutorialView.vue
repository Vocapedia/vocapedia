<template>
    <div>
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