import "./assets/main.css";

import { createApp } from "vue";
import { createPinia } from "pinia";
import mdiVue from "mdi-vue/v3";
import * as mdijs from "@mdi/js";
import { autoAnimatePlugin } from "@formkit/auto-animate/vue";
import { MotionPlugin } from '@vueuse/motion'

import App from "./App.vue";
import router from "./router";
import { i18n } from "./i18n/i18n";

import "animate.css";

import "swiper/css";
import "swiper/css/navigation";
import "swiper/css/pagination";
import "swiper/css/scrollbar";
import "swiper/swiper-bundle.css";


import "splitting/dist/splitting.css";
import "splitting/dist/splitting-cells.css";
import Splitting from "splitting";

import "@hoppscotch/vue-toasted/style.css"
import Toasted from "@hoppscotch/vue-toasted"

const app = createApp(App)
  .use(createPinia())
  .use(i18n)
  .use(mdiVue, {
    icons: mdijs,
  })
  .use(MotionPlugin)
  .use(autoAnimatePlugin)
  .use(router)
  .use(Toasted)
  .directive("splitting", {
    mounted(el) {
      const split = Splitting({ target: el });
      const characters = split[0].chars;
      characters.forEach((char, index) => {
        char.style.animation = `fadeInUp 0.5s ease forwards ${index * 0.05}s`;
      });
      el.addEventListener('mouseenter', () => {
        characters.forEach((char, index) => {
          char.style.animation = `none`;
          char.offsetHeight;
          char.style.animation = `fadeInUp 0.5s ease forwards ${index * 0.05}s`;
        });
      });

    },
  });

router.isReady().then(() => {
  app.mount("#app");
});
