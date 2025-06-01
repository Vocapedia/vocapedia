import { fileURLToPath, URL } from "node:url";

import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import vueDevTools from "vite-plugin-vue-devtools";
import tailwindcss from "@tailwindcss/vite";
import sitemap from 'vite-plugin-sitemap'; // Import the plugin

// Define static routes for the sitemap
const staticRoutes = [
  { url: '/', changefreq: 'daily', priority: 1.0, lastmod: new Date().toISOString() },
  { url: '/landing', changefreq: 'weekly', priority: 0.9 },
  { url: '/compose', changefreq: 'weekly', priority: 0.8 },
  { url: '/trends', changefreq: 'daily', priority: 0.9 },
  { url: '/search', changefreq: 'daily', priority: 0.9 },
  { url: '/notes', changefreq: 'weekly', priority: 0.6 },
  { url: '/login', changefreq: 'monthly', priority: 0.3 },
  { url: '/settings', changefreq: 'monthly', priority: 0.3 },
  { url: '/editor', changefreq: 'monthly', priority: 0.5 },
  { url: '/privacy-policy', changefreq: 'yearly', priority: 0.2 },
];

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    vueDevTools(),
    tailwindcss(),
    sitemap({
      hostname: 'https://vocapedia.space',
      outDir: 'dist', // Explicitly set output directory for sitemap.xml
      customPages: staticRoutes,
      generateRobotsTxt: false // Prevent sitemap plugin from modifying robots.txt
    }),
  ],
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)), // Original alias
      //'@': '/src', // Simplified alias as per task example, assuming it works
    },
  },
  // server: { // server config from example, uncomment if needed
  //   port: 8080,
  // },
});
