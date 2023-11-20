import { defineStore } from 'pinia';

import { fetchWrapper } from '@/_helpers/fetch-wrapper';


export const useUsersStore = defineStore({
    id: 'users',
    state: () => ({
        users: {}
    }),
    actions: {
        async getAll() {
            this.users = { loading: true };
            fetchWrapper.get('/users')
                .then(users => this.users = users)
                .catch(error => this.users = { error })
        }
    }
});