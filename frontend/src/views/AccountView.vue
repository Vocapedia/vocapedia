<template>
  <div class="max-w-160 mx-auto space-y-5">
    <div>
      <div class="flex items-center justify-between">
        <div class="flex space-x-1 items-center">
          <h2 class="text-xl font-extrabold ">{{ user.name }}</h2>
          <!-- <mdicon name="check-decagram" size="18" /> -->
        </div>
        <button class="smooth-click border-zinc-200 dark:border-zinc-700 border px-2 py-1 rounded-full">
          {{ $t('account.edit_profile') }}
        </button>
      </div>
      <p class="text-zinc-400 mt-1">@{{ user.username }}</p>
      <p class="mt-1">{{ user.bio }}</p>
    </div>

    <button @click="createPost"
      class="cursor-pointer transition-all duration-200 hover:scale-101 active:scale-99 w-full bg-sky-100 dark:bg-sky-700 py-3 font-semibold">
      {{ $t('account.create_new_post') }}
    </button>

    <transition name="fade" mode="out-in">
      <WordLists v-if="response.list" :response="response" />
      <div v-else class="loading-spinner mx-auto" />
    </transition>

  </div>

</template>
<script setup>
import WordLists from "@/components/WordLists.vue";
import { useFetch } from "@/composable/useFetch";
import { ref, onMounted } from "vue";
const response = ref("{}")

const user = ref({
  name: "Vocapedia",
  username: "Vocapedia",
  bio: "Explore the World of Words!",
  lists: [
    { id: 1, title: "Kelime Listem", items: 20 },
    { id: 2, title: "Favoriler", items: 10 },
  ],
});
onMounted(async () => {
  response.value = await useFetch("/public/chapters/search?q=")
})


const createPost = () => {
  console.log("Gönderi oluşturma ekranı açıldı");
};
</script>

<style scoped></style>