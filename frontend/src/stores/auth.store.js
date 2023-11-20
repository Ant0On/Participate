import { defineStore } from 'pinia';

import { fetchWrapper } from '@/_helpers/fetch-wrapper';
import { router} from "@/router";

export const useAuthStore = defineStore({
    id: 'auth',
    state: () => ({
        user: JSON.parse(localStorage.getItem('user')),
        returnUrl: null
    }),
    actions: {
        async login(username, password) {
            const user = await fetchWrapper.post('/api/login', { username, password });

            this.user = user;

            localStorage.setItem('user', JSON.stringify(user));

            router.push(this.returnUrl || '/');
        },
        logout() {
            this.user = null;
            localStorage.removeItem('user');
            router.push('/login');
        }
    }
});

