import { createRouter, createWebHistory } from "@ionic/vue-router";

import HomeView from "@/views/HomeView.vue";
import AccountView from "@/views/AccountView.vue";
import NotFoundView from "@/views/NotFoundView.vue";
import ChapterView from "@/views/chapter/index.vue";
import GameView from "@/views/GameView.vue";
import { loadLayout } from "./middleware/loadLayout";
import TestView from "@/views/game/TestView.vue";
import SearchView from "@/views/SearchView.vue";
import FollowedView from "@/views/FollowedView.vue";
import NotesView from "@/views/NotesView.vue";
import FlipWordView from "@/views/game/FlipWordView.vue";
import QuickPickView from "@/views/game/QuickPickView.vue";
import WordMatchView from "@/views/game/WordMatchView.vue";
import WordRushView from "@/views/game/WordRushView.vue";
import LoginView from "@/views/LoginView.vue";
import { authGuard } from "./middleware/authGuard";
import SettingView from "@/views/SettingView.vue";
import ComposePost from "@/views/compose/index.vue";
import EditorView from "@/views/EditorView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/compose",
      name: "compose",
      beforeEnter: [authGuard],
      component: ComposePost,
    },
    {
      path: "/update/:id",
      name: "update",
      beforeEnter: [authGuard],
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
      path: "/l/:id/game/flip-word",
      name: "flip-word",
      component: FlipWordView,
    },
    {
      path: "/l/:id/game/quick-pick",
      name: "quick-pick",
      component: QuickPickView,
    },
    {
      path: "/l/:id/game/word-match",
      name: "word-match",
      component: WordMatchView,
    },
    {
      path: "/l/:id/game/word-rush",
      name: "word-rush",
      component: WordRushView,
    },
    {
      path: "/search",
      name: "search",
      component: SearchView,
    },

    {
      path: "/followed",
      name: "followed",
      beforeEnter: [authGuard],
      component: FollowedView,
    },
    {
      path: "/notes",
      name: "notes",
      component: NotesView,
    },
    {
      path: "/login",
      name: "login",
      component: LoginView
    },
    {
      path: "/settings",
      name: "settings",
      component: SettingView
    },
    {
      path: "/editor",
      name: "editor",
      component: EditorView
    },
    {
      path: "/:username",
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
router.beforeEach(loadLayout);

export default router;
