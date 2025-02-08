<template>
    <div>
        <div class=" flex justify-center items-center space-x-4">
            <h1 class="font-bold text-4xl">
                {{ response.title }}
            </h1>
        </div>

        <div class="p-5 text-center">
            {{
                $t('list_helper', {
                    lang: $t(response.lang),
                    s: response.target_langs.length > 1 ? 'ler' : '',
                    target_lang: response.target_langs.map(x => ' ' + $t(x)).toString()
                })
            }}
        </div>
        <div class="max-w-160 w-full mx-auto flex justify-around items-center py-5">
            <button class="smooth-click">
                <mdicon name="star-outline" class="dark:text-yellow-400 text-yellow-500" size="32" />
            </button>
            <router-link :to="'/l/' + $route.params.id + '/games'" class="smooth-click">
                <mdicon name="gamepad-variant-outline" size="32" />
            </router-link>
        </div>

        <div class="flex py-5 justify-center">
            <input v-model="search" type="text" :placeholder="$t('search_from_list')" class="w-full p-3 border rounded-lg shadow-sm outline-none transition-all 
             bg-white text-zinc-900  border-none
             max-w-160 
             dark:bg-zinc-800 dark:text-white " />
        </div>
        <div v-auto-animate class="space-y-5">

            <div v-for="item in filteredList" :key="item.id" class="flex justify-center">
                <div class="max-w-160 w-full ">
                    <div v-auto-animate class="card transition duration-200  hover:shadow p-4">

                        <div class="font-bold text-xl capitalize ">{{ item.word }}</div>
                        <div class="font-light pt-5">{{ item.description }}</div>

                        <div class="languages">
                            <div v-if="item.example.length > 0">
                                <div v-for="example in item.example" class="p-5 flex items-end font-light text-sm">
                                    <mdicon name="arrow-right" />
                                    <span>{{ example }}</span>
                                </div>
                            </div>
                            <div class="text-sm text-end text-sky-900 dark:text-sky-200 "> {{ $t(item.lang) }}</div>
                            <hr class="border-t-2 border-zinc-200 dark:border-zinc-800 my-4 opacity-50">
                        </div>
                        <div class="languages" v-for="(sub, i) in item.languages" :key="i">
                            <div class="font-bold text-xl capitalize py-5">{{ sub.word }}</div>
                            <div class="font-light">{{ sub.description }}</div>
                            <div v-if="sub.example.length > 0">
                                <div v-for="example in sub.example" class="p-5 flex items-end font-light text-sm">
                                    <mdicon name="arrow-right" />
                                    <span>{{ example }}</span>
                                </div>
                            </div>
                            <div class="text-sm text-end text-sky-900 dark:text-sky-200 "> {{ $t(sub.lang) }}</div>
                            <hr v-if="i < item.languages.length - 1"
                                class="border-t-2 border-zinc-200 dark:border-zinc-800 my-4 opacity-50">
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { computed, ref } from 'vue';

const search = ref('');

const filteredList = computed(() => {
    return response.value.list.filter(x => x.word.toLowerCase().includes(search.value.toLowerCase()));
});

const response = ref({
    id: 1,
    title: 'Sample List',
    lang: 'en',
    target_langs: ['tr'],
    list: [
        {
            id: 1,
            lang: 'en',
            word: 'title',
            description: 'A title is a word or phrase that describes something, such as a book or article.',
            example: ['The title of the book is very inspiring.'],
            languages: [
                {
                    id: 'a',
                    lang: 'tr',
                    word: 'başlık',
                    description: 'Başlık, bir şeyi tanımlayan kelime ya da ifadedir, örneğin bir kitap ya da makale başlığı.',
                    example: ['Bu başlık dikkat çekici olmalı.'],
                    base_id: 1,
                }
            ]
        },
        {
            id: 2,
            lang: 'en',
            word: 'book',
            description: 'A set of written or printed pages, usually bound with a protective cover.',
            example: ['I am reading a book about history.'],
            languages: [
                {
                    id: 'b',
                    lang: 'tr',
                    word: 'kitap',
                    description: 'Yazılı ya da basılı sayfalardan oluşan, genellikle koruyucu bir kapakla bağlanmış eser.',
                    example: ['Tarih hakkında bir kitap okuyorum.'],
                    base_id: 2,
                }
            ]
        },
        {
            id: 3,
            lang: 'en',
            word: 'computer',
            description: 'An electronic device for storing and processing data.',
            example: ['I use my computer to work from home.'],
            languages: [
                {
                    id: 'c',
                    lang: 'tr',
                    word: 'bilgisayar',
                    description: 'Veri depolama ve işleme için kullanılan elektronik cihaz.',
                    example: ['Evden çalışmak için bilgisayarımı kullanıyorum.'],
                    base_id: 3,
                }
            ]
        },
        {
            id: 4,
            lang: 'en',
            word: 'apple',
            description: 'A round fruit with red or green skin and a whitish interior.',
            example: ['I ate an apple for breakfast.'],
            languages: [
                {
                    id: 'd',
                    lang: 'tr',
                    word: 'elma',
                    description: 'Kırmızı veya yeşil kabuğu olan, beyaz içi bulunan yuvarlak meyve.',
                    example: ['Kahvaltıda bir elma yedim.'],
                    base_id: 4,
                }
            ]
        },
        {
            id: 5,
            lang: 'en',
            word: 'car',
            description: 'A road vehicle powered by an engine, typically with four wheels.',
            example: ['She drives a red car.'],
            languages: [
                {
                    id: 'e',
                    lang: 'tr',
                    word: 'araba',
                    description: 'Bir motor tarafından güçlendirilen, tipik olarak dört tekerleği olan kara aracı.',
                    example: ['Kırmızı bir araba kullanıyor.'],
                    base_id: 5,
                }
            ]
        },
        {
            id: 6,
            lang: 'en',
            word: 'city',
            description: 'A large town or urban area with a lot of buildings and people.',
            example: ['I live in a large city.'],
            languages: [
                {
                    id: 'f',
                    lang: 'tr',
                    word: 'şehir',
                    description: 'Binalar ve insanlar bakımından yoğun olan büyük kasaba veya kentsel alan.',
                    example: ['Büyük bir şehirde yaşıyorum.'],
                    base_id: 6,
                }
            ]
        },
        {
            id: 7,
            lang: 'en',
            word: 'friend',
            description: 'A person whom one knows and with whom one has a bond of mutual affection.',
            example: ['My best friend is coming over.'],
            languages: [
                {
                    id: 'g',
                    lang: 'tr',
                    word: 'arkadaş',
                    description: 'Bir kişiyle tanışan ve karşılıklı sevgi bağına sahip olunan kişi.',
                    example: ['En iyi arkadaşım geliyor.'],
                    base_id: 7,
                }
            ]
        },
        {
            id: 8,
            lang: 'en',
            word: 'house',
            description: 'A building for human habitation.',
            example: ['We live in a house with a large garden.'],
            languages: [
                {
                    id: 'h',
                    lang: 'tr',
                    word: 'ev',
                    description: 'İnsanların barınması için kullanılan bina.',
                    example: ['Büyük bir bahçesi olan bir evde yaşıyoruz.'],
                    base_id: 8,
                }
            ]
        },
        {
            id: 9,
            lang: 'en',
            word: 'school',
            description: 'A place where children are educated.',
            example: ['My children go to school every day.'],
            languages: [
                {
                    id: 'i',
                    lang: 'tr',
                    word: 'okul',
                    description: 'Çocukların eğitim aldığı yer.',
                    example: ['Çocuklarım her gün okula gidiyor.'],
                    base_id: 9,
                }
            ]
        },
        {
            id: 10,
            lang: 'en',
            word: 'dog',
            description: 'A domesticated carnivorous mammal.',
            example: ['I have a pet dog.'],
            languages: [
                {
                    id: 'j',
                    lang: 'tr',
                    word: 'köpek',
                    description: 'Evcil etobur memeli.',
                    example: ['Bir evcil köpeğim var.'],
                    base_id: 10,
                }
            ]
        }
    ]
});
</script>
<style scoped>
.languages {
    opacity: 0;
    visibility: hidden;
    height: 0;
    overflow: hidden;
    transition: opacity 0.3s ease, visibility 0s linear 0.3s, height 0.3s ease;
}

.card:hover .languages {
    opacity: 1;
    visibility: visible;
    height: auto;
    transition: opacity 0.3s ease, visibility 0s linear 0s, height 0.3s ease;
}
</style>