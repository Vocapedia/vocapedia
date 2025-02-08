<template>
    <div>
        <NavbarVue />
        <div class="animate__animated animate__fadeInUp container mx-auto p-4">
            <slot />
        </div>
    </div>
</template>
<script setup>
import NavbarVue from '@/components/Navbar.vue'
</script>