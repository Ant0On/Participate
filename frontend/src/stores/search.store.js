import { defineStore } from 'pinia';

export const useSearchStore = defineStore({
    id: 'search',
    state: () => ({
        location: '',
        dateFrom: new Date(),
        dateTo: new Date(),
        numberOfPeople: 0,
    }),
});

