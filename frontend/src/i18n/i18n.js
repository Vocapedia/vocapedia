import { createI18n } from "vue-i18n";

import en from "./en.json";
import tr from "./tr.json";

const messages = {
  en: en,
  tr: tr,
};

const lang = navigator.language || navigator.languages[0];
const i18n = createI18n({
  locale: lang,
  fallbackLocale: "en",
  messages: messages,
});

export { i18n };
