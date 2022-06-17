import { createApp } from 'vue'
import App from './App.vue'
import router from './router.js'
import store from './store'
import { createVuestic } from 'vuestic-ui'
import 'vuestic-ui/dist/vuestic-ui.css'
const app = createApp(App)
app
   .use(router)
   .use(store)
   .use(createVuestic())
app.mount('#app')