<template>
    <div class="max-w-4xl mx-auto px-4">
        <span class="text-2xl font-bold">
            {{ $t('games') }}
        </span>
        <hr class="border-t-2 border-zinc-200 dark:border-zinc-800 my-4 opacity-50">
        <div class="grid grid-cols-12 gap-4">
            <router-link :to="'/l/' + $route.params.id + '/game/test'"
                class="card flex justify-center items-center col-span-12 sm:col-span-6 lg:col-span-4 rounded-xl smooth-click h-48">
                <span class="text-xl font-bold p-4"> {{ $t('test') }}</span>
            </router-link>
        </div>
    </div>
</template>