import { defineConfig } from 'vite'
import angular from '@analogjs/vite-plugin-angular';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [angular()],
  resolve: {
    mainFields: ['module'],
  },
  base: '/auth/',
  build: {
    outDir: '../../../build/auth/admin/'
  }
})
