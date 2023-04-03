import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import { sveltePreprocess } from 'svelte-preprocess/dist/autoProcess'
import rust from '@wasm-tool/rollup-plugin-rust';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    rust({ 
      verbose: true,
      serverPath: "./src-rust/build/"
    }),
    svelte({
        preprocess: sveltePreprocess()
    })
  ],
})
