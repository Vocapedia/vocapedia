import { createRouter, createWebHistory } from "vue-router";
import { authGuard } from "./middleware/authGuard";
import { loadLayout } from "./middleware/loadLayout";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: () => import('@/views/HomeView.vue'),
    },
    {
      path: "/landing",
      name: "landing",
      component: () => import('@/views/LandingView.vue'),
    },
    {
      path: "/compose",
      name: "compose",
      beforeEnter: [authGuard],
      component: () => import('@/views/compose/index.vue'),
    },
    {
      path: "/update/:id",
      name: "update",
      beforeEnter: [authGuard],
      component: () => import('@/views/compose/index.vue'),
    },
    {
      path: "/l/:id",
      name: "word-list",
      component: () => import('@/views/chapter/index.vue'),
    },
    {
      path: "/l/:id/games",
      name: "games",
      component: () => import('@/views/GameView.vue'),
    },
    {
      path: "/l/:id/game/test",
      name: "test",
      meta: {
        layout: "game",
      },
      component: () => import('@/views/game/TestView.vue'),
    },

    {
      path: "/l/:id/game/hangman",
      name: "hangman",
      meta: {
        layout: "game",
      },
      component: () => import('@/views/game/HangmanView.vue'),
    },
    {
      path: "/l/:id/game/flip-word",
      name: "flip-word",
      meta: {
        layout: "game",
      },
      component: () => import('@/views/game/FlipWordView.vue'),
    },
    {
      path: "/l/:id/game/quick-pick",
      name: "quick-pick",
      component: () => import('@/views/game/QuickPickView.vue'),
    },
    {
      path: "/l/:id/game/word-match",
      name: "word-match",
      component: () => import('@/views/game/WordMatchView.vue'),
    },
    {
      path: "/l/:id/game/word-rush",
      name: "word-rush",
      component: () => import('@/views/game/WordRushView.vue'),
    },
    {
      path: "/search",
      name: "search",
      component: () => import('@/views/SearchView.vue'),
    },
    {
      path: "/followed",
      name: "followed",
      beforeEnter: [authGuard],
      component: () => import('@/views/FollowedView.vue'),
    },
    {
      path: "/notes",
      name: "notes",
      component: () => import('@/views/NotesView.vue'),
    },
    {
      path: "/login",
      name: "login",
      component: () => import('@/views/LoginView.vue'),
    },
    {
      path: "/settings",
      name: "settings",
      component: () => import('@/views/SettingView.vue'),
    },
    {
      path: "/editor",
      name: "editor",
      component: () => import('@/views/EditorView.vue'),
    },
    {
      path: "/privacy-policy",
      name: "privacy-policy",
      component: () => import('@/views/PrivacyPolicy.vue'),
    },
    {
      path: "/:username",
      name: "account",
      component: () => import('@/views/AccountView.vue'),
    },
    {
      path: "/:pathMatch(.*)*",
      name: "not-found",
      component: () => import('@/views/NotFoundView.vue'),
    },
  ],
});

router.beforeEach(loadLayout);

export default router;
