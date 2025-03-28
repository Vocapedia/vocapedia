<template>
  <div>
    <router-view v-slot="{ Component, route }">
      <transition name="fade" mode="out-in">
        <component :is="$route.meta.layoutComponent">
          <transition name="fade" mode="out-in">
            <div :key="route.name">
              <component :is="Component" :key="$route.fullPath" />
            </div>
          </transition>
        </component>
      </transition>
    </router-view>
  </div>
</template>
