import { defineStore } from 'pinia';

export const useSearchStore = defineStore({
    id: 'search',
    state: () => ({
       previousSearch: JSON.parse(localStorage.getItem('search')) || []
    }),
    actions: {
        addPreviousSearch(previousSearch){
            if(previousSearch.title)
            {
                this.$patch({previousSearch: [...this.previousSearch, previousSearch]})
                localStorage.setItem('search', JSON.stringify(this.previousSearch));
            }
        }
    }

});

