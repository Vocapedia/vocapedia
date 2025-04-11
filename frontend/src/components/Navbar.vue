<template>
    <div class="animate__animated animate__fadeInDown">

        <transition mode="out-in" name="slide-fade">
            <header v-if="!searchable" class="container mx-auto flex justify-between items-center p-4">
                <RouterLink to="/" class="flex flex-shrink-0 py-2 space-x-1">
                    <span class="text-2xl font-logo">Vocapedia</span>
                    <small class="text-[10px]">{{ GetLang().slice(0, 2).toUpperCase() }}</small>
                </RouterLink>

                <nav :class="searchable ? 'max-w-160 w-full mx-auto' : ''" class="flex items-center space-x-4">
                    <button @click="openSearchable" class="smooth-click">
                        <mdicon name="magnify" size="24" />
                    </button>
                    <RouterLink to="/settings" class="smooth-click">
                        <mdicon name="cog-outline" size="24" />
                    </RouterLink>
                    <RouterLink :to="token ? '/' + getUser().username : '/login'" class="smooth-click">
                        <mdicon name="account-circle-outline" size="24" />
                    </RouterLink>
                </nav>
            </header>
            <header v-else class="container mx-auto  max-w-160 w-full flex justify-between items-center p-4">
                <div class="flex w-full">
                    <input @keyup.enter="searchSomething" ref="searchRef" v-model="search" type="text"
                        :placeholder="$t('search_a_list')" class="w-full px-4 border shadow-sm outline-none transition-all 
             bg-white text-zinc-900 border-none rounded-full
             dark:bg-zinc-800 dark:text-white" />
                    <button @click="searchSomething"
                        class="w-12 cursor-pointer -ml-12 bg-zinc-500/15 rounded-r-full p-3">
                        <mdicon name="magnify" size="24" class="dark:text-zinc-400 text-zinc-600" />
                    </button>

                    <button @click="resetSearch" class="smooth-click -ml-24 rounded-full p-3">
                        <mdicon name="close" size="24" class="dark:text-red-400 text-red-600" />
                    </button>
                    <button @click="closeSearch" class="ml-15 smooth-click">
                        <mdicon name="arrow-up" size="24" class="dark:text-blue-400 text-blue-600" />
                    </button>
                </div>
            </header>
        </transition>
    </div>

</template>

<script setup>
import { getUser } from "@/utils/token";
import { ref, nextTick, watch } from "vue";
import { useRouter } from "vue-router";
import { GetLang } from "../i18n/i18n";
const router = useRouter()
const search = ref('')
const searchRef = ref(null)
const searchable = ref(false)

const token = localStorage.getItem("token")
watch(searchRef, (newVal) => {
    if (newVal) {
        nextTick(() => {
            searchRef.value?.focus();
        });
    }
});

function closeSearch() {
    searchable.value = false
}

function resetSearch() {
    search.value = ''
    nextTick(() => {
        searchRef.value?.focus();
    });
}

function openSearchable() {
    searchable.value = true
}
function searchSomething() {
    if ((search.value).toString().trim().length == 0) return
    router.push({ name: 'search', query: { q: search.value } })
    searchable.value = false
}
</script>