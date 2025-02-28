<template>
    <div v-for="l in props.response.list">
        <router-link :to="'/l/' + l.id">
            <div class="card hover:scale-101 transition-all duration-200">
                <h1 class="text-xl font-semibold p-5">
                    {{ l.title }}
                </h1>
                <h4 class="px-5">
                    {{ l.description }}
                </h4>
                <div class="flex justify-evenly pt-5">
                    <button class="bg-yellow-50 dark:bg-yellow-900 rounded-full px-2 font-bold">
                        <span class="flex items-center space-x-1 dark:text-yellow-400 text-yellow-500">
                            <mdicon name="star-outline" class="dark:text-yellow-400 text-yellow-500" size="27" />
                            <span>{{ l.fav_count }}</span>
                        </span>
                    </button>
                    <button class="bg-blue-50 dark:bg-blue-900 rounded-full px-2 font-bold">
                        <span class="flex items-center space-x-1 dark:text-blue-400 text-blue-500">
                            <mdicon name="format-letter-case" class="dark:text-blue-400 text-blue-500" size="27" />
                            <span>{{ l.word_count }}</span>
                        </span>
                    </button>
                </div>
                <small class="flex justify-end px-2 py-1">
                    <span>
                        {{ l.creator.username }}
                    </span>
                </small>
            </div>
        </router-link>
    </div>
</template>


<script setup>
const props = defineProps({
    response: {
        type: Object,
        required: true,
    }
})
</script>