import {createApp} from 'vue'
import {createPinia} from 'pinia'
import ArcoVue from '@arco-design/web-vue';
import App from './App.vue'
import router from './router'
import '@arco-design/web-vue/dist/arco.css';
import "@/assets/theme.css"
import "@/assets/base.css"
import "font-awesome/css/font-awesome.min.css"
const app = createApp(App)
app.use(ArcoVue)
app.use(createPinia())
app.use(router);
import "nprogress/nprogress.css";

(window as any).isLaptops = document.documentElement.clientWidth > 1400 && document.documentElement.clientWidth < 1800


app.mount('#app')
