<template>
    <div>
        <small class="pb-5 flex justify-center">
            {{
                $t('list_helper', {
                    lang: $t(response.lang),
                    s: response.target_langs.length > 1 ? 'ler' : '',
                    target_lang: response.target_langs.map(x => ' ' + $t(x)).toString()
                })
            }}
        </small>
        <div class="space-y-4 text-center">
            <h1 class="font-bold text-4xl">
                {{ response.title }}
            </h1>
            <div>{{ response.description }}</div>
        </div>

        <div class="max-w-160 w-full mx-auto flex justify-around items-center py-5">
            <button class="flex items-center smooth-click bg-yellow-50 dark:bg-yellow-900 rounded-full px-2 font-bold">
                <mdicon name="star-outline" class="dark:text-yellow-400 text-yellow-500" size="32" />
                <span>4198</span>
            </button>
            <router-link :to="'/l/' + $route.params.id + '/games'" class="smooth-click">
                <mdicon name="gamepad-variant-outline" size="32" />
            </router-link>
        </div>

        <div class="flex py-5 justify-center">
            <input v-model="search" type="text" :placeholder="$t('search_from_list')" class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all 
             bg-white text-zinc-900  border-none
             max-w-160 
             dark:bg-zinc-800 dark:text-white " />
        </div>
        <div v-auto-animate class="space-y-5">
            <div v-for="item in filteredList" :key="item.id" class="flex justify-center">
                <div class="max-w-160 w-full ">
                    <div v-auto-animate class="card transition duration-200  hover:shadow p-4">
                        <div class="flex justify-between items-center space-x-5">
                            <div class="font-bold text-xl capitalize ">{{ item.word }}</div>
                            <span
                                class="languages bg-blue-200 dark:bg-blue-800 px-2 py-1 rounded-full text-blue-800 dark:text-blue-200">
                                {{ item.type }}
                            </span>
                        </div>
                        <div class="font-light pt-5">{{ item.description }}</div>

                        <div class="languages">
                            <div v-if="item.example.length > 0">
                                <div v-for="example in item.example" class="p-5 flex items-end font-light text-sm">
                                    <mdicon name="arrow-right" />
                                    <span>{{ example }}</span>
                                </div>
                            </div>
                            <div class="text-sm text-end text-sky-900 dark:text-sky-200 "> {{ $t(item.lang) }}</div>
                            <hr class="border-t-2 border-zinc-200 dark:border-zinc-800 my-4 opacity-50">
                        </div>
                        <div class="languages" v-for="(sub, i) in item.languages" :key="i">
                            <div class="font-bold text-xl capitalize py-5">{{ sub.word }}</div>
                            <div class="font-light">{{ sub.description }}</div>
                            <div v-if="sub.example.length > 0">
                                <div v-for="example in sub.example" class="p-5 flex items-end font-light text-sm">
                                    <mdicon name="arrow-right" />
                                    <span>{{ example }}</span>
                                </div>
                            </div>
                            <div class="text-sm text-end text-sky-900 dark:text-sky-200 "> {{ $t(sub.lang) }}</div>
                            <hr v-if="i < item.languages.length - 1"
                                class="border-t-2 border-zinc-200 dark:border-zinc-800 my-4 opacity-50">
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { computed, ref } from 'vue';
import fake_response from "@/fake/list.json";

const search = ref('');

const filteredList = computed(() => {
    return response.value.list.filter(x => x.word.toLowerCase().includes(search.value.toLowerCase()));
});

const response = ref(fake_response);
</script>
<style scoped>
.languages {
    opacity: 0;
    visibility: hidden;
    height: 0;
    overflow: hidden;
    transition: opacity 0.3s ease, visibility 0s linear 0.3s, height 0.3s ease;
}

.card:hover .languages {
    opacity: 1;
    visibility: visible;
    height: auto;
    transition: opacity 0.3s ease, visibility 0s linear 0s, height 0.3s ease;
}
</style>