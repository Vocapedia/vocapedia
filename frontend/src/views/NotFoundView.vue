<template>
    <div class="flex flex-col items-center justify-center">
        <h1 class="text-4xl font-bold text-red-500 mb-4">{{ $t('page_not_found.title') }}</h1>
        <p class="text-lg text-gray-600 mb-6">{{ $t('page_not_found.message') }}</p>
        <router-link to="/"
            class="smooth-click px-6 py-3 bg-blue-500 text-white rounded-lg text-xl font-semibold hover:bg-blue-600 transition-colors">
            {{ $t('page_not_found.button_text') }}
        </router-link>
    </div>
</template>