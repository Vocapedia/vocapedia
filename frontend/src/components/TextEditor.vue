<script setup>
import { ref, onBeforeUnmount, onMounted } from "vue";
import { Editor, EditorContent } from "@tiptap/vue-3";
import StarterKit from "@tiptap/starter-kit";
import Image from "@tiptap/extension-image";
import Youtube from "@tiptap/extension-youtube";
import Underline from "@tiptap/extension-underline";
import TextStyle from "@tiptap/extension-text-style";
import Color from "@tiptap/extension-color";
import TextAlign from "@tiptap/extension-text-align";
import FontFamily from "@tiptap/extension-font-family";
import FontSize from "@tiptap/extension-font-size";
import Table from "@tiptap/extension-table";
import TableRow from "@tiptap/extension-table-row";
import TableCell from "@tiptap/extension-table-cell";
import TableHeader from "@tiptap/extension-table-header";
const props = defineProps({
    modelValue: String,
});
const editor = new Editor({
    extensions: [
        StarterKit,
        Image,
        Youtube.configure({ width: 480, height: 270 }),
        Underline,
        TextStyle,
        Color,
        TextAlign.configure({ types: ['heading', 'paragraph'] }),
        FontFamily,
        FontSize,
        Table,
        TableRow,
        TableCell,
        TableHeader,

    ],
    content: props.modelValue,
});


const emit = defineEmits(['update:modelValue']);
onMounted(() => {
    document.body.classList.add('fullscreen-editor');
    editor.on('update', () => {
        emit('update:modelValue', editor.getHTML());
    });
});
onBeforeUnmount(() => {
    editor.destroy();
});

const addImage = () => {
    const url = prompt("Resim URL'si girin:");
    if (url) {
        editor.chain().focus().setImage({ src: url }).run();
    }
};

const addYoutube = () => {
    const url = prompt("YouTube Video URL'si girin:");
    if (url) {
        editor.chain().focus().setYoutubeVideo({ src: url }).run();
    }
};

const isActive = (type, attrs = {}) => {
    return type === "heading"
        ? editor?.isActive(type, { level: attrs.level })
        : editor?.isActive(type, attrs);
};

const setTextAlign = (alignment) => {
    editor.chain().focus().setTextAlign(alignment).run();
};

const setFontFamily = (family) => {
    editor.chain().focus().setFontFamily(family).run();
};

const setFontSize = (size) => {
    editor.chain().focus().setFontSize(size).run();
};



</script>

<template>

    <div class="editor-container ">
        <div class="toolbar dark:bg-zinc-500">
            <button :class="{ active: isActive('bold') }" @click="editor.chain().focus().toggleBold().run()">B</button>
            <button :class="{ active: isActive('italic') }"
                @click="editor.chain().focus().toggleItalic().run()">I</button>
            <button :class="{ active: isActive('underline') }"
                @click="editor.chain().focus().toggleUnderline().run()">U</button>
            <button :class="{ active: isActive('strike') }"
                @click="editor.chain().focus().toggleStrike().run()"><s>S</s></button>
            <button :class="{ active: isActive('code') }" @click="editor.chain().focus().toggleCode().run()">{}</button>

            <select @change="editor.chain().focus().toggleHeading({ level: parseInt($event.target.value) }).run()">
                <option :selected="isActive('heading', { level: 1 })" value="1">H1</option>
                <option :selected="isActive('heading', { level: 2 })" value="2">H2</option>
                <option :selected="isActive('heading', { level: 3 })" value="3">H3</option>
                <option :selected="isActive('heading', { level: 4 })" value="4">H4</option>
                <option :selected="isActive('heading', { level: 5 })" value="5">H5</option>
                <option :selected="isActive('heading', { level: 6 })" value="6">H6</option>
            </select>

            <button :class="{ active: isActive('bulletList') }"
                @click="editor.chain().focus().toggleBulletList().run()">•</button>
            <button :class="{ active: isActive('orderedList') }"
                @click="editor.chain().focus().toggleOrderedList().run()">1.</button>
            <button :class="{ active: isActive('blockquote') }"
                @click="editor.chain().focus().toggleBlockquote().run()">❝</button>
            <button @click="editor.chain().focus().setHorizontalRule().run()">---</button>
            <button :class="{ active: isActive('codeBlock') }" @click="editor.chain().focus().toggleCodeBlock().run()">{
                }</button>

            <button :class="{ active: isActive('textAlign', { align: 'left' }) }"
                @click="setTextAlign('left')">Left</button>
            <button :class="{ active: isActive('textAlign', { align: 'center' }) }"
                @click="setTextAlign('center')">Center</button>
            <button :class="{ active: isActive('textAlign', { align: 'right' }) }"
                @click="setTextAlign('right')">Right</button>
            <button :class="{ active: isActive('textAlign', { align: 'justify' }) }"
                @click="setTextAlign('justify')">Justify</button>

            <select @change="setFontFamily($event.target.value)">
                <option value="Arial">Arial</option>
                <option value="Times New Roman">Times New Roman</option>
                <option value="Courier New">Courier New</option>
            </select>
            <select @change="setFontSize($event.target.value)">
                <option value="12px">12px</option>
                <option value="14px">14px</option>
                <option value="16px">16px</option>
                <option value="18px">18px</option>
                <option value="20px">20px</option>
            </select>

            <button @click="addImage">🖼️</button>
            <button @click="addYoutube">▶️</button>

        </div>
        <div class="prose-xl dark:prose-invert">
            <EditorContent :editor="editor" class="editor-content" />

        </div>
    </div>
</template>

<style scoped>
.editor-container {
    border: 1px solid #ccc;
    padding: 10px;
    border-radius: 8px;
    background: white;
    max-width: 90%;
    margin: 20px auto;
    box-shadow: 0px 2px 5px rgba(0, 0, 0, 0.1);
}

.dark .editor-container {
    background: #1e1e1e;
    /* Koyu arka plan */
    border-color: #444;
    /* Daha koyu bir sınır */
    color: #ddd;
    /* Açık renkli metin */
}



.fullscreen-editor {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: 1000;
    background: white;
    padding: 20px;
}

.fullscreen-editor .editor-container {
    max-width: none;
    height: calc(100% - 40px);
}

.toolbar {
    display: flex;
    flex-wrap: wrap;
    gap: 6px;
    background: #f8f9fa;
    padding: 8px;
    border-radius: 5px;
    margin-bottom: 10px;
    border: 1px solid #ddd;
    align-items: center;
}

.dark .toolbar {
    display: flex;
    flex-wrap: wrap;
    gap: 6px;
    background: #27272a;
    padding: 8px;
    border-radius: 5px;
    margin-bottom: 10px;
    border: 1px solid #27272a;
    align-items: center;
}



button,
select {
    padding: 6px 12px;
    font-weight: bold;
    cursor: pointer;
    border: 1px solid #ccc;
    background: #ffffff;
    border-radius: 4px;
    transition: all 0.2s;
    font-size: 14px;
}

button:hover,
select:hover {
    background: #f0f0f0;
}

button.active {
    background: #007bff;
    color: white;
}





.dark button,
.dark select {
    padding: 6px 12px;
    font-weight: bold;
    cursor: pointer;
    border: 1px solid #18181b;
    background: #27272a;
    border-radius: 4px;
    transition: all 0.2s;
    font-size: 14px;
}

.dark button:hover,
.dark select:hover {
    background: #3f3f46;
}

button.active {
    background: #007bff;
    color: white;
}


.dark .editor-content {
    color: #eee;
    background: #222;
}

.dark .editor-content h1 {
    color: #fff;
    border-bottom: 2px solid #4A90E2;
}

.dark .editor-content blockquote {
    border-left-color: #4A90E2;
    background: #333;
    color: #bbb;
}


.editor-content {
    padding: 10px;
    font-size: 16px;
    line-height: 1.6;
}

.editor-content h1 {
    font-size: 2.5em;
    font-weight: bold;
    margin: 1em 0;
    color: #222;
    border-bottom: 2px solid #007bff;
    padding-bottom: 0.3em;
}

textarea {
    outline: none
}

.editor-content blockquote {
    position: relative;
    border-left: 4px solid #007bff;
    padding: 10px 15px;
    margin: 1em 0;
    font-style: italic;
    background: #f8f9fa;
    color: #555;
}

.editor-content img {
    max-width: 100%;
    height: auto;
    border-radius: 8px;
}

.editor-content table {
    width: 100%;
    border-collapse: collapse;
    margin: 1em 0;
}

.editor-content th,
.editor-content td {
    padding: 8px 12px;
    border: 1px solid #ddd;
    text-align: left;
}

.editor-content th {
    background-color: #f7f7f7;
    font-weight: bold;
}

.no-tailwind {
    all: revert !important;
}
</style>
