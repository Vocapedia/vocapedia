import { createRouter, createWebHistory } from "vue-router";
import { authGuard } from "./middleware/authGuard";
import { loadLayout } from "./middleware/loadLayout";
import { block } from "./middleware/block";

import HomeView from '@/views/HomeView.vue';
import LandingView from '@/views/LandingView.vue';
import ComposeView from '@/views/compose/index.vue';
import StreamView from '@/views/StreamView.vue';
import StreamersView from '@/views/StreamersView.vue';
import ChapterView from '@/views/chapter/index.vue';
import GameView from '@/views/GameView.vue';
import TestView from '@/views/game/TestView.vue';
import HangmanView from '@/views/game/HangmanView.vue';
import FlipWordView from '@/views/game/FlipWordView.vue';
import QuickPickView from '@/views/game/QuickPickView.vue';
import WordMatchView from '@/views/game/WordMatchView.vue';
import WordRushView from '@/views/game/WordRushView.vue';
import SearchView from '@/views/SearchView.vue';
import FollowedView from '@/views/FollowedView.vue';
import NotesView from '@/views/NotesView.vue';
import LoginView from '@/views/LoginView.vue';
import SettingView from '@/views/SettingView.vue';
import EditorView from '@/views/EditorView.vue';
import PrivacyPolicy from '@/views/PrivacyPolicy.vue';
import AccountView from '@/views/AccountView.vue';
import NotFoundView from '@/views/NotFoundView.vue';
import TrendsView from "@/views/TrendsView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/landing",
      name: "landing",
      component: LandingView,
    },
    {
      path: "/compose",
      name: "compose",
      beforeEnter: [authGuard],
      component: ComposeView,
    },
    {
      path: "/trends",
      name: "trends",
      component: TrendsView,
    },
    {
      path: "/stream/:id",
      name: "stream",
      beforeEnter: [block, authGuard],
      component: StreamView,
    },
    {
      path: "/streamers",
      name: "streamers",
      beforeEnter: [block, authGuard],
      component: StreamersView,
    },
    {
      path: "/update/:id",
      name: "update",
      beforeEnter: [authGuard],
      component: ComposeView,
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
      meta: {
        layout: "game",
      },
      component: TestView,
    },
    {
      path: "/l/:id/game/hangman",
      name: "hangman",
      meta: {
        layout: "game",
      },
      component: HangmanView,
    },
    {
      path: "/l/:id/game/flip-word",
      name: "flip-word",
      meta: {
        layout: "game",
      },
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
      component: LoginView,
    },
    {
      path: "/settings",
      name: "settings",
      component: SettingView,
    },
    {
      path: "/editor",
      name: "editor",
      component: EditorView,
    },
    {
      path: "/privacy-policy",
      name: "privacy-policy",
      component: PrivacyPolicy,
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
