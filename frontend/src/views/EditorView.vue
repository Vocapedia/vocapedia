<template>
    <div>

        <button @click="copyHTMLText" class="smooth-click2 w-full p-3 bg-sky-100 dark:bg-sky-700">{{$t('editor.copy')}}</button>

        <TextEditor v-model="editorContent" />
    </div>
</template>

<script setup>
import { ref } from "vue"
import TextEditor from "@/components/TextEditor.vue"
import { i18n } from '@/i18n/i18n';
import { useToast } from "@/composable/useToast"
const toast = useToast()

const editorContent = ref("<p>"+i18n.global.t('editor.placeholder')+"</p>")
function copyHTMLText (){
    navigator.clipboard.writeText(editorContent.value);
    toast.show(i18n.global.t('editor.copied'))
}
</script>
