import { ref, getCurrentInstance } from 'vue'
import { useI18n } from 'vue-i18n'

export function useTourManager() {
    const app = getCurrentInstance()
    const i18n = useI18n()
    const isTourOpen = ref(false)

    const LANGUAGE_TOUR_KEY = 'language_tour'
    const WORD_BASE_TOUR_KEY = 'word_base_tour'

    const stepsLangs = [
        {
            target: '[data-v-step="main_lang"]',
            content: i18n.t('compose.tour.main_lang'),
            params: {
                placement: 'bottom'
            }
        },
        {
            target: '[data-v-step="target_lang"]',
            content: i18n.t('compose.tour.target_lang'),
            params: {
                placement: 'bottom'
            }
        },
    ]

    const stepsWordBase = [
        {
            target: '[data-v-step="list_name"]',
            content: i18n.t('compose.tour.list_name'),
            params: {
                placement: 'top'
            }
        },
        {
            target: '[data-v-step="list_description"]',
            content: i18n.t('compose.tour.list_description'),
            params: {
                placement: 'bottom'
            }
        },
        {
            target: '[data-v-step="word_list"]',
            content: i18n.t('compose.tour.word_list'),
            params: {
                placement: 'bottom'
            }
        },
        {
            target: '[data-v-step="tutorial"]',
            content: i18n.t('compose.tour.tutorial'),
            params: {
                placement: 'bottom'
            }
        },
        {
            target: '[data-v-step="word_type"]',
            content: i18n.t('compose.tour.word_type'),
            params: {
                placement: 'bottom'
            }
        },
        {
            target: '[data-v-step="word_lang1"]',
            content: i18n.t('compose.tour.word_lang1'),
            params: {
                placement: 'top'
            }
        },
        {
            target: '[data-v-step="word_lang2"]',
            content: i18n.t('compose.tour.word_lang2'),
            params: {
                placement: 'top'
            }
        },
        {
            target: '[data-v-step="new_word_button"]',
            content: i18n.t('compose.tour.new_word_button'),
            params: {
                placement: 'left'
            }
        },
        {
            target: '[data-v-step="new_word_base"]',
            content: i18n.t('compose.tour.new_word_base'),
            params: {
                placement: 'top'
            }
        },
        {
            target: '[data-v-step="save_button"]',
            content: i18n.t('compose.tour.save_button'),
            params: {
                placement: 'bottom'
            }
        }
    ]

    const getTourKey = (tourName) => {
        if (tourName === 'language_tour') {
            return LANGUAGE_TOUR_KEY
        } else if (tourName === 'word_base_tour') {
            return WORD_BASE_TOUR_KEY
        }
        return null
    }

    const isTourCompleted = (tourName) => {
        const tourKey = getTourKey(tourName)
        return tourKey ? localStorage.getItem(tourKey) === 'true' : false
    }

    const startTour = (tourName) => {
        if (isTourCompleted(tourName)) {
            return
        }

        if (app && app.proxy && app.proxy.$tours && app.proxy.$tours[tourName]) {
            app.proxy.$tours[tourName].start()
        }
        localStorage.setItem(tourName, 'true')
    }
    const restartTours = (tourname) => {
        localStorage.removeItem(tourname)
        startTour(tourname)
    }
    return {
        isTourOpen,
        stepsLangs,
        stepsWordBase,
        startTour,
        isTourCompleted,
        restartTours
    }
}
