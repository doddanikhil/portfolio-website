import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    port: 3000, // Development server port
  },
  build: {
    outDir: 'build', // Output directory for production builds
  },
});
