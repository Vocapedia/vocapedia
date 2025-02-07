<template>
    <div class="max-w-6xl mx-auto ">
        <div class="flex justify-between items-center animate__animated animate__fadeInDown">
            <nav>
                <RouterLink to="/">
                    <img width="100" src="@/assets/logo.png" alt="">
                </RouterLink>
            </nav>
            <nav class="flex items-center space-x-4 p-4">
                <RouterLink to="/account" class="smooth-click">
                    <mdicon name="account-circle-outline" size="32" />
                </RouterLink>
                <SwitchThemeVue />
            </nav>

        </div>
    </div>
</template>
<script setup>
import SwitchThemeVue from './SwitchTheme.vue';

</script>