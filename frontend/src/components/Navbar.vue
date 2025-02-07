<template>
    <div class="max-w-6xl mx-auto">
        <div class="flex justify-between items-center">
            <nav>
                <RouterLink to="/">
                    <img width="100" src="@/assets/logo.png" alt="">
                </RouterLink>
            </nav>
            <nav class="space-x-4 p-4">
                <RouterLink to="/">{{ $t('home') }}</RouterLink>
                <RouterLink to="/account">{{ $t('account') }}</RouterLink>
            </nav>
        </div>
    </div>

</template>