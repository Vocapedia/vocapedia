<template>
    <button class="flex items-center justify-center space-x-2" :class="buttonClasses" @click="handleClick">
        <mdicon v-if="icon" :name="icon" size="24" />
        <div v-if="text">
            <slot>{{ text }}</slot>
        </div>
    </button>
</template>

<script setup>
import { computed } from "vue";
const emit = defineEmits(["click"]);
const props = defineProps({
    theme: {
        type: String,
        default: "gray",
        validator: (val) => ["gray", "red", "blue"].includes(val),
    },
    text: {
        type: String,
        default: "",
    },
    icon: {
        type: String,
        default: "",
    },
});
const handleClick = (event) => {
   emit("click", event);
};

const buttonClasses = computed(() => {
    switch (props.theme) {
        case "red":
            return "smooth-click2 w-full flex-1 px-4 py-2 font-semibold text-xl bg-red-100 text-red-900 hover:bg-red-200 dark:hover:bg-red-600/20 dark:bg-red-700/20 dark:text-red-100 transition-colors rounded-sm";
        case "blue":
            return "smooth-click2 w-full flex-1 px-4 py-2 font-semibold text-xl bg-sky-100 text-sky-900 hover:bg-sky-200 dark:hover:bg-sky-600/20 dark:bg-sky-700/20 dark:text-sky-100 transition-colors rounded-sm";
        default:
            return "smooth-click2 w-full flex-1 px-4 py-2 font-semibold text-xl bg-gray-100 text-gray-900 hover:bg-gray-200 dark:hover:bg-zinc-600/20 dark:bg-zinc-700/20 dark:text-zinc-100 transition-colors rounded-sm";
    }
});
</script>
