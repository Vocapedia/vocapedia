import { createRouter, createWebHistory } from "vue-router";
import { authGuard } from "./middleware/authGuard";
import { loadLayout } from "./middleware/loadLayout";
import { block } from "./middleware/block";

import HomeView from '@/views/HomeView.vue';
import LandingView from '@/views/LandingView.vue';
import ComposeView from '@/views/ComposeView.vue';
import StreamView from '@/views/StreamView.vue';
import StreamersView from '@/views/StreamersView.vue';
import CreateStreamView from '@/views/CreateStreamView.vue';
import ChapterView from '@/views/ChapterView.vue';
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
import PricingView from '@/views/PricingView.vue';
import TrendsView from "@/views/TrendsView.vue";
import { i18n } from "@/i18n/i18n";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
      meta: {
        titleKey: 'seo.home.title',
        descriptionKey: 'seo.home.description',
        ogTitleKey: 'seo.home.title',
        ogDescriptionKey: 'seo.home.description',
        twitterTitleKey: 'seo.home.title',
        twitterDescriptionKey: 'seo.home.description'
      }
    },
    {
      path: "/landing",
      name: "landing",
      component: LandingView,
      meta: {
        titleKey: 'seo.landing.title',
        descriptionKey: 'seo.landing.description',
        ogTitleKey: 'seo.landing.title',
        ogDescriptionKey: 'seo.landing.description',
        twitterTitleKey: 'seo.landing.title',
        twitterDescriptionKey: 'seo.landing.description'
      }
    },
    {
      path: "/compose",
      name: "compose",
      beforeEnter: [authGuard],
      component: ComposeView,
      meta: {
        titleKey: 'seo.compose.title',
        descriptionKey: 'seo.compose.description'
      }
    },
    {
      path: "/trends",
      name: "trends",
      component: TrendsView,
      meta: {
        titleKey: 'seo.trends.title',
        descriptionKey: 'seo.trends.description'
      }
    },
    {
      path: "/stream/:id",
      name: "stream",
      beforeEnter: [block, authGuard],
      component: StreamView,
      meta: { // Assuming a generic title/desc for streams or specific keys if they exist
        titleKey: 'seo.default.title', // Placeholder, should be specific if available
        descriptionKey: 'seo.default.description' // Placeholder
      }
    },
    {
      path: "/streamers",
      name: "streamers",
      beforeEnter: [block, authGuard],
      component: StreamersView,
      meta: { // Assuming a generic title/desc
        titleKey: 'seo.default.title', // Placeholder
        descriptionKey: 'seo.default.description' // Placeholder
      }
    },
    {
      path: "/create-stream",
      name: "create-stream",
      beforeEnter: [block, authGuard],
      component: CreateStreamView,
      meta: {
        titleKey: 'seo.default.title',
        descriptionKey: 'seo.default.description'
      }
    },
    {
      path: "/update/:id",
      name: "update",
      beforeEnter: [authGuard],
      component: ComposeView,
      meta: { // Assuming it reuses compose or has its own
        titleKey: 'seo.compose.title', // Or 'seo.updateList.title' if defined
        descriptionKey: 'seo.compose.description' // Or 'seo.updateList.description'
      }
    },
    {
      path: "/l/:id",
      name: "word-list",
      component: ChapterView,
      meta: {
        titleKey: 'seo.wordList.title',
        descriptionKey: 'seo.wordList.description',
        ogTitleKey: 'seo.wordList.title',
        ogDescriptionKey: 'seo.wordList.description',
        twitterTitleKey: 'seo.wordList.title',
        twitterDescriptionKey: 'seo.wordList.description'
      }
    },
    {
      path: "/l/:id/games",
      name: "games",
      component: GameView,
      meta: {
        titleKey: 'seo.breadcrumbs.games', // Using breadcrumb key as title for now
        descriptionKey: 'seo.default.description' // Generic description
      }
    },
    {
      path: "/l/:id/game/test",
      name: "test",
      meta: {
        layout: "game",
        titleKey: 'seo.breadcrumbs.testGame', // Using breadcrumb key
        descriptionKey: 'seo.default.description'
      },
      component: TestView,
    },
    {
      path: "/l/:id/game/hangman",
      name: "hangman",
      meta: {
        layout: "game",
        titleKey: 'seo.breadcrumbs.hangmanGame', // Using breadcrumb key
        descriptionKey: 'seo.default.description'
      },
      component: HangmanView,
    },
    {
      path: "/l/:id/game/flip-word",
      name: "flip-word",
      meta: {
        layout: "game",
        titleKey: 'seo.breadcrumbs.flipWordGame', // Using breadcrumb key
        descriptionKey: 'seo.default.description'
      },
      component: FlipWordView,
    },
    {
      path: "/l/:id/game/quick-pick",
      name: "quick-pick",
      component: QuickPickView,
      meta: {
        titleKey: 'seo.breadcrumbs.quickPickGame', // Using breadcrumb key
        descriptionKey: 'seo.default.description'
      }
    },
    {
      path: "/l/:id/game/word-match",
      name: "word-match",
      component: WordMatchView,
      meta: {
        titleKey: 'seo.breadcrumbs.wordMatchGame', // Using breadcrumb key
        descriptionKey: 'seo.default.description'
      }
    },
    {
      path: "/l/:id/game/word-rush",
      name: "word-rush",
      component: WordRushView,
      meta: {
        titleKey: 'seo.breadcrumbs.wordRushGame', // Using breadcrumb key
        descriptionKey: 'seo.default.description'
      }
    },
    {
      path: "/search",
      name: "search",
      component: SearchView,
      meta: {
        titleKey: 'seo.search.title',
        descriptionKey: 'seo.search.description'
      }
    },
    {
      path: "/followed",
      name: "followed",
      beforeEnter: [authGuard],
      component: FollowedView,
      meta: { // Assuming generic or specific keys if available for followed lists
        titleKey: 'seo.default.title', // Placeholder
        descriptionKey: 'seo.default.description' // Placeholder
      }
    },
    {
      path: "/notes",
      name: "notes",
      component: NotesView,
      meta: {
        titleKey: 'seo.notes.title',
        descriptionKey: 'seo.notes.description'
      }
    },
    {
      path: "/login",
      name: "login",
      component: LoginView,
      meta: {
        titleKey: 'seo.login.title',
        descriptionKey: 'seo.login.description'
      }
    },
    {
      path: "/settings",
      name: "settings",
      component: SettingView,
      meta: {
        titleKey: 'seo.settings.title',
        descriptionKey: 'seo.settings.description'
      }
    },
    {
      path: "/editor",
      name: "editor",
      component: EditorView,
      meta: {
        titleKey: 'seo.editor.title',
        descriptionKey: 'seo.editor.description'
      }
    },
    {
      path: "/privacy-policy",
      name: "privacy-policy",
      component: PrivacyPolicy,
      meta: {
        titleKey: 'seo.privacyPolicy.title',
        descriptionKey: 'seo.privacyPolicy.description'
      }
    },
    {
      path: "/:username",
      name: "account",
      component: AccountView,
      meta: {
        titleKey: 'seo.userProfile.title',
        descriptionKey: 'seo.userProfile.description'
      }
    },
    {
      path: "/pricing",
      name: "pricing",
      component: PricingView,
      meta: {
        titleKey: 'seo.pricing.title',
        descriptionKey: 'seo.pricing.description',
        ogTitleKey: 'seo.pricing.title',
        ogDescriptionKey: 'seo.pricing.description',
        twitterTitleKey: 'seo.pricing.title',
        twitterDescriptionKey: 'seo.pricing.description'
      }
    },
    {
      path: "/:pathMatch(.*)*",
      name: "not-found",
      component: NotFoundView,
      meta: {
        titleKey: 'seo.notFound.title',
        descriptionKey: 'seo.notFound.description'
      }
    },
  ],
});

router.beforeEach(loadLayout);

router.afterEach((to, from) => {
  const siteUrl = 'https://vocapedia.space'; // Base URL

  // Update document title
  document.title = i18n.global.t(to.meta.titleKey || 'seo.default.title');

  // Update meta description
  const descriptionTag = document.querySelector('meta[name="description"]');
  if (descriptionTag) {
    descriptionTag.setAttribute('content', i18n.global.t(to.meta.descriptionKey || 'seo.default.description'));
  }

  // Update Open Graph tags
  const ogTitleTag = document.querySelector('meta[property="og:title"]');
  if (ogTitleTag) {
    ogTitleTag.setAttribute('content', i18n.global.t(to.meta.ogTitleKey || to.meta.titleKey || 'seo.default.title'));
  }
  const ogDescriptionTag = document.querySelector('meta[property="og:description"]');
  if (ogDescriptionTag) {
    ogDescriptionTag.setAttribute('content', i18n.global.t(to.meta.ogDescriptionKey || to.meta.descriptionKey || 'seo.default.description'));
  }
  const ogUrlTag = document.querySelector('meta[property="og:url"]');
  if (ogUrlTag) {
    ogUrlTag.setAttribute('content', siteUrl + to.fullPath); // Use to.fullPath for accuracy
  }

  // Update Twitter card tags
  const twitterTitleTag = document.querySelector('meta[name="twitter:title"]');
  if (twitterTitleTag) {
    twitterTitleTag.setAttribute('content', i18n.global.t(to.meta.twitterTitleKey || to.meta.titleKey || 'seo.default.title'));
  }
  const twitterDescriptionTag = document.querySelector('meta[name="twitter:description"]');
  if (twitterDescriptionTag) {
    twitterDescriptionTag.setAttribute('content', i18n.global.t(to.meta.twitterDescriptionKey || to.meta.descriptionKey || 'seo.default.description'));
  }

  const canonicalLink = document.querySelector('link[rel="canonical"]');
  if (canonicalLink) {
    canonicalLink.setAttribute('href', siteUrl + to.fullPath); // Use to.fullPath for accuracy with queries/hashes
  }

  // Breadcrumb JSON-LD update
  const breadcrumbScript = document.getElementById('breadcrumb-json-ld');
  if (breadcrumbScript) {
    const breadcrumbs = [
      {
        "@type": "ListItem",
        "position": 1,
        "name": i18n.global.t('seo.breadcrumbs.home'), // Localized "Home"
        "item": siteUrl + "/"
      }
    ];

    // Accumulator for the path of parent routes
    // let accumulatedPath = ''; // Not strictly needed if routeRecord.path is always absolute

    to.matched.forEach((routeRecord, index) => {
      // Skip the root route if it's the first in matched, as "Home" is already added.
      // This check is important if the root route (path: '/') has a name and would otherwise be processed.
      if (index === 0 && routeRecord.path === '/') {
        // If the current page IS the home page, its name should be the 'Home' breadcrumb name.
        // The loop for breadcrumbs should effectively generate only the "Home" item.
        // The current logic might add a second "Home" if not handled.
        // For now, we let the last item logic handle the name of the current page.
        // If to.matched.length is 1 and it's home, breadcrumbs will have one item.
        // If to.matched.length > 1, the first item is Home, then children.
      }

      let name = '';
      let itemPath = ''; // This will be the path for the 'item' URL for intermediate items

      // Determine breadcrumb name key
      let breadcrumbNameKey = '';
      if (routeRecord.meta && routeRecord.meta.breadcrumbKey) {
        breadcrumbNameKey = routeRecord.meta.breadcrumbKey;
      } else if (routeRecord.name) {
        const routeNameStr = String(routeRecord.name);
        // Convert camelCase or PascalCase route names to simple lowercase for keys like 'wordList' or 'privacyPolicy'
        // This assumes breadcrumb keys in JSON are like seo.breadcrumbs.wordList, seo.breadcrumbs.privacyPolicy
        const keyName = routeNameStr.replace(/([A-Z])/g, '-$1').toLowerCase().replace(/^-/, ''); // e.g., 'word-list', 'privacy-policy'
                                                                                                // or simply use as is if keys match:
                                                                                                // const keyName = routeNameStr.charAt(0).toLowerCase() + routeNameStr.slice(1);
        breadcrumbNameKey = `seo.breadcrumbs.${keyName}`;
        // Fallback for names that might not match simple lowercase conversion if JSON keys are specific
        // e.g. if JSON has seo.breadcrumbs.testGame and route name is 'test'
        // For game routes, their titleKey was already like 'seo.breadcrumbs.testGame'
        if(routeRecord.meta?.titleKey?.startsWith('seo.breadcrumbs.')) {
             breadcrumbNameKey = routeRecord.meta.titleKey;
        }
      }

      // Path for the item URL. routeRecord.path is the path pattern defined in routes.
      itemPath = routeRecord.path;
      // Substitute params for the current route context.
      for (const paramKey in to.params) {
        if (itemPath.includes(`:${paramKey}`)) {
          itemPath = itemPath.replace(`:${paramKey}`, to.params[paramKey]);
        }
      }
      // Ensure path starts with a slash if it's not the root path itself
      if (itemPath !== '/' && !itemPath.startsWith('/')) {
        itemPath = '/' + itemPath;
      }


      // Determine the name for the breadcrumb
      if (index === to.matched.length - 1) { // Current page
        name = i18n.global.t(to.meta.titleKey || 'seo.default.title');
        itemPath = to.fullPath; // Current page's item URL is its full path
      } else if (breadcrumbNameKey) { // Intermediate breadcrumb
        // Use route name as fallback if key is not found by i18n.t
        const fallbackName = routeRecord.name ? (String(routeRecord.name).charAt(0).toUpperCase() + String(routeRecord.name).slice(1)) : 'Unnamed';
        name = i18n.global.t(breadcrumbNameKey, fallbackName);
         // If translation is same as key (i.e., not found), use fallback
        if (name === breadcrumbNameKey) name = fallbackName;
      } else { // Fallback if no key and no name
        name = routeRecord.name ? (String(routeRecord.name).charAt(0).toUpperCase() + String(routeRecord.name).slice(1)) : 'Page';
      }

      // Add breadcrumb item, but avoid duplicating "Home" if it's the first matched route.
      // The first item in `breadcrumbs` is already "Home".
      // So, if routeRecord.path is '/', we only add it if it's the *only* item (i.e., current page is home).
      // Otherwise, we've already added "Home" and this iteration for '/' should be skipped.
      if (routeRecord.path === '/' && breadcrumbs.length === 1 && to.matched.length > 1) {
        // This is the root path '/' being processed in a sequence of matched routes,
        // and "Home" is already the first breadcrumb. Skip adding another.
      } else if (name) { // Only add if a name was resolved
         // If this is the first actual page in matched (e.g. '/home' or '/dashboard')
         // and it's not the root path being handled as the current page.
        if (breadcrumbs.length === 1 && routeRecord.path === '/' && index === to.matched.length -1){
            // Current page is home, update the name of the first breadcrumb.
            breadcrumbs[0].name = name;
        } else {
            breadcrumbs.push({
              "@type": "ListItem",
              "position": breadcrumbs.length + 1,
              "name": name,
              "item": siteUrl + itemPath
            });
        }
      }
    });

    // If after processing, the breadcrumbs list only has one item and it's not the current page (e.g. Home linking to itself),
    // or if the current page is home and its name got updated.
    // Ensure the last item has its URL as the full current path if it wasn't set that way.
    if (breadcrumbs.length > 0 && breadcrumbs[breadcrumbs.length -1].item !== siteUrl + to.fullPath) {
        // This case should be rare if the logic above is correct for last item.
        // breadcrumbs[breadcrumbs.length -1].item = siteUrl + to.fullPath;
    }


    const breadcrumbLdJson = {
      "@context": "https://schema.org",
      "@type": "BreadcrumbList",
      "itemListElement": breadcrumbs
    };
    breadcrumbScript.textContent = JSON.stringify(breadcrumbLdJson, null, 2);
  }

  // Manage hreflang tags
  const supportedLangs = ['en', 'tr', 'de', 'es', 'fr', 'zh'];
  // const defaultLang = 'en'; // Not explicitly used in the provided snippet for href construction
  const siteBaseUrl = siteUrl; // siteUrl is already defined as https://vocapedia.space

  // Remove previously added hreflang tags
  document.querySelectorAll('link[data-hreflang-tag="true"]').forEach(tag => tag.remove());

  const head = document.head;
  const pagePath = to.fullPath; // e.g., /search?query=term or /l/some-id

  // Add new hreflang tags for each supported language
  supportedLangs.forEach(langCode => {
    const link = document.createElement('link');
    link.setAttribute('rel', 'alternate');
    link.setAttribute('hreflang', langCode);
    link.setAttribute('href', siteBaseUrl + pagePath);
    link.setAttribute('data-hreflang-tag', 'true');
    head.appendChild(link);
  });

  // Add x-default hreflang tag
  const xDefaultLink = document.createElement('link');
  xDefaultLink.setAttribute('rel', 'alternate');
  xDefaultLink.setAttribute('hreflang', 'x-default');
  xDefaultLink.setAttribute('href', siteBaseUrl + pagePath);
  xDefaultLink.setAttribute('data-hreflang-tag', 'true');
  head.appendChild(xDefaultLink);

});

export default router;
