import { fileURLToPath, URL } from "node:url";

import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import vueDevTools from "vite-plugin-vue-devtools";
import tailwindcss from "@tailwindcss/vite";
import sitemap from 'vite-plugin-sitemap'; // Import the plugin

// Define static routes for the sitemap
const staticRoutes = [
  '/',
  '/landing',
  '/compose', // Assuming this is a public-facing page to create lists
  '/trends',
  '/search',
  '/notes', // Assuming notes can be public or user wants them indexed
  '/login',
  '/settings', // Publicly discoverable, though content is user-specific after login
  '/editor', // If it's a general purpose editor page
  '/privacy-policy',
  // Routes like /l/:id, /:username, /update/:id, /stream/:id are dynamic
  // and would require fetching actual IDs/params during build.
  // Auth-guarded routes like /followed, /streamers are typically excluded unless intended.
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
      dynamicRoutes: staticRoutes, // Provide the list of static routes
      generateRobotsTxt: false // Prevent sitemap plugin from modifying robots.txt
    }),
  ],
  resolve: {
    alias: {
      // "@": fileURLToPath(new URL("./src", import.meta.url)), // Original alias
      '@': '/src', // Simplified alias as per task example, assuming it works
    },
  },
  // server: { // server config from example, uncomment if needed
  //   port: 8080,
  // },
});
