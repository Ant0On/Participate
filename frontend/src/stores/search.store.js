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
                const newSearch = [previousSearch,
                    ...JSON.parse(localStorage.getItem('search'))?.filter(item => item.title !== previousSearch.title) || []]
                this.$patch({previousSearch: newSearch})
                localStorage.setItem('search', JSON.stringify(this.previousSearch));
            }
        }
    }

});

