import '../css/main.scss'

import "babel-polyfill";

import App from "./app.vue";
import {createApp} from "vue";

createApp(App).mount('#app');