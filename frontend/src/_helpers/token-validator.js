import {useAuthStore} from "@/stores/auth.store";
import {fetchWrapper} from "@/_helpers/fetch-wrapper";

function validateToken() {
    const authStore = useAuthStore();

    fetchWrapper.get('/api/check-token')
        .then(response => {
            if (response.status === 401) {
                authStore.logout()
            }
        }).catch(error => {

        });
}

export { validateToken }