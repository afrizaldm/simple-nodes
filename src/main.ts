import './bootstrap';

import { createApp } from 'vue'
import router from "./router/index";
import App from './pages/App.vue'


const app = createApp(App)
app.use(router)

app.mount('#app')

// createApp(App).mount('#app')
