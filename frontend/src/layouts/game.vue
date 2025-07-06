<template>
    <div>
        <NavbarVue />
        <div class="animate__animated animate__fadeIn container mx-auto p-4">
            <div class="flex justify-between items-center">
                <router-link :to="'/l/' + $route.params.id + '/games'"
                    class="smooth-click dark:text-blue-100 text-blue-900 px-2 py-1 rounded-full">
                    <small class="flex justify-between items-center space-x-2">
                        <mdicon name="arrow-left" size="15" />
                        <span class="sm:text-base font-semibold">{{ $t('games.header') }}</span>
                    </small>
                </router-link>
            </div>
            <transition name="fade" mode="out-in">
                <slot />
            </transition>

        </div>
    </div>

</template>
<script setup>
import NavbarVue from '@/components/shared/Navbar.vue'
</script>