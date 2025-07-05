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
            <mdicon v-if="user.is_teacher" name="school" class="text-xl text-red-500"/>
          </div>
          <div class="flex space-x-2">

            <button @click="triggerSettingsPopup = true" v-if="isUsersAccount"
              class="smooth-click border-zinc-200 dark:border-zinc-700 border px-2 py-1 rounded-full">
              {{ $t('account.edit_profile') }}
            </button>
            <button @click="triggerDiscardedChaptersPopup = true" v-if="isUsersAccount"
              class="smooth-click border-zinc-200 dark:border-zinc-700 border px-2 py-1 rounded-full">
              <mdicon name="archive-arrow-down-outline" />
            </button>
            <button @click="triggerLanguagePreferencesPopup = true" v-if="isUsersAccount"
              class="smooth-click border-zinc-200 dark:border-zinc-700 border px-2 py-1 rounded-full">
              <mdicon name="translate" />
            </button>
            <button @click="triggerTeacherRequestPopup = true" v-if="isUsersAccount && !user.is_teacher"
              class="smooth-click border-zinc-200 dark:border-zinc-700 border px-2 py-1 rounded-full">
              <mdicon name="school" />
            </button>
          </div>

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

      <div v-motion-slide-bottom>
        <WordLists :key="$route.params.username" :uri="'/public/chapters/user?username=' + $route.params.username" />
      </div>
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

      <!-- Language Preferences Popup -->
      <SettingsPopup v-model="triggerLanguagePreferencesPopup">
        <template #header>
          <h2 class="text-xl font-semibold">
            {{ $t('language_preferences.title') }}
          </h2>
        </template>
        <template #description>
          <div class="space-y-5">
            <hr class="border-t-2 border-zinc-200 dark:border-zinc-700 opacity-50">
            <LanguagePreferencesForm ref="languagePreferencesForm" @update:loading="setLanguagePreferencesLoading" />
          </div>
        </template>
        <template #buttons>
          <div class="flex justify-around space-x-5">
            <button @click="triggerLanguagePreferencesPopup = false"
              class="smooth-click2 cursor-pointer w-full px-4 py-2 font-semibold text-xl rounded bg-red-100 text-red-900 dark:bg-red-700/20 dark:text-red-100">
              {{ $t('common.cancel') }}
            </button>
            <div v-auto-animate class="w-full">
              <div v-if="isLanguagePreferencesLoading">
                <div class="loading-spinner mx-auto"></div>
              </div>
              <button v-else @click="saveLanguagePreferences"
                class="smooth-click2 cursor-pointer w-full px-4 py-2 font-semibold text-xl rounded bg-sky-100 text-sky-900 dark:bg-sky-700/20 dark:text-sky-100">
                {{ $t('common.save') }}
              </button>
            </div>
          </div>
        </template>
      </SettingsPopup>

      <!-- Teacher Request Popup -->
      <SettingsPopup v-model="triggerTeacherRequestPopup">
        <template #header>
          <h2 class="text-xl font-semibold">
            {{ $t('teacher_request.title') }}
          </h2>
        </template>
        <template #description>
          <div class="space-y-5">
            <hr class="border-t-2 border-zinc-200 dark:border-zinc-700 opacity-50">
            <TeacherRequestForm ref="teacherRequestForm" @update:loading="setTeacherRequestLoading" />
          </div>
        </template>
        <template #buttons>
          <div class="flex justify-around space-x-5">
            <button @click="triggerTeacherRequestPopup = false"
              class="smooth-click2 cursor-pointer w-full px-4 py-2 font-semibold text-xl rounded bg-red-100 text-red-900 dark:bg-red-700/20 dark:text-red-100">
              {{ $t('common.cancel') }}
            </button>
            <div v-auto-animate class="w-full">
              <div v-if="isTeacherRequestLoading">
                <div class="loading-spinner mx-auto"></div>
              </div>
              <button v-else @click="submitTeacherRequest"
                class="smooth-click2 cursor-pointer w-full px-4 py-2 font-semibold text-xl rounded bg-green-100 text-green-900 dark:bg-green-700/20 dark:text-green-100">
                {{ $t('teacher_request.request') }}
              </button>
            </div>
          </div>
        </template>
      </SettingsPopup>

      <!-- Discarded Chapters Popup -->
      <SettingsPopup v-model="triggerDiscardedChaptersPopup">
        <template #header>
          <h2 class="text-xl font-semibold">
            {{ $t('discarded_chapters.title') }}
          </h2>
        </template>
        <template #description>
          <div class="space-y-5">
            <hr class="border-t-2 border-zinc-200 dark:border-zinc-700 opacity-50">
            <DiscardedChaptersForm />
          </div>
        </template>
        <template #buttons>
          <div class="flex justify-center">
            <button @click="triggerDiscardedChaptersPopup = false"
              class="smooth-click2 cursor-pointer px-6 py-2 font-semibold text-xl rounded bg-zinc-100 text-zinc-900 dark:bg-zinc-700/20 dark:text-zinc-100">
              {{ $t('common.close') }}
            </button>
          </div>
        </template>
      </SettingsPopup>

    </div>
  </div>
</template>
<script setup>
import SettingsPopup from "@/components/Popup.vue"
import LanguagePreferencesForm from "@/components/LanguagePreferencesForm.vue"
import TeacherRequestForm from "@/components/TeacherRequestForm.vue"
import DiscardedChaptersForm from "@/components/DiscardedChaptersForm.vue"

import WordLists from "@/components/WordLists.vue";
import { useFetch } from "@/composable/useFetch";
import { useToast } from "@/composable/useToast";
import { i18n } from "@/i18n/i18n";
import { getDevice, getUser } from "@/utils/token";
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import { useRoute } from "vue-router";
const route = useRoute()
const router = useRouter()
const triggerSettingsPopup = ref(false)
const triggerLanguagePreferencesPopup = ref(false)
const triggerTeacherRequestPopup = ref(false)
const triggerDiscardedChaptersPopup = ref(false)
const isUsersAccount = ref(false)
const isLoading = ref(false)
const isLanguagePreferencesLoading = ref(false)
const isTeacherRequestLoading = ref(false)

// Component refs
const languagePreferencesForm = ref(null)
const teacherRequestForm = ref(null)
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


const toast = useToast()
const isUserLoading = ref(false)
onMounted(async () => {
  isUserLoading.value = true
  isUsersAccount.value = route.params.username.toLowerCase() == (getUser().username ?? "").toLowerCase()
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

// Language Preferences Functions
const setLanguagePreferencesLoading = (loading) => {
  isLanguagePreferencesLoading.value = loading
}

const saveLanguagePreferences = async () => {
  const success = await languagePreferencesForm.value?.saveLanguagePreferences()
  if (success) {
    triggerLanguagePreferencesPopup.value = false
  }
}

// Teacher Request Functions
const setTeacherRequestLoading = (loading) => {
  isTeacherRequestLoading.value = loading
}

const submitTeacherRequest = async () => {
  const success = await teacherRequestForm.value?.requestTeacherStatus()
  if (success) {
    triggerTeacherRequestPopup.value = false
  }
}
</script>