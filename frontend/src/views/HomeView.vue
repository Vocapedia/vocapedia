<template>
  <div class="space-y-16">
    <div>
      <h1 class="p-4 text-3xl font-bold">{{ $t('trends') }}</h1>
      <div class="flex justify-center">
        <div class="flex overflow-x-auto space-x-4 p-4 scrollbar-hide">
          <div v-for="i in 10" :key="i" @click="$router.push('/l/' + i)"
            class="min-w-[250px] md:min-w-[300px] smooth-click card h-50 flex justify-center items-center">
            <p class="text-xl p-3">trend / {{ i }}</p>
          </div>
        </div>
      </div>
    </div>

    <div>
      <h1 class="p-4 text-3xl font-bold">{{ $t('personal_info') }}</h1>
      <div class="container mx-auto px-4">
        <div class="grid grid-cols-12 gap-4">

          <div @click="triggerComposePopup = true"
            class="card col-span-12 sm:col-span-12 lg:col-span-4 rounded-xl smooth-click">
            <img class="max-h-48 md:hidden lg:block mx-auto" src="@/assets/logo_full.png" alt="compose">
            <div class="bg-zinc-100 dark:bg-zinc-800 flex justify-center items-center text-xl font-bold p-4 space-x-1">
              <mdicon name="pencil-outline" />
              <span> {{ $t('compose_post') }}</span>
            </div>
          </div>

          <router-link class="card rounded-xl smooth-click col-span-12 sm:col-span-6 lg:col-span-4 " to="/followed">
            <div class="  p-4 max-h-48 overflow-hidden">
              <div class="select-none blur" v-for="i in follows">
                <div>
                  <div class="text-xl font-semibold">{{ i.title }}</div>
                  <div class="text-sm font-thin">{{ i.content }}</div>
                </div>
              </div>
            </div>
            <div class="bg-zinc-100 dark:bg-zinc-800 flex justify-center items-center text-xl font-bold p-4">
              <mdicon name="star-outline" />
              {{ $t('followed') }}
            </div>
          </router-link>

          <router-link class="card rounded-xl smooth-click col-span-12 sm:col-span-6 lg:col-span-4 " to="/notes">
            <div class="p-4 max-h-48 overflow-hidden">
              <div class="select-none blur" v-for="i in notes">
                <div>
                  <div class="text-xl font-semibold">{{ i.title }}</div>
                  <div class="text-sm font-thin">{{ i.content }}</div>
                </div>
              </div>
            </div>
            <div class="bg-zinc-100 dark:bg-zinc-800 flex justify-center items-center space-x-1 text-xl font-bold p-4">
              <mdicon name="text" />
              <span>{{ $t('notes') }}</span>
            </div>
          </router-link>

        </div>
      </div>
    </div>
  </div>
  <div class="">
    <ComposePost v-model="triggerComposePopup">
      <template #description>
        <div class="text-center">
          <h5>!!! Şu an yapım aşamasında !!!</h5>
        </div>
      </template>
      <template #buttons>
        <div class="flex justify-around">
          <button @click="triggerComposePopup = false"
            class="cursor-pointer w-full px-4 py-2 bg-red-500 text-white rounded">
            {{ $t('close') }}
          </button>
        </div>
      </template>
    </ComposePost>
  </div>

</template>

<script setup>
import { ref } from 'vue';
import ComposePost from '@/components/Popup.vue';
const triggerComposePopup = ref(false)


const follows = ref([
  {
    username: 'user1', id: 1,
    title: 'title1',
    content: 'content1',
  },
  {
    username: 'user2', id: 2,
    title: 'title2',
    content: 'content2',
  },
  {
    username: 'user3', id: 3,
    title: 'title3',
    content: 'content3',
  },
  {
    username: 'user4', id: 4,
    title: 'title4',
    content: 'content4',
  }
]);
const notes = ref([{
  username: 'user1', id: 1,
  title: 'title1',
  content: 'content1',
},
{
  username: 'user2', id: 2,
  title: 'title2',
  content: 'content2',
},
{
  username: 'user3', id: 3,
  title: 'title3',
  content: 'content3',
},
{
  username: 'user4', id: 4,
  title: 'title4',
  content: 'content4',
}]);
</script>