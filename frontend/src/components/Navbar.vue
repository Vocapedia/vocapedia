<template>
    <div v-auto-animate class="animate__animated animate__fadeInDown">


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
            <div class="relative flex w-full">
                <input @input="searchShort" @keyup.enter="searchSomething" ref="searchRef" v-model="search" type="text"
                    :placeholder="$t('search_a_list')" class="w-full px-4 border shadow-sm outline-none transition-all 
             bg-white text-zinc-900 border-none rounded-full
             dark:bg-zinc-800 dark:text-white" />
                <button @click="searchSomething"
                    class="w-12 cursor-pointer -ml-12 bg-zinc-100 dark:bg-zinc-700 rounded-r-full p-3">
                    <mdicon name="magnify" size="24" class="dark:text-zinc-400 text-zinc-600" />
                </button>

                <button @click="resetSearch" class="smooth-click2 bg-white dark:bg-zinc-800 -ml-24  p-3">
                    <mdicon name="close" size="24" class="dark:text-red-400 text-red-600" />
                </button>
                <button @click="closeSearch" class="ml-15 smooth-click">
                    <mdicon name="arrow-up" size="24" class="dark:text-blue-400 text-blue-600" />
                </button>
            </div>

        </header>
        <div v-if="searchable && searchList.length > 0"
            class="max-w-160 mx-auto top-full left-0 w-full border border-zinc-200 dark:border-zinc-700 h-64 overflow-y-auto scrollbar-hide rounded-md mt-1">
            <ul>
                <li v-for="item in searchList" :key="item.id"
                    class="p-3 hover:bg-zinc-100 dark:hover:bg-zinc-800 cursor-pointer">
                    <RouterLink @click="searchable = false" :to="'/l/' + BigInt(item.id)" class="flex items-center">
                        <div class="flex-grow">
                            <div class="font-semibold">{{ item.title }}</div>
                            <div class="text-sm text-zinc-500 dark:text-zinc-400">{{ item.description }}</div>
                        </div>
                    </RouterLink>
                </li>
            </ul>
        </div>

        <div v-else-if="search.length >= 3 && searchList.length === 0"
            class="p-2 max-w-160 mx-auto top-full left-0 w-full rounded-md mt-1">
            <div class="p-5 w-full border border-zinc-200 dark:border-zinc-700">
                {{ $t('no_results_found') }}
            </div>
        </div>
    </div>

</template>

<script setup>
import { getUser } from "@/utils/token";
import { ref, nextTick, watch } from "vue";
import { useRouter } from "vue-router";
import { GetLang } from "../i18n/i18n";
import { useFetch } from "@/composable/useFetch";
const router = useRouter()
const search = ref('')
const searchList = ref([])
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
    searchList.value = []
    nextTick(() => {
        searchRef.value?.focus();
    });
}
async function searchShort() {
    if (search.value.length < 3) {
        searchList.value = []
        return
    }
    await useFetch("/public/chapters/search-short?q=" + search.value).then(r => {
        searchList.value = r.list
    })
}
function openSearchable() {
    searchable.value = true
}
function searchSomething() {
    if ((search.value).toString().trim().length == 0) return

    const trimmedSearch = search.value.trim();

    router.push({ name: 'search', query: { q: encodeURIComponent(trimmedSearch) } })
    searchable.value = false
}
</script>