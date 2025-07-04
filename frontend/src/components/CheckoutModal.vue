<template>
    <div v-if="isVisible" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50" @click="closeModal">
        <div class="bg-white dark:bg-gray-800 rounded-lg p-6 max-w-md w-full mx-4 relative overflow-hidden" @click.stop>
            <!-- Close button -->
            <button @click="closeModal" class="absolute top-4 right-4 text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200">
                <mdicon name="close" size="24" />
            </button>
            
            <!-- Header -->
            <div class="text-center mb-6">
                <div class="w-16 h-16 bg-gradient-to-r from-purple-600 to-pink-600 rounded-full flex items-center justify-center mx-auto mb-4">
                    <mdicon name="credit-card" size="24" class="text-white" />
                </div>
                <h2 class="text-2xl font-bold text-gray-900 dark:text-white">
                    {{ $t('checkout.title') }}
                </h2>
                <p class="text-gray-600 dark:text-gray-300 mt-2">
                    {{ $t('checkout.subtitle') }}
                </p>
            </div>

            <!-- Progress Steps -->
            <div class="flex justify-center mb-6">
                <div class="flex space-x-2">
                    <div v-for="step in 3" :key="step" 
                         :class="[
                             'w-3 h-3 rounded-full transition-all duration-300',
                             step <= currentStep ? 'bg-purple-600' : 'bg-gray-300 dark:bg-gray-600'
                         ]">
                    </div>
                </div>
            </div>

            <!-- Step 1: Package Info -->
            <div v-if="currentStep === 1" class="animate-fade-in">
                <div class="bg-gray-50 dark:bg-gray-700 rounded-lg p-4 mb-6">
                    <h3 class="font-semibold text-gray-900 dark:text-white mb-2">
                        {{ $t('checkout.package_details') }}
                    </h3>
                    <div class="space-y-2">
                        <div class="flex justify-between">
                            <span class="text-gray-600 dark:text-gray-300">{{ $t('checkout.tokens') }}:</span>
                            <span class="font-semibold text-gray-900 dark:text-white">{{ packageInfo.tokens }}</span>
                        </div>
                        <div class="flex justify-between">
                            <span class="text-gray-600 dark:text-gray-300">{{ $t('checkout.price') }}:</span>
                            <span class="font-semibold text-gray-900 dark:text-white">${{ packageInfo.price }}</span>
                        </div>
                        <div v-if="packageInfo.discount > 0" class="flex justify-between text-green-600 dark:text-green-400">
                            <span>{{ $t('checkout.discount') }}:</span>
                            <span>-{{ packageInfo.discount }}%</span>
                        </div>
                    </div>
                </div>
                <button @click="nextStep" class="w-full bg-purple-600 hover:bg-purple-700 text-white font-semibold py-3 px-6 rounded-lg transition duration-200">
                    {{ $t('checkout.continue') }}
                </button>
            </div>

            <!-- Step 2: Payment Details -->
            <div v-if="currentStep === 2" class="animate-fade-in">
                <form @submit.prevent="nextStep" class="space-y-4">
                    <!-- Card Number -->
                    <div>
                        <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                            {{ $t('checkout.card_number') }}
                        </label>
                        <div class="relative">
                            <input 
                                v-model="paymentData.cardNumber" 
                                type="text" 
                                placeholder="1234 5678 9012 3456"
                                maxlength="19"
                                @input="formatCardNumber"
                                class="w-full px-4 py-3 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-transparent dark:bg-gray-700 dark:text-white pr-12"
                                required
                            />
                            <div class="absolute right-3 top-3">
                                <mdicon name="credit-card-outline" size="20" class="text-gray-400" />
                            </div>
                        </div>
                    </div>

                    <!-- Expiry and CVV -->
                    <div class="grid grid-cols-2 gap-4">
                        <div>
                            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                                {{ $t('checkout.expiry_date') }}
                            </label>
                            <input 
                                v-model="paymentData.expiryDate" 
                                type="text" 
                                placeholder="MM/YY"
                                maxlength="5"
                                @input="formatExpiryDate"
                                class="w-full px-4 py-3 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-transparent dark:bg-gray-700 dark:text-white"
                                required
                            />
                        </div>
                        <div>
                            <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                                {{ $t('checkout.cvv') }}
                            </label>
                            <input 
                                v-model="paymentData.cvv" 
                                type="text" 
                                placeholder="123"
                                maxlength="4"
                                class="w-full px-4 py-3 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-transparent dark:bg-gray-700 dark:text-white"
                                required
                            />
                        </div>
                    </div>

                    <!-- Cardholder Name -->
                    <div>
                        <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                            {{ $t('checkout.cardholder_name') }}
                        </label>
                        <input 
                            v-model="paymentData.cardholderName" 
                            type="text" 
                            placeholder="John Doe"
                            class="w-full px-4 py-3 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-transparent dark:bg-gray-700 dark:text-white"
                            required
                        />
                    </div>

                    <!-- Email -->
                    <div>
                        <label class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2">
                            {{ $t('checkout.email') }}
                        </label>
                        <input 
                            v-model="paymentData.email" 
                            type="email" 
                            placeholder="john@example.com"
                            class="w-full px-4 py-3 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-transparent dark:bg-gray-700 dark:text-white"
                            required
                        />
                    </div>

                    <div class="flex space-x-4">
                        <button @click="prevStep" type="button" class="flex-1 bg-gray-600 hover:bg-gray-700 text-white font-semibold py-3 px-6 rounded-lg transition duration-200">
                            {{ $t('checkout.back') }}
                        </button>
                        <button type="submit" class="flex-1 bg-purple-600 hover:bg-purple-700 text-white font-semibold py-3 px-6 rounded-lg transition duration-200">
                            {{ $t('checkout.process_payment') }}
                        </button>
                    </div>
                </form>
            </div>

            <!-- Step 3: Processing/Success -->
            <div v-if="currentStep === 3" class="animate-fade-in text-center">
                <div v-if="isProcessing" class="space-y-4">
                    <div class="w-16 h-16 mx-auto">
                        <div class="animate-spin rounded-full h-16 w-16 border-4 border-purple-600 border-t-transparent"></div>
                    </div>
                    <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
                        {{ $t('checkout.processing') }}
                    </h3>
                    <p class="text-gray-600 dark:text-gray-300">
                        {{ $t('checkout.processing_message') }}
                    </p>
                </div>
                
                <div v-else class="space-y-4">
                    <div class="w-16 h-16 bg-green-100 dark:bg-green-900 rounded-full flex items-center justify-center mx-auto">
                        <mdicon name="check" size="24" class="text-green-600 dark:text-green-400" />
                    </div>
                    <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
                        {{ $t('checkout.success') }}
                    </h3>
                    <p class="text-gray-600 dark:text-gray-300">
                        {{ $t('checkout.success_message') }}
                    </p>
                    <div class="bg-green-50 dark:bg-green-900/20 rounded-lg p-4">
                        <p class="text-green-800 dark:text-green-400 font-semibold">
                            +{{ packageInfo.tokens }} {{ $t('checkout.tokens_added') }}
                        </p>
                    </div>
                    <button @click="closeModal" class="w-full bg-purple-600 hover:bg-purple-700 text-white font-semibold py-3 px-6 rounded-lg transition duration-200">
                        {{ $t('checkout.close') }}
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, reactive, watch } from 'vue'
import { useI18n } from 'vue-i18n'
import { useToast } from '@/composable/useToast'

const { t } = useI18n()
const toast = useToast()

// Props
const props = defineProps({
    isVisible: {
        type: Boolean,
        default: false
    },
    packageInfo: {
        type: Object,
        default: () => ({
            tokens: 0,
            price: '0.00',
            discount: 0
        })
    }
})

// Emits
const emit = defineEmits(['close', 'success'])

// State
const currentStep = ref(1)
const isProcessing = ref(false)
const paymentData = reactive({
    cardNumber: '',
    expiryDate: '',
    cvv: '',
    cardholderName: '',
    email: ''
})

// Methods
function closeModal() {
    emit('close')
    resetForm()
}

function resetForm() {
    currentStep.value = 1
    isProcessing.value = false
    Object.assign(paymentData, {
        cardNumber: '',
        expiryDate: '',
        cvv: '',
        cardholderName: '',
        email: ''
    })
}

function nextStep() {
    if (currentStep.value < 3) {
        currentStep.value++
        if (currentStep.value === 3) {
            processPayment()
        }
    }
}

function prevStep() {
    if (currentStep.value > 1) {
        currentStep.value--
    }
}

function formatCardNumber() {
    let value = paymentData.cardNumber.replace(/\s+/g, '').replace(/[^0-9]/gi, '')
    let formattedValue = value.match(/.{1,4}/g)?.join(' ') || ''
    paymentData.cardNumber = formattedValue
}

function formatExpiryDate() {
    let value = paymentData.expiryDate.replace(/\D/g, '')
    if (value.length >= 2) {
        value = value.substring(0, 2) + '/' + value.substring(2, 4)
    }
    paymentData.expiryDate = value
}

async function processPayment() {
    isProcessing.value = true
    
    // Simulate payment processing
    await new Promise(resolve => setTimeout(resolve, 3000))
    
    isProcessing.value = false
    
    // Simulate successful payment
    toast.show('success', t('checkout.payment_successful'))
    emit('success', props.packageInfo)
    
    // Auto close after 2 seconds
    setTimeout(() => {
        closeModal()
    }, 2000)
}

// Watch for modal visibility changes
watch(() => props.isVisible, (newVal) => {
    if (!newVal) {
        resetForm()
    }
})
</script>

<style scoped>
.animate-fade-in {
    animation: fadeIn 0.3s ease-in-out;
}

@keyframes fadeIn {
    from {
        opacity: 0;
        transform: translateY(10px);
    }
    to {
        opacity: 1;
        transform: translateY(0);
    }
}
</style>
