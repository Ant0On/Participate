import {defineStore} from 'pinia';

import {fetchWrapper} from '@/_helpers/fetch-wrapper';
import {router} from "@/router";

export const useAuthStore = defineStore({
    id: 'auth',
    state: () => ({
        user: JSON.parse(localStorage.getItem('user')),
        token: JSON.parse(localStorage.getItem('token')),
    }),
    actions: {
        async login(email, password) {
            const response = await fetchWrapper.post('/api/login', {email, password});

            this.user = response.user;
            this.token = response.token;

            localStorage.setItem('user', JSON.stringify(response.user));
            localStorage.setItem('token', JSON.stringify(response.token));

            router.push('/');
        },
        async logout() {
            this.user = null;
            this.token = null;
            localStorage.removeItem('user');
            localStorage.removeItem('token');
            router.push('/');
        },
        async signUp(firstName, lastName, email, password, passwordConfirmation, image) {
            const request = await fetchWrapper.post('/api/register', {
                "first_name": firstName,
                "last_name": lastName,
                "email": email,
                "password": password,
                "password_confirmation": passwordConfirmation,
                "image": image,
            }, "multipart/form-data")

            if (request.message === "registration success!") {
                const user = await fetchWrapper.post('/api/login', {email, password});
                localStorage.setItem('user', JSON.stringify(user));
                alert('Registration successful! You will now be redirected to the homepage.');
                router.push('/')
            }
        },
        async saveToLocalStorage(){
            localStorage.setItem('user', JSON.stringify(this.user));
        }
    }
});