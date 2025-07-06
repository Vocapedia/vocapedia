import { createI18n } from "vue-i18n";
import { ref } from "vue"

const modules = import.meta.glob("./**/*.json", { eager: true });

const messages = {};

for (const path in modules) {
  const pathParts = path
    .replace("./", "")   // en/home.json
    .replace(".json", "")        // en/home
    .split("/");                 // ["en", "home"]

  const [lang, namespace] = pathParts;
  const content = modules[path].default;

  if (!messages[lang]) {
    messages[lang] = {};
  }

  messages[lang][namespace] = content;
}


const lang = ref(localStorage.getItem('lang') || navigator.language || navigator.languages[0] || 'en');
const i18n = createI18n({
  locale: lang.value,
  fallbackLocale: "en",
  messages: messages,
});
export function ChangeLang(n) {
  lang.value = n
  localStorage.setItem("lang", n)
  i18n.global.locale = n;
}
export function GetLang() {
  return lang.value
}
export { i18n };
