import { createRouter, createWebHistory } from 'vue-router'
import {storeToRefs} from 'pinia';

import { useAuthStore } from '@/stores/auth.store';
import {useNavStore} from "@/stores/nav.store";

import Login from '@/views/LoginPage.vue'
import Home from '@/views/HomePage.vue'
import Accommodations from "@/views/AccommodationsPage.vue";
import Events from "@/views/EventsPage.vue";
import Activities from "@/views/ActivitiesPage.vue";
import Offer from "@/views/OfferDetailsPage.vue";
import OfferEdit from '@/views/OfferEditPage.vue';
import AccountBecomeHost from "@/components/account/become_host/AccountBecomeHost.vue";
import ReservationsHistory from "@/components/account/reservations_history/ReservationsHistory.vue";
import AccountMyOffers from "@/components/account/my_offers/AccountMyOffers.vue";
import AccountNewOffer from "@/components/account/new_offer/AccountNewOffer.vue";
import CurrentReservations from "@/components/account/current_reservations/CurrentReservations.vue";

const routes = [
    {
        path: '/',
        name: "Home",
        component: Home
    },
    {
        path: '/login',
        alias: '/account',
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
        path: '/offers/:type/edit/:id',
        name: 'Offer Edit',
        component: OfferEdit,
        props: true
    },
    {
        path: '/account/become_host',
        name: "Become a Host",
        component: AccountBecomeHost
    },
    {
        path: '/account/reservations_history',
        name: "Reservations History",
        component: ReservationsHistory
    },
    {
        path: '/account/my_offers',
        name: "My Offers",
        component: AccountMyOffers
    },
    {
        path: '/account/new_offer',
        name: "New offer",
        component: AccountNewOffer
    },
     {
        path: '/account/current_reservations',
        name: "Current reservations",
        component: CurrentReservations
    },


]

export const router = createRouter({
    history: createWebHistory(),
    routes
})

router.beforeEach(async (to) => {
    const hostPages = ['/account/my_offers', '/account/new_offer', '/account/current_reservations'];
    const customerPages = ['/account/become_host', '/account/reservations_history'];
    const hostRequired = hostPages.includes(to.path);
    const customerRequired = customerPages.includes(to.path);
    const userStore = useAuthStore();
    const navStore = useNavStore();
    const {user: user} = storeToRefs(userStore)

    if (hostRequired && user.value?.Role !== "host" || customerRequired && user.value?.Role !== "customer") {
        router.push('/')
        user.value.returnUrl = to.fullPath;
        navStore.changePage('/')
        return
    }
    navStore.changePage(to)
});