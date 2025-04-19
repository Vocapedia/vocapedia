<template>
    <div class="max-w-160 mx-auto">
        <p class="py-5 text-center text-xl font-semibold">
            {{ $t('trends') }}
        </p>


        <WordLists uri="/public/chapters/trends" />
    </div>
</template>
<script setup>
import WordLists from '@/components/WordLists.vue';
</script>