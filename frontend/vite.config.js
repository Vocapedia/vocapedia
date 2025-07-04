import { fileURLToPath, URL } from "node:url";

import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import vueDevTools from "vite-plugin-vue-devtools";
import tailwindcss from "@tailwindcss/vite";
import sitemap from 'vite-plugin-sitemap'; // Import the plugin

const now = new Date().toISOString();
const staticRoutes = [
  // Main pages
  { url: '/', changefreq: 'daily', priority: 1.0 },
  { url: '/landing', changefreq: 'weekly', priority: 0.9 },
  { url: '/compose', changefreq: 'weekly', priority: 0.8 },
  { url: '/trends', changefreq: 'daily', priority: 0.9 },
  { url: '/search', changefreq: 'daily', priority: 0.9 },
  
  // User features
  { url: '/notes', changefreq: 'weekly', priority: 0.7 },
  { url: '/followed', changefreq: 'weekly', priority: 0.6 },
  { url: '/editor', changefreq: 'monthly', priority: 0.6 },
  
  // Authentication
  { url: '/login', changefreq: 'monthly', priority: 0.3 },
  { url: '/settings', changefreq: 'monthly', priority: 0.3 },
  
  // Legal and info
  { url: '/privacy-policy', changefreq: 'yearly', priority: 0.2 },
].map(route => ({
  ...route,
  lastmod: now
}));


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
      generateRobotsTxt: true, // Generate robots.txt with sitemap reference
      robots: [
        {
          userAgent: '*',
          allow: '/',
          disallow: ['/admin/', '/api/', '/private/'],
          sitemap: 'https://vocapedia.space/sitemap.xml'
        }
      ],
      // Additional sitemap options for better SEO
      exclude: ['/admin', '/api', '/private'],
      readable: true, // Generate human-readable sitemap
      trailingSlash: false
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
