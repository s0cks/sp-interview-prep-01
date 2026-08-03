import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      '/sensors': {
        target: 'http://localhost:8080', // Your backend container URL
        changeOrigin: true,
        secure: false,
      }
    }
  }
})
