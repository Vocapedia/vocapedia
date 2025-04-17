<template>
    <div>
        <div v-if="!isPDFGenerating">

            <transition name="fade" mode="out-in">
                <div v-if="response.chapter">
                    <small v-motion-slide-visible-once-left class="pb-5 flex justify-center">
                        {{
                            $t('list_helper', {
                                lang: getLangByCode(response.chapter.lang).name,
                                target_lang: getLangByCode(response.chapter.target_lang).name
                            })
                        }}
                    </small>
                    <div v-motion-slide-visible-once-top class="space-y-4 text-center">
                        <h1 class="font-bold text-4xl">
                            {{ response.chapter.title }}
                        </h1>
                        <div>{{ response.chapter.description }}</div>
                    </div>

                    <div v-motion-slide-visible-once-bottom
                        class="max-w-160 w-full mx-auto flex justify-around items-center py-5">
                        <button @click="response.chapter.is_favorited ? deleteFavChapter() : favoriteChapter()"
                            class="space-x-2 flex items-center smooth-click bg-yellow-50 dark:bg-yellow-900 rounded-full py-px pl-2 pr-4 font-bold">
                            <mdicon v-if="response.chapter.is_favorited" name="star"
                                class="dark:text-yellow-400 text-yellow-500 " size="32" />
                            <mdicon v-else name="star-outline" class="dark:text-yellow-400 text-yellow-500 "
                                size="32" />
                            <span>{{ response.chapter.fav_count }}</span>
                        </button>
                        <router-link :to="'/l/' + $route.params.id + '/games'" class="smooth-click">
                            <mdicon name="gamepad-variant-outline" size="32" />
                        </router-link>
                    </div>

                    <div v-motion-slide-visible-once-right class="max-w-160 w-full mx-auto">
                        <div class="flex items-center justify-between space-x-5">
                            <div class="flex items-center space-x-2">
                                <router-link :to="$route.path + '?variant=word-list'"
                                    :class="$route.query.variant == 'tutorial' ? 'text-zinc-400 dark:text-zinc-600' : ''"
                                    class="truncate cursor-pointer text-lg font-semibold">
                                    {{ $t('word-list') }}
                                </router-link>
                                <router-link :to="$route.path + '?variant=tutorial'"
                                    :class="$route.query.variant == 'tutorial' ? '' : 'text-zinc-400 dark:text-zinc-600'"
                                    class="truncate cursor-pointer text-lg font-semibold">
                                    {{ $t('tutorial') }}
                                </router-link>
                                <button class="smooth-click bg-sky-100 dark:bg-sky-700 rounded-full p-1"
                                    @click="generatePDF">
                                    <mdicon name="download-outline" />
                                </button>
                                <button class="smooth-click bg-red-100 dark:bg-red-700 rounded-full p-1"
                                    @click="archiveChapter">
                                    <mdicon name="archive-arrow-down-outline" />
                                </button>
                                <router-link
                                    v-if="BigInt(response.chapter.creator.id) == BigInt(getUser().user_id ?? 0)"
                                    :to="'/update/' + route.params.id"
                                    class="smooth-click bg-sky-100 dark:bg-sky-700 rounded-full p-1">
                                    <mdicon name="puzzle-edit-outline" />
                                </router-link>
                            </div>
                            <router-link class="flex truncate items-center space-x-1 font-semibold"
                                :to="'/' + response.chapter.creator.username">
                                <span class="overflow truncate">
                                    @{{ response.chapter.creator.username }}
                                </span>
                            </router-link>
                        </div>
                        <hr class="border-t-2 border-zinc-200 dark:border-zinc-800 my-4 opacity-50">
                    </div>
                    <transition name="fade" mode="out-in">
                        <component :is="currentComponent" :response="response" />
                    </transition>
                </div>
                <div v-else class="loading-spinner mx-auto" />

            </transition>
        </div>
    </div>

</template>

<script setup>
import TutorialView from "./TutorialView.vue"
import WordListView from "./WordListView.vue"
import { ref, watch, shallowRef, onMounted } from "vue"
import { useRoute } from "vue-router"
import { jsPDF } from "jspdf";
import { useFetch } from '@/composable/useFetch';
import { useToast } from "@/composable/useToast"
import { i18n } from "@/i18n/i18n";
import { getLangByCode } from "@/utils/language/languages";
import { getUser } from "@/utils/token";
import { useRouter } from "vue-router";
import "@/assets/Roboto-bold"
import "@/assets/Roboto-normal"
const router = useRouter()

let savedStyles = {
    links: [],
    styles: [],
    inlineStyles: []
};
const saveAndRemoveStyles = () => {
    savedStyles.links = Array.from(document.querySelectorAll('link[rel="stylesheet"]'));
    savedStyles.styles = Array.from(document.querySelectorAll('style'));
    savedStyles.inlineStyles = Array.from(document.querySelectorAll('*')).map(element => ({
        element,
        style: element.getAttribute('style')
    }));
    savedStyles.links.forEach(link => link.remove());
    savedStyles.styles.forEach(style => style.remove());
    savedStyles.inlineStyles.forEach(item => item.element.removeAttribute('style'));
};

const restoreStyles = () => {
    savedStyles.links.forEach(link => document.head.appendChild(link));
    savedStyles.styles.forEach(style => document.head.appendChild(style));
    savedStyles.inlineStyles.forEach(item => {
        if (item.style) {
            item.element.setAttribute('style', item.style);
        }
    });
};
async function archiveChapter() {
    await useFetch("/chapters/archive/" + route.params.id, {
        method: "DELETE",
    }).then(r => {
        router.replace("/")
        toast.show(r.message)
    }).catch(e => {
        toast.show(e.error)
    })
}
const isPDFGenerating = ref(false)
const generatePDF = async () => {
    isPDFGenerating.value = true
    const doc = new jsPDF();
    const marginX = 14;
    let startY = 20;

    doc.setFont('Roboto', "bold");
    doc.setFontSize(20);
    doc.text(response.value.chapter.title, marginX, startY);
    startY += 8;

    doc.setFont('Roboto', "normal");
    doc.setFontSize(12);
    doc.text(response.value.chapter.description, marginX, startY);
    startY += 8;

    doc.setFontSize(12);
    doc.setTextColor(100);
    doc.text(`@${response.value.chapter.creator.username}`, marginX, startY);
    startY += 10;

    saveAndRemoveStyles();
    const tutorialDiv = document.createElement('div');
    tutorialDiv.style.width = '160mm';

    const innerDiv = document.createElement('div');
    innerDiv.style.width = '100%';
    innerDiv.style.fontFamily = "'Roboto', sans-serif";

    innerDiv.innerHTML = response.value.chapter.tutorial;

    const images = innerDiv.querySelectorAll('img');
    innerDiv.querySelectorAll('iframe').forEach(iframe => {
        iframe.remove()
    })
    images.forEach(img => {
        img.style.maxWidth = '50%';
        img.style.width = '100%';
        img.style.height = 'auto';
        img.style.display = 'block';
        img.style.marginLeft = 'auto';
        img.style.marginRight = 'auto';
    });
    let wb = response.value.chapter.word_bases
    tutorialDiv.appendChild(innerDiv);
    const tableContainer = document.createElement('div');
    tableContainer.style.fontFamily = "'Roboto', sans-serif";
    tableContainer.style.gap = '16px';
    tableContainer.style.width = '100%';
    tableContainer.innerHTML = `
<table style="width: 100%; border-collapse: separate;border-spacing: 0 8px; font-size: 13px; font-family: 'Roboto', sans-serif;">
  <thead>
    <tr style="background-color: #2980b9; color: #fff; font-size: 15px;">
      <th style="border: 1px solid #2980b9; padding: 15px 12px; border-top-left-radius: 6px;">${i18n.global.t("chapter.type")}</th>
      <th style="border: 1px solid #2980b9; padding: 15px 12px;">${i18n.global.t("chapter.language")}</th>
      <th style="border: 1px solid #2980b9; padding: 15px 12px; border-top-right-radius: 6px;">${i18n.global.t("chapter.word")}</th>
    </tr>
  </thead>
  <tbody>
    ${wb.map((base) =>
        base.words.map((word, i) => `
      ${i % 2 == 0 ? '<tr style="background-color: #ecf0f1">' : '<tr>'}
      
        ${i === 0 ? `
          <td rowspan="${base.words.length}" style="border-top: 0.5px solid #ccc;border-bottom: 0.5px solid #ccc;border-left: 0.5px solid #ccc; padding: 6px 8px; text-align:center; font-weight: 500; background-color: #ecf0f1;">
            ${i18n.global.t("word_types." + base.type)}
          </td>` : ``}
          <td style="padding: 10px 8px; text-align:center;">
            ${getLangByCode(word.lang).name}
          </td>
          <td style="padding: 10px 8px;">
            ${word.word}
          </td>
        </tr>
      `).join('')
    ).join('')}
  </tbody>
</table>
`;

    tutorialDiv.appendChild(tableContainer);
    document.body.appendChild(tutorialDiv);
    await doc.html(tutorialDiv,
        {
            useCORS: true,
            foreignObjectRendering: true,
            x: marginX,
            y: startY,
            html2canvas: { scale: 0.3 },
            callback: () => {
                restoreStyles()
                document.body.removeChild(tutorialDiv);
            }
        },
    )
    doc.save(`${response.value.chapter.title}.pdf`);
    isPDFGenerating.value = false
};

const route = useRoute()

const response = ref({})
const toast = useToast()

onMounted(async () => {
    response.value = await useFetch("/public/chapters/" + route.params.id)
    if (route.query.variant == 'tutorial') {
        currentComponent.value = TutorialView
    } else {
        currentComponent.value = WordListView
    }
})

async function favoriteChapter() {
    await useFetch("/chapters/favorite?chapter_id=" + route.params.id, {
        method: "POST",
    }).then(r => {
        response.value.chapter.is_favorited = 1
        response.value.chapter.fav_count += 1
        toast.show(r.message)
    }).catch(e => {
        toast.show(e.error)
    })
}
async function deleteFavChapter() {
    await useFetch("/chapters/favorite?chapter_id=" + route.params.id, {
        method: "DELETE",
    }).then(r => {
        response.value.chapter.is_favorited = 0
        response.value.chapter.fav_count -= 1
        toast.show(r.message)
    })
        .catch(e => {
            toast.show(e.error)
        })
}

const currentComponent = shallowRef(WordListView)
watch(route, (newV) => {
    if (newV.query.variant == 'tutorial') {
        currentComponent.value = TutorialView
    } else {
        currentComponent.value = WordListView
    }
})

</script>