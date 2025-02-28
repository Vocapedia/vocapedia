<template>
    <div class="max-w-160 mx-auto">
        <p class="py-5">
            {{ $t('followed_results', { count: response.list.length, }) }}
        </p>
        <WordLists :response="response" />
    </div>
</template>

<script setup>
import fake_response from "@/fake/search_list.json";
import WordLists from '@/components/WordLists.vue';
const response = fake_response


</script>