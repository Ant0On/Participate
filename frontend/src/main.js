import {createApp} from 'vue'
import { createPinia } from 'pinia'

import { router } from '@/router'
import App from "./App.vue";

import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import '@mdi/font/css/materialdesignicons.css'

function clearLocalStorage() {
    localStorage.clear();
}

function handleBeforeUnload(event) {
    if (navigator.sendBeacon) {
        navigator.sendBeacon('/clear-local-storage');
    } else {
        clearLocalStorage();
    }
}

window.addEventListener('beforeunload', handleBeforeUnload);

const vuetify = createVuetify({
    components,
    directives,
})

const app = createApp(App)
const pinia = createPinia()

app.use(vuetify)
app.use(pinia)
app.use(router)

app.mount('#app')
