import "./assets/main.css";

import { createApp } from "vue";
import { createPinia } from "pinia";
import mdiVue from "mdi-vue/v3";
import * as mdijs from "@mdi/js";
import { autoAnimatePlugin } from "@formkit/auto-animate/vue";

import App from "./App.vue";
import router from "./router";
import { i18n } from "./i18n/i18n";
import { IonicVue } from "@ionic/vue";

import "animate.css";

import "swiper/css";
import "swiper/css/navigation";
import "swiper/css/pagination";
import "swiper/css/scrollbar";
import "swiper/swiper-bundle.css";


import "splitting/dist/splitting.css";
import "splitting/dist/splitting-cells.css";
import Splitting from "splitting";

const app = createApp(App)
  .use(createPinia())
  .use(i18n)
  .use(mdiVue, {
    icons: mdijs,
  })
  .use(IonicVue)
  .use(autoAnimatePlugin)
  .use(router)
  .directive("splitting", {
    mounted(el) {
      const split = Splitting({ target: el });
      const characters = split[0].chars;
      characters.forEach((char, index) => {
        char.style.animation = `fadeInLeft 0.5s ease forwards ${index * 0.1}s`;
      });
      el.addEventListener('mouseenter', () => {
        characters.forEach((char, index) => {
          char.style.animation = `none`;
          char.offsetHeight;
          char.style.animation = `fadeInLeft 0.5s ease forwards ${index * 0.1}s`;
        });
      });

    },
  });

router.isReady().then(() => {
  app.mount("#app");
});
