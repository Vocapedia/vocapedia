import { createI18n } from "vue-i18n";
import { ref } from "vue"

import en from "./en.json";
import tr from "./tr.json";
const messages = {
  en: en,
  tr: tr,
};

const lang = ref(localStorage.getItem('lang') || navigator.language || navigator.languages[0]);
const i18n = createI18n({
  locale: lang.value,
  fallbackLocale: "en",
  messages: messages,
});
export function ChangeLang(n) {
  lang.value = n
  localStorage.setItem("lang", n)
}
export function GetLang() {
  return lang.value
}
export { i18n };
