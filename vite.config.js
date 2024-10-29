import { defineConfig } from 'vite';

import vue from '@vitejs/plugin-vue'
import path from 'path'

export default defineConfig({
    base: "public",
    plugins: [
        vue(),
    ],
    build: {
        rollupOptions: {
            input: {
                main: '/src/main.ts',
                index: '/src/app.html',
            },
        },
        outDir: 'public',
        emptyOutDir: true,
    },
    resolve: {
        alias: {
            '@': path.resolve(__dirname, 'src'),
            '~': path.resolve(__dirname, 'js'),
            '$': path.resolve(__dirname, ''),
        },
    },
});