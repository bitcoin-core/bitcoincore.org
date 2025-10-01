// @ts-check
import { defineConfig } from 'astro/config';
import mdx from '@astrojs/mdx';
import sitemap from '@astrojs/sitemap';

import tailwindcss from '@tailwindcss/vite';

// https://astro.build/config
export default defineConfig({
  site: 'https://bitcoincore.yachts',
  integrations: [mdx(), sitemap()],

  i18n: {
    defaultLocale: 'en',
    locales: ['en', 'es', 'ja', 'zh-cn', 'zh-tw'],
    routing: {
      prefixDefaultLocale: false
    }
  },

  markdown: {
    shikiConfig: {
      theme: 'github-dark'
    }
  },

  vite: {
    plugins: [tailwindcss()]
  }
});