import { createRouter, createWebHistory } from "vue-router";

import HomeView from "@/views/HomeView.vue";
import AccountView from "@/views/AccountView.vue";
import NotFoundView from "@/views/NotFoundView.vue";
import { loadLayoutMiddleware } from "./middleware/loadLayoutMiddleware";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/account",
      name: "account",
      component: AccountView,
    },
    {
      path: "/:pathMatch(.*)*",
      name: "not-found",
      component: NotFoundView,
    },
  ],
});
router.beforeEach(loadLayoutMiddleware);

export default router;
