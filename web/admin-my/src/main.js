import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'
import VueAxios from 'vue-axios'
import {antArr} from "@/assets/AntDesign";
import "@/assets/css/style.css"

axios.defaults.baseURL = 'http://localhost:8080/api/v1'
const app = createApp(App)
app.use(store).use(router)
app.use(VueAxios, axios)
app.provide('axios', app.config.globalProperties.axios)
app.config.globalProperties.$http = axios
// 自动使用antd ui
for (const item of antArr) {
    console.log(item)
    app.use(item);
}
app.mount('#app')
