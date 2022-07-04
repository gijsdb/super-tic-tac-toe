import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router';
import App from './App.vue'
import { routes } from './router/routes.js'
import { createPinia } from 'pinia'
import './index.css'

const router = createRouter({
    // 4. Provide the history implementation to use. We are using the hash history for simplicity here.
    history: createWebHistory(),
    routes
})

createApp(App).use(createPinia()).use(router).mount('#app')
