<template>
    <div>
        <div class="flex justify-center">
            <input v-model="search" type="text" :placeholder="$t('search_from_list')" class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all 
             bg-white text-zinc-900  border-none
             max-w-160 
             dark:bg-zinc-800 dark:text-white " />
        </div>
        <div v-auto-animate>
            <div :id="item.word" v-for="item in filteredList" :key="item.id" class="flex justify-center py-2.5">
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
import { computed, ref, onMounted } from 'vue';
import { useRoute } from "vue-router"
const search = ref('');
const route = useRoute()
const filteredList = computed(() => {
    return props.response.list.filter(x => x.word.toLowerCase().includes(search.value.toLowerCase()));
});

const props = defineProps({
    response: {
        type: Object,
        required: true,
    }
})

onMounted(() => {
    if (window.location.hash && route.query.variant != 'tutorial') {
        const element = document.querySelector(window.location.hash);
        if (element) {
            element.scrollIntoView({ behavior: 'smooth' });
        }
    }
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