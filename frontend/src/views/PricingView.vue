<template>
    <div class="container mx-auto max-w-7xl px-4 py-8 pb-32">
        <!-- Hero Section -->
        <div class="text-center mb-12">
            <h1 class="text-4xl font-bold text-gray-900 dark:text-white mb-4">
                {{ $t('pricing.title') }}
            </h1>
            <p class="text-xl text-gray-600 dark:text-gray-300 max-w-3xl mx-auto">
                {{ $t('pricing.description') }}
            </p>
        </div>

        <!-- Token Info Banner -->
        <div class="bg-gradient-to-r from-blue-600 to-purple-600 rounded-lg p-6 mb-12 text-white text-center">
            <div class="flex items-center justify-center mb-4">
                <h2 class="text-2xl font-bold">{{ $t('pricing.tokens.title') }}</h2>
            </div>
            <p class="text-lg mb-4">{{ $t('pricing.tokens.description') }}</p>
            <div class="grid md:grid-cols-3 gap-4 text-sm">
                <div class="bg-white/20 rounded-lg p-4 flex items-center justify-center">
                    <mdicon name="book-open" size="24" class="mr-3 flex-shrink-0" />
                    <p class="text-center">{{ $t('pricing.tokens.use1') }}</p>
                </div>
                <div class="bg-white/20 rounded-lg p-4 flex items-center justify-center">
                    <mdicon name="gamepad-variant-outline" size="24" class="mr-3 flex-shrink-0" />
                    <p class="text-center">{{ $t('pricing.tokens.use2') }}</p>
                </div>
                <div class="bg-white/20 rounded-lg p-4 flex items-center justify-center">
                    <mdicon name="account-voice" size="24" class="mr-3 flex-shrink-0" />
                    <p class="text-center">{{ $t('pricing.tokens.use3') }}</p>
                </div>
            </div>
        </div>

        <!-- Security Banner -->
        <div
            class="bg-green-50 dark:bg-green-900/20 border border-green-200 dark:border-green-800 rounded-lg p-4 mb-8 text-center">
            <div class="flex items-center justify-center mb-2">
                <mdicon name="shield-check" size="24" class="text-green-600 dark:text-green-400 mr-2" />
                <span class="text-green-800 dark:text-green-200 font-medium">
                    {{ $t('pricing.purchase.secure_payment') }}
                </span>
            </div>
            <p class="text-sm text-green-700 dark:text-green-300">
                {{ $t('pricing.purchase.secure_description') }}
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
                        {{ $t('pricing.tokens.label') }}
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

        <!-- Payment Provider Selection -->
        <div id="calculator-summary" v-if="availableProviders.length > 0"
            class="bg-white dark:bg-gray-800 rounded-lg p-6 mb-8">
            <h3 class="text-xl font-bold text-gray-900 dark:text-white mb-4 text-center">
                {{ $t('pricing.payment.choose_provider') }}
            </h3>

            <!-- Loading State -->
            <div v-if="isLoadingProviders" class="text-center py-8">
                <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-purple-600 mx-auto"></div>
                <p class="text-gray-500 mt-2">{{ $t('pricing.payment.loading') }}</p>
            </div>

            <!-- Provider Selection -->
            <div v-else class="max-w-2xl mx-auto">
                <!-- Country Detection Info -->
                <div v-if="detectedCountry" class="bg-blue-50 dark:bg-blue-900/20 rounded-lg p-3 mb-4 text-center">
                    <div class="flex items-center justify-center text-sm text-blue-700 dark:text-blue-300">
                        <mdicon name="map-marker" size="16" class="mr-2" />
                        <span>{{ $t('pricing.payment.detected_location') }}: {{ detectedCountry }}</span>
                    </div>
                </div>

                <!-- Provider Cards -->
                <div class="grid md:grid-cols-2 gap-4">
                    <div v-for="provider in availableProviders" :key="provider.provider"
                        class="border-2 rounded-lg p-4 cursor-pointer transition-all hover:shadow-lg" :class="selectedProvider === provider.provider
                            ? 'border-purple-500 bg-purple-50 dark:bg-purple-900/20 shadow-md'
                            : 'border-gray-200 dark:border-gray-700 hover:border-gray-300 dark:hover:border-gray-600'"
                        @click="selectProvider(provider.provider, provider.currency)">

                        <div class="flex items-center justify-between mb-3">
                            <div class="flex items-center">
                                <!-- Provider Icon -->
                                <div class="w-10 h-10 rounded-full flex items-center justify-center mr-3"
                                    :class="provider.provider === 'iyzico' ? 'bg-orange-100 dark:bg-orange-900/20' : 'bg-blue-100 dark:bg-blue-900/20'">
                                    <mdicon name="credit-card" size="20"
                                        :class="provider.provider === 'iyzico' ? 'text-orange-600 dark:text-orange-400' : 'text-blue-600 dark:text-blue-400'" />
                                </div>

                                <div>
                                    <div class="font-semibold text-gray-900 dark:text-white">
                                        {{ provider.display_name || (provider.provider === 'iyzico' ? 'İyzico' :
                                        'PayPal') }}
                                    </div>
                                    <div class="text-sm text-gray-600 dark:text-gray-300">
                                        {{ provider.description || provider.currency }}
                                    </div>
                                </div>
                            </div>

                            <!-- Selection Indicator -->
                            <div v-if="selectedProvider === provider.provider"
                                class="w-6 h-6 rounded-full bg-purple-500 flex items-center justify-center">
                                <mdicon name="check" size="16" class="text-white" />
                            </div>
                        </div>

                        <!-- Currency and Recommended Badge -->
                        <div class="flex items-center justify-between">
                            <div class="text-sm font-medium text-gray-700 dark:text-gray-300">
                                {{ provider.currency }}
                            </div>
                            <div v-if="provider.recommended" class="flex items-center">
                                <span
                                    class="text-xs bg-green-100 dark:bg-green-900/20 text-green-800 dark:text-green-400 px-2 py-1 rounded-full font-medium flex items-center">
                                    <mdicon name="star" size="12" class="mr-1" />
                                    <span>{{ $t('pricing.payment.recommended') }}</span>
                                </span>
                            </div>
                        </div>
                        <!-- Reason -->
                        <div class="text-xs text-gray-500 dark:text-gray-400 mt-2">
                            {{ provider.reason }}
                        </div>
                    </div>
                </div>

                <!-- Selected Provider Summary -->
                <div v-if="selectedProvider" class="mt-4 p-3 bg-gray-50 dark:bg-gray-900 rounded-lg">
                    <div class="flex items-center text-sm text-gray-700 dark:text-gray-300">
                        <mdicon name="information" size="16" class="mr-2 text-blue-500" />
                        <span>{{ $t('pricing.payment.selected_provider') }}: </span>
                        <span class="font-semibold ml-1">
                            {{ selectedProvider === 'iyzico' ? 'İyzico' : 'PayPal' }} ({{ selectedCurrency }})
                        </span>
                    </div>
                </div>
            </div>
        </div>

        <!-- Custom Token Calculator -->
        <div class="bg-gray-50 dark:bg-gray-900 rounded-lg p-8 mb-12">
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
                    <div class="text-xs text-gray-500 dark:text-gray-400">1-49 {{ $t('pricing.tokens.label') }}</div>
                </div>
                <div class="text-center p-4 border-2 border-green-500 rounded-lg relative">
                    <div class="absolute -top-2 left-1/2 transform -translate-x-1/2">
                        <span class="bg-green-500 text-white px-3 py-1 rounded-full text-xs font-semibold">10%
                            OFF</span>
                    </div>
                    <div class="text-3xl font-bold text-green-600 dark:text-green-400 mb-2">$0.90</div>
                    <div class="text-sm text-gray-600 dark:text-gray-300 mb-2">{{ $t('pricing.tiers.per_token') }}</div>
                    <div class="text-xs text-gray-500 dark:text-gray-400">50-99 {{ $t('pricing.tokens.label') }}</div>
                </div>
                <div class="text-center p-4 border-2 border-purple-500 rounded-lg relative">
                    <div class="absolute -top-2 left-1/2 transform -translate-x-1/2">
                        <span class="bg-purple-500 text-white px-3 py-1 rounded-full text-xs font-semibold">20%
                            OFF</span>
                    </div>
                    <div class="text-3xl font-bold text-purple-600 dark:text-purple-400 mb-2">$0.80</div>
                    <div class="text-sm text-gray-600 dark:text-gray-300 mb-2">{{ $t('pricing.tiers.per_token') }}</div>
                    <div class="text-xs text-gray-500 dark:text-gray-400">100+ {{ $t('pricing.tokens.label') }}</div>
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
                        {{ selectedPackage.tokens }} {{ $t('pricing.tokens.label') }} - ${{ selectedPackage.price }}
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
// CheckoutModal artık kullanılmıyor - Gumroad entegrasyonu
// import CheckoutModal from '@/components/CheckoutModal.vue'

const { t } = useI18n()
const toast = useToast()

// Reactive data
const customTokens = ref(50)
const selectedPackage = ref({})
const availableProviders = ref([])
const selectedProvider = ref('')
const selectedCurrency = ref('')
const detectedCountry = ref('')
const isLoadingProviders = ref(false)

// Initialize with custom token value and load providers
onMounted(async () => {
    handleCustomTokenChange()
    await loadAvailableProviders()
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

// Load available payment providers
async function loadAvailableProviders() {
    try {
        isLoadingProviders.value = true

        // Mock payment providers for UI testing
        const mockProviders = [
            {
                provider: 'iyzico',
                currency: 'TRY',
                recommended: true,
                reason: 'Türkiye için en uygun ödeme yöntemi',
                display_name: 'İyzico',
                description: 'Türkiye\'nin güvenilir ödeme sistemi'
            },
            {
                provider: 'paypal',
                currency: 'USD',
                recommended: false,
                reason: 'Uluslararası ödemeler için ideal',
                display_name: 'PayPal',
                description: 'Dünya çapında güvenli ödeme'
            }
        ]

        // Simulate API call delay
        await new Promise(resolve => setTimeout(resolve, 500))

        // Try to load from real API first, fallback to mock
        try {
            const response = await useFetch('/user/payment-providers')

            if (response && response.success) {
                availableProviders.value = response.providers || mockProviders
                selectedProvider.value = response.default_provider || 'iyzico'
                selectedCurrency.value = response.default_currency || 'TRY'
                detectedCountry.value = response.detected_country || 'TR'

                console.log('✅ Real API - Available providers:', availableProviders.value)
            } else {
                throw new Error('API returned no data')
            }
        } catch (apiError) {
            console.warn('⚠️ API failed, using mock data:', apiError.message)

            // Use mock data as fallback
            availableProviders.value = mockProviders
            selectedProvider.value = 'iyzico'
            selectedCurrency.value = 'TRY'
            detectedCountry.value = 'TR'

            console.log('🔄 Mock - Available providers:', availableProviders.value)
        }

        console.log('Selected provider:', selectedProvider.value, selectedCurrency.value)
    } catch (error) {
        console.error('Error loading payment providers:', error)
        // Ultimate fallback
        availableProviders.value = []
        selectedProvider.value = 'iyzico'
        selectedCurrency.value = 'TRY'
    } finally {
        isLoadingProviders.value = false
    }
}

// Select payment provider
function selectProvider(provider, currency) {
    selectedProvider.value = provider
    selectedCurrency.value = currency

    // Recalculate prices when currency changes
    if (selectedPackage.value.tokens) {
        handleCustomTokenChange()
    }

    console.log('Selected provider:', provider, currency)
}

// Token price calculation function (matching the Go logic)
function calculateTokenPrice(tokenCount) {
    // Float değerleri üste yuvarla ve integer'a çevir
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

// Calculate price with formatting
function calculatePrice(tokens, withDiscount = true) {
    // Float değerleri üste yuvarla ve integer'a çevir
    const integerTokens = Math.ceil(Number(tokens) || 0)

    if (!integerTokens || integerTokens < 1) return '0.00'

    const price = withDiscount ? calculateTokenPrice(integerTokens) : integerTokens * 1.00
    return price.toFixed(2)
}

// Get discount percentage
function getDiscount(tokens) {
    // Float değerleri üste yuvarla ve integer'a çevir
    const integerTokens = Math.ceil(Number(tokens) || 0)

    if (integerTokens < 50) return 0
    if (integerTokens < 100) return 10
    return 20
}

// Get discount text
function getDiscountText(tokens) {
    // Float değerleri üste yuvarla ve integer'a çevir
    const integerTokens = Math.ceil(Number(tokens) || 0)
    const discount = getDiscount(integerTokens)

    if (discount === 0) {
        return t('pricing.calculator.no_discount')
    } else {
        return t('pricing.calculator.discount', { discount })
    }
}

// Toggle FAQ
function toggleFaq(index) {
    faqs.value[index].open = !faqs.value[index].open
}

// Select tokens for purchase
function selectTokens(tokens) {
    // Float değerleri üste yuvarla ve integer'a çevir
    const integerTokens = Math.ceil(Number(tokens) || 0)

    if (!integerTokens || integerTokens < 1) return

    const discount = getDiscount(integerTokens)

    selectedPackage.value = {
        tokens: integerTokens,
        price: calculatePrice(integerTokens),
        discount
    }
    customTokens.value = integerTokens

    // Scroll to calculator-summary id smoothly
    const summaryElement = document.getElementById('calculator-summary')
    if (summaryElement) {
        summaryElement.scrollIntoView({ behavior: 'smooth' })
    }
}

// Handle custom token input change
function handleCustomTokenChange() {
    // Float değerleri üste yuvarla ve integer'a çevir
    const integerTokens = Math.ceil(Number(customTokens.value) || 0)

    // Input değerini de integer'a çevir
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

// Proceed to backend purchase
async function proceedToPurchase() {
    // Float değerleri üste yuvarla ve integer'a çevir
    const integerTokens = Math.ceil(Number(selectedPackage.value.tokens) || 0)

    if (!integerTokens) return

    try {
        // For demo purposes, create a mock response
        const mockResponse = {
            success: true,
            payment_url: selectedProvider.value === 'iyzico'
                ? 'https://sandbox-cpp.iyzipay.com/?token=demo-token-12345'
                : 'https://www.sandbox.paypal.com/checkoutnow?token=demo-paypal-token-67890',
            transaction_id: 'txn_' + Date.now(),
            amount: calculatePrice(integerTokens),
            currency: selectedCurrency.value,
            tokens: integerTokens,
            provider: selectedProvider.value
        }

        // Simulate API delay
        await new Promise(resolve => setTimeout(resolve, 500))

        // Try real API first, fallback to mock
        let response
        try {
            response = await useFetch('/user/purchase-tokens', {
                method: 'POST',
                body: {
                    tokens: integerTokens,
                    provider: selectedProvider.value,
                    currency: selectedCurrency.value
                },
                headers: {
                    'Content-Type': 'application/json'
                }
            })

            if (!response || !response.success) {
                throw new Error(response?.error || 'API failed')
            }

            console.log('✅ Real API Response:', response)
        } catch (apiError) {
            console.warn('⚠️ API failed, using mock response:', apiError.message)
            response = mockResponse
            console.log('🔄 Mock Response:', response)
        }

        if (response && response.success) {
            // Show the payment URL in an alert for demo purposes
            if (response.payment_url.includes('demo-token') || response.payment_url.includes('demo-paypal')) {
                // This is a demo URL
                const providerName = selectedProvider.value === 'iyzico' ? 'İyzico' : 'PayPal'
                const message = `DEMO: ${providerName} ödeme sayfasına yönlendiriliyorsunuz...\n\nURL: ${response.payment_url}\n\nGerçek ortamda bu URL otomatik olarak açılacaktır.`
                alert(message)

                // For demo, simulate successful payment after 3 seconds
                setTimeout(() => {
                    handleMockPaymentSuccess(response)
                }, 3000)
            } else {
                // Real payment URL - open in new tab
                window.open(response.payment_url, '_blank')

                // Start polling for purchase status
                pollPurchaseStatus(response.transaction_id)
            }

            toast.show('info', `${selectedProvider.value === 'iyzico' ? 'İyzico' : 'PayPal'} ödeme sayfasına yönlendiriliyorsunuz...`)
        } else {
            // Handle API error response
            const errorMessage = response?.error || 'Bilinmeyen hata oluştu'
            toast.show('error', errorMessage)
        }
    } catch (error) {
        console.error('Purchase initiation failed:', error)
        // Check if error has message from API
        const errorMessage = error?.error || error?.message || 'Bağlantı hatası'
        toast.show('error', errorMessage)
    }
}

// Handle mock payment success for demo
function handleMockPaymentSuccess(purchaseData) {
    toast.show('success', `🎉 DEMO: ${purchaseData.tokens} token başarıyla hesabınıza eklendi!`)

    // Reset selected package
    selectedPackage.value = {}
    customTokens.value = 50

    // Emit event for parent components to refresh user data
    window.dispatchEvent(new CustomEvent('tokensUpdated', { detail: { tokens: purchaseData.tokens } }))
}

// Poll purchase status
async function pollPurchaseStatus(transactionId) {
    const maxAttempts = 120 // 10 minutes of polling
    let attempts = 0

    // Show a persistent notification about polling
    const pollingToast = toast.show('info', 'Ödeme durumunuz kontrol ediliyor...', 0) // 0 = persistent

    const checkStatus = async () => {
        try {
            const response = await useFetch(`/user/purchase-status/${transactionId}`)

            if (response.status === 'completed') {
                // Hide polling notification
                if (pollingToast && pollingToast.hide) {
                    pollingToast.hide()
                }

                toast.show('success', `🎉 ${response.tokens} token başarıyla hesabınıza eklendi!`)

                // Reset selected package
                selectedPackage.value = {}
                customTokens.value = 50

                // Optionally emit event for parent components to refresh user data
                window.dispatchEvent(new CustomEvent('tokensUpdated', { detail: { tokens: response.tokens } }))

                return true
            } else if (response.status === 'pending') {
                attempts++
                if (attempts < maxAttempts) {
                    setTimeout(checkStatus, 5000) // Check every 5 seconds
                } else {
                    // Hide polling notification
                    if (pollingToast && pollingToast.hide) {
                        pollingToast.hide()
                    }
                    toast.show('warning', 'Ödeme durumu kontrol edilemedi. Lütfen hesabınızı kontrol edin veya destek ile iletişime geçin.')
                }
            }
        } catch (error) {
            console.error('Status check failed:', error)
            attempts++
            if (attempts < maxAttempts) {
                setTimeout(checkStatus, 5000)
            } else {
                // Hide polling notification
                if (pollingToast && pollingToast.hide) {
                    pollingToast.hide()
                }
                toast.show('error', 'Ödeme durumu kontrol edilirken bir hata oluştu.')
            }
        }
    }

    // Start checking after 10 seconds (give user time to complete payment)
    setTimeout(checkStatus, 10000)
}


// Gumroad entegrasyonu için artık bu fonksiyonlar kullanılmıyor
/*
// Close checkout modal
function closeCheckoutModal() {
    showCheckoutModal.value = false
    selectedPackage.value = {}
}

// Handle successful purchase
function handlePurchaseSuccess(packageInfo) {
    toast.show('success', t('pricing.purchase.success', { tokens: packageInfo.tokens }))
    // Here you would update the user's token balance
    // updateUserTokens(packageInfo.tokens)
}
*/
</script>
