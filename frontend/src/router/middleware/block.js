import { useToast } from "@/composable/useToast";
import { i18n } from "@/i18n/i18n";
export function block(to, from, next) {
    //useToast().show(i18n.global.t("not_ready_yet"))
    //next({ name: "home" });
    next();
}
