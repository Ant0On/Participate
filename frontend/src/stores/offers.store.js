import { defineStore } from 'pinia';


export const useOfferStore = defineStore({
    id: 'offers',
    state: () => ({
        isLocalization: false,
        inputValue: null
    }),
});

