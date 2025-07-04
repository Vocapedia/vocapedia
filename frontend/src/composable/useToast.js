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
    
    const success = (message, options = {}) => {
        showToast(message, {
            ...options,
            type: 'success',
            className: 'bg-green-500 text-white',
        });
    }
    
    const error = (message, options = {}) => {
        showToast(message, {
            ...options,
            type: 'error',
            className: 'bg-red-500 text-white',
        });
    }
    
    return { show, success, error }
}
