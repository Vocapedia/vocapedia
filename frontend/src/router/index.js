import { createRouter, createWebHistory } from "vue-router";

import HomeView from "@/views/HomeView.vue";
import AccountView from "@/views/AccountView.vue";
import NotFoundView from "@/views/NotFoundView.vue";
import ListView from "@/views/ListView.vue";
import GameView from "@/views/GameView.vue";
import { loadLayoutMiddleware } from "./middleware/loadLayoutMiddleware";
import TestView from "@/views/game/TestView.vue";

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
      path: "/l/:id",
      name: "list",
      component: ListView,
    },
    {
      path: "/l/:id/games",
      name: "games",
      component: GameView,
    },
    {
      path: "/l/:id/game/test",
      name: "test",
      component: TestView,
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
