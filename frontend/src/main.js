import "./assets/main.css";

import { createApp } from "vue";
import { createPinia } from "pinia";
import mdiVue from "mdi-vue/v3";
import * as mdijs from "@mdi/js";

import App from "./App.vue";
import router from "./router";
import { i18n } from "./i18n/i18n";

import "animate.css";
const app = createApp(App)
  .use(createPinia())
  .use(i18n)
  .use(mdiVue, {
    icons: mdijs,
  })
  .use(router);

app.mount("#app");
