import { ref } from 'vue';
import { useFetch } from '@/composable/useFetch';
import { useToast } from "@/composable/useToast";
import { useRouter } from "vue-router";
import { getUser } from "@/utils/token";

export function useChapterAPI() {
    const router = useRouter();
    const toast = useToast();
    const response = ref({});

    const loadChapter = async (chapterId) => {
        try {
            response.value = await useFetch("/public/chapters/" + chapterId);
        } catch (error) {
            toast.show(error.message || "Chapter loading failed");
        }
    };

    const favoriteChapter = async (chapterId) => {
        if (!getUser()) {
            router.push("/login");
            return;
        }
        
        try {
            const result = await useFetch("/chapters/favorite?chapter_id=" + chapterId, {
                method: "POST",
            });
            
            response.value.chapter.is_favorited = 1;
            response.value.chapter.fav_count += 1;
            toast.show(result.message);
        } catch (error) {
            toast.show(error.error || "Favorite failed");
        }
    };

    const deleteFavChapter = async (chapterId) => {
        try {
            const result = await useFetch("/chapters/favorite?chapter_id=" + chapterId, {
                method: "DELETE",
            });
            
            response.value.chapter.is_favorited = 0;
            response.value.chapter.fav_count -= 1;
            toast.show(result.message);
        } catch (error) {
            toast.show(error.error || "Unfavorite failed");
        }
    };

    const archiveChapter = async (chapterId) => {
        try {
            const result = await useFetch("/chapters/archive/" + chapterId, {
                method: "DELETE",
            });
            
            router.replace("/");
            toast.show(result.message);
        } catch (error) {
            toast.show(error.error || "Archive failed");
        }
    };

    const toggleFavorite = async (chapterId, isFavorited) => {
        if (isFavorited) {
            await deleteFavChapter(chapterId);
        } else {
            await favoriteChapter(chapterId);
        }
    };

    return {
        response,
        loadChapter,
        favoriteChapter,
        deleteFavChapter,
        archiveChapter,
        toggleFavorite
    };
}
