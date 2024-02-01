import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth.store';

import Login from '@/views/LoginPage.vue'
import Home from '@/views/HomePage.vue'
import Accommodations from "@/views/AccommodationsPage.vue";
import Events from "@/views/EventsPage.vue";
import Activities from "@/views/ActivitiesPage.vue";
import Offer from "@/views/OfferDetailsPage.vue";
import Recommended from "@/views/RecommendedPage.vue"

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
    {
        path: '/offers/:type/:id',
        name: 'Offers',
        component: Offer,
        props: true
    },
    {
        path: '/recommended',
        name: "Recommended",
        component: Recommended
    },

]

export const router = createRouter({
    history: createWebHistory(),
    routes
})

router.beforeEach(async (to) => {
    const privatePages = [];
    const authRequired = privatePages.includes(to.path);
    const auth = useAuthStore();

    if (authRequired && !auth.user) {
        auth.returnUrl = to.fullPath;
        return '/';
    }
});