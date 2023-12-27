import {defineStore} from 'pinia';

import {fetchWrapper} from '@/_helpers/fetch-wrapper';
import {router} from "@/router";

export const useAuthStore = defineStore({
    id: 'auth',
    state: () => ({
        user: JSON.parse(localStorage.getItem('user')),
        token: JSON.parse(localStorage.getItem('token'))
    }),
    actions: {
        async login(email, password) {
            const user = await fetchWrapper.post('/api/login', {email, password});

            this.user = user;
            localStorage.setItem('user', JSON.stringify(user.user));
            localStorage.setItem('token', JSON.stringify(user.token));
            router.go('/');
        },
        logout() {
            this.user = null;
            localStorage.removeItem('user');
            router.push('/login');
        },
        async signUp(name, email, password) {
            const user = await fetchWrapper.post('/api/register', {name, email, password})
            this.user = user;
            localStorage.setItem('user', JSON.stringify(user));

            router.push('/')
        }
    }
});