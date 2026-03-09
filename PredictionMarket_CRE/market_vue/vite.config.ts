import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
    server: {
        host: true, // 或者 '0.0.0.0'，这样允许局域网和本机所有地址访问
        port: 5173, // 你当前的端口
    }
})

