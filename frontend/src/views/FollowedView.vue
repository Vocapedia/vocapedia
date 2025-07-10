<template>
    <div class="max-w-160 mx-auto">
        <p class="py-5 text-center text-xl font-semibold">
            {{ $t('followeds.header') }}
        </p>

        <div v-motion-slide-bottom>
            <WordLists uri="/chapters/favorite" />
        </div>
    </div>
</template>

<script setup>
import WordLists from '@/components/WordLists.vue';
</script>