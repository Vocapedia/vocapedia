<template>
  <div>
    <div v-if="isUserLoading">
      <div class="loading-spinner mx-auto" />
    </div>
    <div v-else class="max-w-160 mx-auto space-y-5">
      <div>
        <div v-motion-slide-visible-once-top class="flex items-center justify-between">
          <div class="flex space-x-1 items-center">
            <h2 v-if="user.name" class="text-xl font-extrabold ">{{ user.name }}</h2>
            <mdicon v-if="user.approved" name="check-decagram" class="text-xl text-sky-500" />
          </div>
          <button @click="triggerSettingsPopup = true" v-if="isUsersAccount"
            class="smooth-click border-zinc-200 dark:border-zinc-700 border px-2 py-1 rounded-full">
            {{ $t('account.edit_profile') }}
          </button>
        </div>
        <p v-motion-slide-visible-once-right v-if="user.username" class="text-zinc-400 mt-1">@{{ user.username }}</p>
        <p v-motion-slide-visible-once-left v-if="user.biography" class="mt-1"
          v-html="transformBiography(user.biography)" />
      </div>
      <div v-motion-slide-visible-once-bottom class="text-center flex" v-if="isUsersAccount">
        <router-link to="/compose"
          class="cursor-pointer w-full smooth-click2 w-full bg-sky-100 dark:bg-sky-700 py-3 font-semibold">
          {{ $t('account.create_new_post') }}
        </router-link>
      </div>
      <hr class="border-t-2 border-zinc-200 dark:border-zinc-800 my-4 opacity-50">

      <transition name="fade" mode="out-in">
        <div v-if="response.list">
          <WordLists :response="response" />
          <!-- <WordLists v-infinite-scroll="onLoadMore" :response="response" />-->
        </div>

        <div v-else class="loading-spinner mx-auto" />
      </transition>
      <SettingsPopup v-model="triggerSettingsPopup">
        <template #header>
          <h2 class="text-xl font-semibold">
            {{ $t('account.edit_profile') }}
          </h2>
        </template>
        <template #description>
          <div class="space-y-5">
            <hr class="border-t-2 border-zinc-200 dark:border-zinc-700 opacity-50">
            <div class="space-y-3">
              <div>
                <label>name</label>
                <input v-model="EditUser.name" type="text" :placeholder="$t('account.edit.name')" class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all 
             bg-zinc-100 text-zinc-900  border-none
             max-w-160 
             dark:bg-zinc-900 dark:text-white " />

              </div>

              <input v-model="EditUser.username" type="text" :placeholder="$t('account.edit.username')" class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all 
             bg-zinc-100 text-zinc-900  border-none
             max-w-160 
             dark:bg-zinc-900 dark:text-white " />

              <textarea v-model="EditUser.biography" type="text" :placeholder="$t('account.edit.biography')" class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all 
             bg-zinc-100 text-zinc-900  border-none
             max-w-160 
             dark:bg-zinc-900 dark:text-white " />

            </div>

          </div>


        </template>
        <template #buttons>
          <div class="flex justify-around space-x-5">
            <button @click="triggerSettingsPopup = false"
              class="smooth-click2 cursor-pointer w-full px-4 py-2 font-semibold text-xl rounded bg-red-100 text-red-900 dark:bg-red-700/20 dark:text-red-100">
              {{ $t('close') }}
            </button>
            <div v-auto-animate class="w-full">
              <div v-if="isLoading">
                <div class="loading-spinner mx-auto"></div>
              </div>
              <button v-else @click="Edit"
                class="smooth-click2 cursor-pointer w-full px-4 py-2 font-semibold text-xl rounded bg-sky-100 text-sky-900 dark:bg-sky-700/20 dark:text-sky-100">
                {{ $t('save') }}
              </button>
            </div>
          </div>
        </template>
      </SettingsPopup>

    </div>
  </div>
</template>
<script setup>
import SettingsPopup from "@/components/Popup.vue"
import { vInfiniteScroll } from '@vueuse/components'

import WordLists from "@/components/WordLists.vue";
import { useFetch } from "@/composable/useFetch";
import { useToast } from "@/composable/useToast";
import { i18n } from "@/i18n/i18n";
import { getDevice, getUser } from "@/utils/token";
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import { useRoute } from "vue-router";
const response = ref("{}")
const route = useRoute()
const router = useRouter()
const triggerSettingsPopup = ref(false)
const isUsersAccount = ref(false)
const isLoading = ref(false)
const user = ref({
  name: "",
  username: "",
  biography: "",
});
const EditUser = ref({
  name: "",
  username: "",
  biography: "",
})
const transformBiography = (text) => {
  return (text ?? "").replace(/@([a-zA-Z0-9_]+)/g, (match, username) => {
    return `<a class="dark:text-blue-400 text-blue-700" href="/${username}">@${username}</a>`;
  });
};

async function onLoadMore() {
  console.log("sdkjafnk")
  return
  isLoading.value = true;

  // Load more data when scrolled to the bottom
  try {
    const newItems = await fetchData(response.value.list.length);
    response.value.list = [...response.value.list, ...newItems]; // Append new items to the list
  } catch (error) {
    console.error('Error loading data:', error);
  } finally {
    isLoading.value = false;
  }
}


const fetchData = async (startIndex) => {
  // Fetch additional items from API (update this with your real API endpoint)
  const response = await useFetch(`/public/chapters/user?username=${route.params.username}&page=${startIndex}&limit=10`);
  return response.list; // Return the new items for the list
};

const toast = useToast()
const isUserLoading = ref(false)
onMounted(async () => {
  isUserLoading.value = true
  isUsersAccount.value = route.params.username.toLowerCase() == (getUser().username ?? "").toLowerCase()
  response.value = await useFetch("/public/chapters/user?username=" + route.params.username)
  user.value = await useFetch("/public/user?username=" + route.params.username).catch(e => {
    router.replace("/search?q=" + route.params.username)
  })
  EditUser.value.biography = user.value.biography
  EditUser.value.username = user.value.username
  EditUser.value.name = user.value.name
  isUserLoading.value = false
})

async function Edit() {
  isLoading.value = true
  await useFetch("/user", {
    method: "PUT",
    body: {
      name: EditUser.value.name,
      username: EditUser.value.username,
      biography: EditUser.value.biography,
      device: getDevice(),
    }
  }).then(r => {
    toast.show(i18n.global.t('account.edit.success'))
    localStorage.setItem("token", r.token)
    router.replace("/" + EditUser.value.username).then(() => router.go())
  }).catch(e => {
    toast.show(e.error)
    isLoading.value = false
  })
  isLoading.value = false

}
</script>