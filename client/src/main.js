import { createApp } from 'vue'
import App from './App.vue'
import { Router } from './router/router.js'
import { createPinia } from 'pinia'
import './index.css'


createApp(App).use(createPinia()).use(Router).mount('#app')
