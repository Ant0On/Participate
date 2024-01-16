import {defineStore} from 'pinia';

import {fetchWrapper} from '@/_helpers/fetch-wrapper';
import {router} from "@/router";

export const useAuthStore = defineStore({
    id: 'auth',
    state: () => ({
        user: JSON.parse(localStorage.getItem('user')),
    }),
    actions: {
        async login(email, password) {
            const user = await fetchWrapper.post('/api/login', {email, password});

            this.user = user;

            localStorage.setItem('user', JSON.stringify(user));

            router.push('/');
        },
        logout() {
            this.user = null;
            localStorage.removeItem('user');
            router.push('/login');
        },
        async signUp(name, email, password) {
            const request = await fetchWrapper.post('/api/register/customer', {
                "first_name": name,
                "email": email,
                "password": password
            })

            if (request.message === "registration success!") {
                const user = await fetchWrapper.post('/api/login', {email, password});
                localStorage.setItem('user', JSON.stringify(user));
                router.push('/')

            }


        }
    }
});