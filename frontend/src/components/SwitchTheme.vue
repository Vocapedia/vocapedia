<template>
    <button @click="toggleDark()" class="cursor-pointer transition ease-in-out delay-50 hover:scale-105 duration-300">
        <mdicon name="white-balance-sunny" class="text-orange-400 dark:text-orange-300" />
    </button>
</template>

<script setup>
import { useDark, useToggle } from "@vueuse/core";

const isDark = useDark();
const toggleDark = useToggle(isDark);
</script>