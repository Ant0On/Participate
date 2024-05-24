import { defineStore } from 'pinia';

export const useNavStore = defineStore({
    id: 'nav',
    state: () => ({
        currentPage: JSON.parse(localStorage.getItem('currentPage')) || 'home'
    }),
    actions: {
        changePage(pageTo){
            if(pageTo)
            {
                const page = pageTo?.fullPath.split('/')[1] || 'home'
                localStorage.setItem('currentPage', JSON.stringify(page));
                this.currentPage = page;
            }
        }
    }

});

