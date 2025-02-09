<template>
    <header class="container mx-auto flex justify-between items-center p-4">
        <RouterLink to="/" class="flex-shrink-0">
            <img class="w-24 sm:w-16 md:w-20 lg:w-24" src="@/assets/logo.png" alt="logo">
        </RouterLink>

        <nav class="flex items-center space-x-4">
            <RouterLink to="/settings" class="smooth-click">
                <mdicon name="cog-outline" :size="iconSize" />
            </RouterLink>
            <RouterLink to="/account" class="smooth-click">
                <mdicon name="account-circle-outline" :size="iconSize" />
            </RouterLink>
            <SwitchThemeVue />
        </nav>
    </header>
</template>

<script setup>
import SwitchThemeVue from "@/components/SwitchTheme.vue";
import { ref, onMounted, onBeforeUnmount } from "vue";

const iconSize = ref(getSize());

function getSize() {
    const width = window.innerWidth;

    if (width < 640) return 24
    else return 30
}

const updateSize = () => {
    iconSize.value = getSize();
};

onMounted(() => {
    window.addEventListener("resize", updateSize);
});
onBeforeUnmount(() => {
    window.removeEventListener("resize", updateSize);
});
</script>