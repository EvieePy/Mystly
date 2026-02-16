import { sveltekit } from '@sveltejs/kit/vite';
import extractorSvelte from '@unocss/extractor-svelte'
import UnoCSS from 'unocss/vite'
import { defineConfig } from 'vite';


export default defineConfig({
	plugins: [
    UnoCSS({
      extractors: [
        extractorSvelte(),
      ],
    }),
    sveltekit(),
  ],
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:7171',
        changeOrigin: true
      }
    }
  }
});
