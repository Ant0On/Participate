import { defineStore } from 'pinia';

import { fetchWrapper } from '@/_helpers/fetch-wrapper';

export const useSearchStore = defineStore({
    id: 'search',
    state: () => ({
        location: JSON.parse(localStorage.getItem('location')),
        dateFrom: JSON.parse(localStorage.getItem('dateFrom')),
        dateTo: JSON.parse(localStorage.getItem('dateTo')),
        numberOfPeople: JSON.parse(localStorage.getItem('numberOfPeople')),
    }),
    actions: {
         setSearchValues(searchValues) {
            console.log("values", searchValues)
            localStorage.setItem('location', JSON.stringify(searchValues.location));
            localStorage.setItem('dateFrom', JSON.stringify(searchValues.dateFrom));
            localStorage.setItem('dateTo', JSON.stringify(searchValues.dateTo));
            localStorage.setItem('numberOfPeople', JSON.stringify(searchValues.numberOfPeople));

        },
    }
});

