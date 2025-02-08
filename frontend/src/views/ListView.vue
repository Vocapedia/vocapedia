<template>
    <div>
        <!-- <div class="max-w-120 mx-auto grid grid-cols-12 py-10">
            <span class="col-span-4 text-center">
                {{ $t(response.lang) }}
            </span>
            <mdicon class="col-span-4 mx-auto" name="translate" />
            <div class="space-x-2 col-span-4 text-center">
                <span v-for="(lang, i) in response.target_langs">
                    {{ $t(lang) }}
                    <span v-if="i < response.target_langs.length - 1">,</span>
                </span>
            </div>
        </div> -->
        <h1 class="font-bold text-4xl flex justify-center">{{ response.title }}</h1>
        <div class="flex py-5 justify-center">
            <input v-model="search" type="text" :placeholder="$t('search_from_list')" class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all 
             bg-white text-zinc-900  border-none
             max-w-120 
             dark:bg-zinc-800 dark:text-white " />
        </div>
        <div v-auto-animate class="space-y-5">

            <div v-for="item in response.list" :key="item.id" class="flex justify-center">
                <div class="max-w-120 w-full">
                    <div v-auto-animate class="card p-4">

                        <div class="font-bold text-xl capitalize ">{{ item.word }}</div>
                        <div class="font-light">{{ item.description }}</div>

                        <div class="languages">
                            <div class="p-5 flex items-end font-light text-sm">
                                <mdicon name="arrow-right" />
                                <span>{{ item.example }}</span>
                            </div>
                            <div class="text-sm text-end text-sky-900 dark:text-sky-200 "> {{ $t(item.lang) }}</div>
                            <hr class="border-t-2 border-zinc-200 dark:border-zinc-800 my-4 opacity-50">
                        </div>
                        <div class="languages" v-for="(sub, i) in item.languages" :key="i">
                            <div class="font-bold text-xl capitalize py-5">{{ sub.word }}</div>
                            <div class="font-light">{{ sub.description }}</div>
                            <div class="p-5 flex items-end font-light text-sm">
                                <mdicon name="arrow-right" />
                                <span>{{ sub.example }}</span>
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
import { ref } from 'vue';

const search = ref('');
const response = ref({
    id: 1,
    title: 'List',
    lang: 'en',
    target_langs: ['tr'],
    list: [
        {
            id: 1,
            lang: 'en',
            word: 'title',
            description: 'A title is a word or phrase that describes something, such as a book or article.',
            example: 'The title of the book is very inspiring.',
            languages: [
                {
                    id: 'a',
                    lang: 'tr',
                    word: 'başlık',
                    description: 'Başlık, bir şeyi tanımlayan kelime ya da ifadedir, örneğin bir kitap ya da makale başlığı.',
                    example: 'Bu başlık dikkat çekici olmalı.',
                    base_id: 1,
                }
            ]
        }
    ]
});
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