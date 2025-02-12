<template>
    <div>
        <NavbarVue />
        <div class="animate__animated animate__fadeIn container mx-auto p-4">
            <transition name="fade" mode="out-in">
                <slot />
            </transition>

        </div>
    </div>

</template>
<script setup>
import NavbarVue from '@/components/Navbar.vue'
</script>