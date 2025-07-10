<template>
    <div class="container mx-auto max-w-7xl px-4 py-8 pb-32">
        <!-- Hero Section -->
        <div class="text-center mb-12">
            <h1 class="text-4xl font-bold text-gray-900 dark:text-white mb-4">
                {{ $t('pricing.header') }}
            </h1>
            <p class="text-xl text-gray-600 dark:text-gray-300 max-w-3xl mx-auto">
                {{ $t('pricing.description') }}
            </p>
        </div>

        <!-- Token Info Banner -->
        <div class="bg-gradient-to-r from-blue-600 to-purple-600 rounded-lg p-6 mb-12 text-white text-center">
            <div class="flex items-center justify-center mb-4">
                <h2 class="text-2xl font-bold">{{ $t('pricing.banner.title') }}</h2>
            </div>
            <p class="text-lg mb-4">{{ $t('pricing.banner.description') }}</p>
            <div class="grid md:grid-cols-3 gap-4 text-sm">
                <router-link to="/" class="bg-white/20 rounded-lg p-4 flex items-center justify-center">
                    <mdicon name="book-open" size="24" class="mr-3 flex-shrink-0" />
                    <p class="text-center">{{ $t('pricing.banner.use1') }}</p>
                </router-link>
                <router-link to="/" class="bg-white/20 rounded-lg p-4 flex items-center justify-center">
                    <mdicon name="puzzle" size="24" class="mr-3 flex-shrink-0" />
                    <p class="text-center">{{ $t('pricing.banner.use2') }}</p>
                </router-link>
                <router-link to="/" class="bg-white/20 rounded-lg p-4 flex items-center justify-center">
                    <mdicon name="account-voice" size="24" class="mr-3 flex-shrink-0" />
                    <p class="text-center">{{ $t('pricing.banner.use3') }}</p>
                </router-link>
            </div>
        </div>

        <!-- Security Banner -->
        <div
            class="bg-green-50 dark:bg-green-900/20 border border-green-200 dark:border-green-800 rounded-lg p-4 mb-8 text-center">
            <div class="flex items-center justify-center mb-2">
                <mdicon name="shield-check" size="24" class="text-green-600 dark:text-green-400 mr-2" />
                <span class="text-green-800 dark:text-green-200 font-medium">
                    {{ $t('pricing.purchase.title') }}
                </span>
            </div>
            <p class="text-sm text-green-700 dark:text-green-300">
                {{ $t('pricing.purchase.description') }}
            </p>
        </div>

        <!-- Quick Buy Options -->
        <div class="grid md:grid-cols-4 gap-6 mb-12">
            <!-- Quick Buy Cards -->
            <div v-for="option in quickBuyOptions" :key="option.tokens"
                class="bg-white dark:bg-gray-800 rounded-lg shadow-lg p-6 border-2 border-gray-200 dark:border-gray-700 hover:shadow-xl hover:border-purple-500 dark:hover:border-purple-400 transition-all cursor-pointer"
                :class="{ 'border-purple-500 dark:border-purple-400 bg-purple-50 dark:bg-purple-900/20': selectedPackage.tokens === option.tokens }"
                @click="selectTokens(option.tokens)">
                <div class="text-center">
                    <div class="flex items-center justify-center mb-3">
                        <div class="text-4xl font-bold text-gray-900 dark:text-white">
                            {{ option.tokens }}
                        </div>
                    </div>
                    <div class="text-sm text-gray-600 dark:text-gray-300 mb-2">
                        {{ $t('pricing.label') }}
                    </div>
                    <div class="text-xl font-semibold text-purple-600 dark:text-purple-400 mb-2">
                        ${{ calculatePrice(option.tokens) }}
                    </div>
                    <div v-if="option.discount > 0" class="text-xs text-green-600 dark:text-green-400">
                        {{ option.discount }}% {{ $t('pricing.discount') }}
                    </div>
                </div>
            </div>
        </div>

        <!-- Custom Token Calculator -->
        <div id="calculator-summary" class="bg-gray-50 dark:bg-gray-900 rounded-lg p-8 mb-12">
            <h3 class="text-2xl font-bold text-gray-900 dark:text-white mb-6 text-center">
                {{ $t('pricing.calculator.title') }}
            </h3>
            <div class="max-w-md mx-auto">
                <div class="mb-6 flex items-center space-x-2">
                    <label class="w-full block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                        {{ $t('pricing.calculator.token_count') }}
                    </label>
                    <div class="relative w-full">
                        <input v-model.number="customTokens" type="number" min="1" max="1000" step="1"
                            class="w-full px-4 py-3 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-transparent dark:bg-gray-800 dark:text-white text-center text-lg"
                            :placeholder="$t('pricing.calculator.placeholder')" @input="handleCustomTokenChange" />
                    </div>
                </div>

                <div class="text-center mb-6">
                    <div class="text-4xl font-bold text-gray-900 dark:text-white mb-2">
                        ${{ calculatePrice(customTokens) }}
                    </div>
                    <p class="text-sm text-gray-600 dark:text-gray-300">
                        {{ getDiscountText(customTokens) }}
                    </p>
                </div>

                <div class="bg-white dark:bg-gray-800 rounded-lg p-4 mb-6">
                    <div class="text-sm text-gray-600 dark:text-gray-300">
                        <div class="flex justify-between mb-2">
                            <span>{{ $t('pricing.calculator.base_price') }}:</span>
                            <span>${{ (Math.ceil(Number(customTokens) || 0) * 1.00).toFixed(2) }}</span>
                        </div>
                        <div v-if="getDiscount(customTokens) > 0"
                            class="flex justify-between mb-2 text-green-600 dark:text-green-400">
                            <span>{{ $t('pricing.calculator.discount') }}:</span>
                            <span>-${{ ((Math.ceil(Number(customTokens) || 0) * 1.00) -
                                calculateTokenPrice(customTokens)).toFixed(2) }}</span>
                        </div>
                        <hr class="my-2 border-gray-200 dark:border-gray-600">
                        <div class="flex justify-between font-semibold text-gray-900 dark:text-white">
                            <span>{{ $t('pricing.calculator.total') }}:</span>
                            <span>${{ calculatePrice(customTokens) }}</span>
                        </div>
                    </div>
                </div>

                <div class="text-center">
                    <p class="text-sm text-gray-600 dark:text-gray-300">
                        {{ $t('pricing.calculator.purchase_instruction') }}
                    </p>
                </div>
            </div>
        </div>

        <!-- Pricing Tiers Info -->
        <div class="bg-white dark:bg-gray-800 rounded-lg p-8 mb-12">
            <h3 class="text-2xl font-bold text-gray-900 dark:text-white mb-6 text-center">
                {{ $t('pricing.tiers.title') }}
            </h3>
            <div class="grid md:grid-cols-3 gap-6">
                <div class="text-center p-4 border border-gray-200 dark:border-gray-700 rounded-lg">
                    <div class="text-3xl font-bold text-gray-900 dark:text-white mb-2">$1.00</div>
                    <div class="text-sm text-gray-600 dark:text-gray-300 mb-2">{{ $t('pricing.tiers.per_token') }}</div>
                    <div class="text-xs text-gray-500 dark:text-gray-400">1-49 {{ $t('pricing.label') }}</div>
                </div>
                <div class="text-center p-4 border-2 border-green-500 rounded-lg relative">
                    <div class="absolute -top-2 left-1/2 transform -translate-x-1/2">
                        <span class="bg-green-500 text-white px-3 py-1 rounded-full text-xs font-semibold">10%
                            OFF</span>
                    </div>
                    <div class="text-3xl font-bold text-green-600 dark:text-green-400 mb-2">$0.90</div>
                    <div class="text-sm text-gray-600 dark:text-gray-300 mb-2">{{ $t('pricing.tiers.per_token') }}</div>
                    <div class="text-xs text-gray-500 dark:text-gray-400">50-99 {{ $t('pricing.label') }}</div>
                </div>
                <div class="text-center p-4 border-2 border-purple-500 rounded-lg relative">
                    <div class="absolute -top-2 left-1/2 transform -translate-x-1/2">
                        <span class="bg-purple-500 text-white px-3 py-1 rounded-full text-xs font-semibold">20%
                            OFF</span>
                    </div>
                    <div class="text-3xl font-bold text-purple-600 dark:text-purple-400 mb-2">$0.80</div>
                    <div class="text-sm text-gray-600 dark:text-gray-300 mb-2">{{ $t('pricing.tiers.per_token') }}</div>
                    <div class="text-xs text-gray-500 dark:text-gray-400">100+ {{ $t('pricing.label') }}</div>
                </div>
            </div>
        </div>

        <!-- FAQ Section -->
        <div class="mb-12">
            <h3 class="text-2xl font-bold text-gray-900 dark:text-white mb-8 text-center">
                {{ $t('pricing.faq.title') }}
            </h3>
            <div class="space-y-4 max-w-3xl mx-auto">
                <div v-for="(faq, index) in faqs" :key="index"
                    class="border border-gray-200 dark:border-gray-700 rounded-lg">
                    <button @click="toggleFaq(index)"
                        class="w-full px-6 py-4 text-left flex justify-between items-center hover:bg-gray-50 dark:hover:bg-gray-800 rounded-lg transition-colors">
                        <span class="font-semibold text-gray-900 dark:text-white">{{ faq.question }}</span>
                        <mdicon :name="faq.open ? 'chevron-up' : 'chevron-down'" size="20" class="text-gray-500" />
                    </button>
                    <div v-if="faq.open" class="px-6 pb-4">
                        <p class="text-gray-700 dark:text-gray-300">{{ faq.answer }}</p>
                    </div>
                </div>
            </div>
        </div>

        <!-- CheckoutModal artık kullanılmıyor - Gumroad entegrasyonu -->
        <!--
        <CheckoutModal 
            :isVisible="showCheckoutModal" 
            :packageInfo="selectedPackage"
            @close="closeCheckoutModal"
            @success="handlePurchaseSuccess"
        />
        -->

        <!-- Fixed Buy Button -->
        <div
            class="fixed bottom-0 left-0 right-0 bg-white dark:bg-gray-800 border-t border-gray-200 dark:border-gray-700 p-4 z-40">
            <div class="container mx-auto max-w-md">
                <div v-if="selectedPackage.tokens" class="text-center mb-3">
                    <div class="text-lg font-semibold text-gray-900 dark:text-white">
                        {{ selectedPackage.tokens }} {{ $t('pricing.label') }} - ${{ selectedPackage.price }}
                    </div>
                    <div v-if="selectedPackage.discount > 0" class="text-sm text-green-600 dark:text-green-400">
                        {{ selectedPackage.discount }}% {{ $t('pricing.discount') }}
                    </div>
                </div>
                <button @click="proceedToPurchase" :disabled="!selectedPackage.tokens"
                    class="smooth-click w-full bg-gradient-to-r from-purple-600 to-pink-600 hover:from-purple-700 hover:to-pink-700 disabled:from-gray-400 disabled:to-gray-500 text-white font-semibold py-4 px-6 rounded-lg transition duration-200 transform disabled:transform-none disabled:cursor-not-allowed flex items-center justify-center">
                    <mdicon name="credit-card" size="20" class="mr-2" />
                    {{ selectedPackage.tokens ? $t('pricing.buy_now') : $t('pricing.calculator.buy_tokens') }}
                </button>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useToast } from '@/composable/useToast'
import { useI18n } from 'vue-i18n'
import { useFetch } from '@/composable/useFetch'
const { t } = useI18n()
const toast = useToast()

// Reactive data
const customTokens = ref(50)
const selectedPackage = ref({})
const selectedProvider = ref('')
const selectedCurrency = ref('')

// Initialize with custom token value and load providers
onMounted(async () => {
    handleCustomTokenChange()
})

// Quick buy options
const quickBuyOptions = ref([
    { tokens: 25, discount: 0 },
    { tokens: 50, discount: 10 },
    { tokens: 100, discount: 20 },
    { tokens: 200, discount: 20 }
])

// FAQ data
const faqs = ref([
    {
        question: computed(() => t('pricing.faq.q1')),
        answer: computed(() => t('pricing.faq.a1')),
        open: false
    },
    {
        question: computed(() => t('pricing.faq.q2')),
        answer: computed(() => t('pricing.faq.a2')),
        open: false
    },
    {
        question: computed(() => t('pricing.faq.q3')),
        answer: computed(() => t('pricing.faq.a3')),
        open: false
    },
    {
        question: computed(() => t('pricing.faq.q4')),
        answer: computed(() => t('pricing.faq.a4')),
        open: false
    }
])



function calculateTokenPrice(tokenCount) {
    const integerTokens = Math.ceil(Number(tokenCount) || 0)

    if (!integerTokens || integerTokens < 1) return 0

    if (integerTokens < 50) {
        return integerTokens * 1.00
    } else if (integerTokens < 100) {
        return integerTokens * 0.90
    } else {
        return integerTokens * 0.80
    }
}

function calculatePrice(tokens, withDiscount = true) {
    const integerTokens = Math.ceil(Number(tokens) || 0)

    if (!integerTokens || integerTokens < 1) return '0.00'

    const price = withDiscount ? calculateTokenPrice(integerTokens) : integerTokens * 1.00
    return price.toFixed(2)
}

function getDiscount(tokens) {
    const integerTokens = Math.ceil(Number(tokens) || 0)

    if (integerTokens < 50) return 0
    if (integerTokens < 100) return 10
    return 20
}

function getDiscountText(tokens) {
    const integerTokens = Math.ceil(Number(tokens) || 0)
    const discount = getDiscount(integerTokens)

    if (discount === 0) {
        return t('pricing.calculator.no_discount')
    } else {
        return t('pricing.calculator.discount', { discount: discount + "%" })
    }
}

function toggleFaq(index) {
    faqs.value[index].open = !faqs.value[index].open
}

function selectTokens(tokens) {
    const integerTokens = Math.ceil(Number(tokens) || 0)

    if (!integerTokens || integerTokens < 1) return

    const discount = getDiscount(integerTokens)

    selectedPackage.value = {
        tokens: integerTokens,
        price: calculatePrice(integerTokens),
        discount
    }
    customTokens.value = integerTokens

    const summaryElement = document.getElementById('calculator-summary')
    if (summaryElement) {
        summaryElement.scrollIntoView({ behavior: 'smooth' })
    }
}

function handleCustomTokenChange() {
    const integerTokens = Math.ceil(Number(customTokens.value) || 0)
    customTokens.value = integerTokens

    if (integerTokens && integerTokens >= 1) {
        const discount = getDiscount(integerTokens)

        selectedPackage.value = {
            tokens: integerTokens,
            price: calculatePrice(integerTokens),
            discount
        }
    } else {
        selectedPackage.value = {}
    }
}

async function proceedToPurchase() {
    const res = await useFetch("/payment/token", {
        method: "POST",
        body: {
            token_count: selectedPackage.value.tokens,
        }
    }
    );
}

</script>
