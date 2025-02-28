import { createRouter, createWebHistory } from "@ionic/vue-router";

import HomeView from "@/views/HomeView.vue";
import AccountView from "@/views/AccountView.vue";
import NotFoundView from "@/views/NotFoundView.vue";
import ChapterView from "@/views/chapter/index.vue";
import GameView from "@/views/GameView.vue";
import { loadLayoutMiddleware } from "./middleware/loadLayoutMiddleware";
import TestView from "@/views/game/TestView.vue";
import SearchView from "@/views/SearchView.vue";
import ComposePost from "@/views/ComposePost.vue";
import FollowedView from "@/views/FollowedView.vue";
import NotesView from "@/views/NotesView.vue";

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
      path: "/compose",
      name: "compose",
      component: ComposePost,
    },
    {
      path: "/l/:id",
      name: "word-list",
      component: ChapterView,
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
      path: "/search",
      name: "search",
      component: SearchView,
    },

    {
      path: "/followed",
      name: "followed",
      component: FollowedView,
    },
    {
      path: "/notes",
      name: "notes",
      component: NotesView,
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
