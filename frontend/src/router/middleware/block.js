import { useToast } from "@/composable/useToast";
import { i18n } from "@/i18n/i18n";
const mode = import.meta.env.VITE_MODE;
const built = import.meta.env.VITE_BUILT;

export function block(to, from, next) {
    if (mode === "development" || built == 1) {
        next();
    }
    else {
        useToast().show(i18n.global.t("not_ready_yet"))
        next({ name: "home" });
    }
}
