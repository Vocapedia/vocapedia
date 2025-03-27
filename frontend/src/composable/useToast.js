import { useToasted } from "@hoppscotch/vue-toasted"

export const useToast = () => {
    const toast = useToasted();

    const showToast = (message, options = {}) => {
        toast.show(message, {
            ...options,
            position: 'bottom-right',
            duration: 3000,
        });
    };
    const show = (message, options = {}) => {
        showToast(message, options);
    }
    return { show }
}
