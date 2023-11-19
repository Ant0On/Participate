import { createRouter, createWebHistory } from 'vue-router'

import Login from '@/views/Login.vue'
import Home from '@/views/Home.vue'
import Accommodations from "@/views/Accommodations.vue";
import Events from "@/views/Events.vue";
import Activities from "@/views/Activities.vue";

const routes = [
    {
        path: '/',
        name: "Home",
        component: Home
    },
    {
        path: '/login',
        name: "Login",
        component: Login
    },
    {
        path: '/accommodations',
        name: "Accommodations",
        component: Accommodations
    },
    {
        path: '/activities',
        name: "Activities",
        component: Activities
    },
    {
        path: '/events',
        name: "Events",
        component: Events
    },

]

export const router = createRouter({
    history: createWebHistory(),
    routes
})

