<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>

  <router-view v-slot="{ Component, route }">
    <transition name="fade" mode="out-in">
      <component :is="$route.meta.layoutComponent">
        <transition name="fade" mode="out-in">
          <div :key="route.name">
            <component :is="Component" />
          </div>
        </transition>
      </component>
    </transition>
  </router-view>
</template>
