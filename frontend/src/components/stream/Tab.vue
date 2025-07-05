<template>
    <div  class="space-y-4">
        <div v-if="streams.length === 0" class="text-center py-12">
            <div
                class="w-16 h-16 bg-gray-100 dark:bg-gray-700 rounded-full flex items-center justify-center mx-auto mb-4">
                <mdicon name="video-off" size="24" class="text-gray-400" />
            </div>
            <p class="text-gray-500 dark:text-gray-400">{{ $t('stream.noRecentStreams') }}</p>
        </div>

        <div v-else class="grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">

            <PreviewCard v-for="stream in streams" :form="{
                title: stream.title,
                description: stream.description,
                lang: stream.lang,
                target_lang: stream.target_lang,
                scheduled_at: stream.scheduled_at,
                duration: stream.duration,
                max_participants: stream.max_participants,
            }" :username="stream.creator.username" />
        </div>
    </div>
</template>
<script setup>
import PreviewCard from './PreviewCard.vue';
defineProps({
    streams: {
        type: Array,
        default: () => []
    }
});
</script>