<template>
    <div class="animate__animated animate__fadeInDown">

        <transition mode="out-in" name="slide-fade">
            <header v-if="!searchable" class="container mx-auto flex justify-between items-center p-4">
                <RouterLink to="/" class="flex-shrink-0 py-2">
                    <span class="text-2xl font-logo">Vocapedia</span>
                </RouterLink>

                <nav :class="searchable ? 'max-w-160 w-full mx-auto' : ''" class="flex items-center space-x-4">
                    <button @click="openSearchable" class="smooth-click">
                        <mdicon name="magnify" :size="iconSize" />
                    </button>
                    <RouterLink to="/settings" class="smooth-click">
                        <mdicon name="cog-outline" :size="iconSize" />
                    </RouterLink>
                    <RouterLink to="/account" class="smooth-click">
                        <mdicon name="account-circle-outline" :size="iconSize" />
                    </RouterLink>
                    <SwitchThemeVue />
                </nav>
            </header>
            <header v-else class="container mx-auto  max-w-160 w-full flex justify-between items-center p-4">
                <div class="flex w-full">
                    <input ref="searchRef" v-model="search" type="text" :placeholder="$t('search_a_list')" class="w-full p-3 border shadow-sm outline-none transition-all 
             bg-white text-zinc-900 border-none rounded-full
             dark:bg-zinc-800 dark:text-white" />
                    <button @click="searchSomething" class="w-12 cursor-pointer -ml-12 bg-zinc-500/15 rounded-r-full p-3">
                        <mdicon name="magnify" class="dark:text-zinc-400 text-zinc-600" />
                    </button>

                    <button @click="closeSearch" class="smooth-click -ml-24 rounded-full p-3">
                        <mdicon name="close" class="dark:text-red-400 text-red-600" />
                    </button>

                </div>
            </header>
        </transition>
    </div>

</template>

<script setup>
import SwitchThemeVue from "@/components/SwitchTheme.vue";
import { ref, onMounted, onBeforeUnmount, nextTick, watch } from "vue";

const iconSize = ref(getSize());
const search = ref('')
const searchRef = ref(null)
const searchable = ref(false)
function getSize() {
    const width = window.innerWidth;

    if (width < 640) return 24
    else return 30
}

watch(searchRef, (newVal) => {
    if (newVal) {
        nextTick(() => {
            searchRef.value?.focus();
        });
    }
});

function closeSearch() {
    searchable.value = false
    search.value = ''
}

function openSearchable() {
    searchable.value = true
}
function searchSomething() {
    console.log(search.value)
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