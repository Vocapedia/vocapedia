<template>
    <transition name="fade" mode="out-in">
        <div v-if="response.chapter">
            <small class="pb-5 flex justify-center">
                {{
                    $t('list_helper', {
                        lang: getLangByCode(response.chapter.lang).name,
                        target_lang: getLangByCode(response.chapter.target_lang).name
                    })
                }}
            </small>
            <div class="space-y-4 text-center">
                <h1 class="font-bold text-4xl">
                    {{ response.chapter.title }}
                </h1>
                <div>{{ response.chapter.description }}</div>
            </div>

            <div class="max-w-160 w-full mx-auto flex justify-around items-center py-5">
                <button @click="response.chapter.is_favorited ? deleteFavChapter() : favoriteChapter()"
                    class="space-x-2 flex items-center smooth-click bg-yellow-50 dark:bg-yellow-900 rounded-full py-px pl-2 pr-4 font-bold">
                    <mdicon v-if="response.chapter.is_favorited" name="star"
                        class="dark:text-yellow-400 text-yellow-500 " size="32" />
                    <mdicon v-else name="star-outline" class="dark:text-yellow-400 text-yellow-500 " size="32" />
                    <span>{{ response.chapter.fav_count }}</span>
                </button>
                <router-link :to="'/l/' + $route.params.id + '/games'" class="smooth-click">
                    <mdicon name="gamepad-variant-outline" size="32" />
                </router-link>
            </div>

            <div class="max-w-160 w-full mx-auto">
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
                        <button class="smooth-click bg-sky-100 dark:bg-sky-700 rounded-full p-1" @click="generatePDF">
                            <mdicon name="download-outline" />
                        </button>
                        <router-link :to="'/update/' + route.params.id"
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

</template>

<script setup>
import TutorialView from "./TutorialView.vue"
import WordListView from "./WordListView.vue"
import { ref, watch, shallowRef, onMounted } from "vue"
import { useRoute } from "vue-router"
import { jsPDF } from "jspdf";
import autoTable from "jspdf-autotable";
import "@/assets/Roboto-bold"
import "@/assets/Roboto-normal"

const generatePDF = async () => {
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
    response.value.chapter.word_bases.forEach((base) => {
        doc.setFontSize(12);
        const tableData = base.words.map((word) =>
            [
                i18n.global.t("word_types." + base.type),
                word.lang.toUpperCase(),
                word.word,
                word.description,
                (word.examples ?? []).join("\n")
            ]
        );
        if (doc.lastAutoTable && (doc.lastAutoTable.finalY + tableData.length * 25) > doc.internal.pageSize.height) {
            doc.addPage();
            startY = 20;
        }
        autoTable(doc, {
            startY,
            head: doc.lastAutoTable ? null : [[i18n.global.t("chapter.type"), i18n.global.t("chapter.language"), i18n.global.t("chapter.word"), i18n.global.t("chapter.description"), i18n.global.t("chapter.example")]],
            body: tableData,
            theme: "grid",
            styles: { font: "Roboto", fontSize: 10, cellPadding: 4, overflow: 'ellipsize', halign: "center" },
            headStyles: { font: "Roboto", fillColor: [41, 128, 185], textColor: 255, fontStyle: "bold" },
            columnStyles: {
                0: { halign: "center", cellWidth: 30 },
                1: { halign: "center", cellWidth: 15 },
                2: { halign: "left", cellWidth: 35 },
                3: { halign: "left", cellWidth: 50 },
                4: { halign: "left", cellWidth: 51.77 }
            }
        });

        startY = doc.lastAutoTable.finalY + 7;
    });


    doc.save(`${response.value.chapter.title}.pdf`);
};

const route = useRoute()
import { useFetch } from '@/composable/useFetch';
import { useToast } from "@/composable/useToast"
import { i18n } from "@/i18n/i18n";
import { getLangByCode } from "@/utils/language/languages";

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
    })
        .catch(e => {
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