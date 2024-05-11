import { useAuthStore } from '@/stores/auth.store';

export const fetchWrapper = {
    get: request('GET'),
    post: request('POST'),
    put: request('PUT'),
    delete: request('DELETE')
};

function request(method) {
    return (url, data, contentType = 'application/json') => {
        const requestOptions = {
            method,
            headers: authHeader(url)
        };

        if (data) {
            console.log('data:', data)
            if (method === 'GET') {
                const queryParams = new URLSearchParams(data);
                url = `${url}?${queryParams}`;
            } else {
                if (contentType === 'application/json') {
                    requestOptions.headers['Content-Type'] = 'application/json';
                    requestOptions.body = JSON.stringify(data);
                } else if (contentType === 'multipart/form-data') {
                    const formData = new FormData();
                    for (const key in data) {
                        if (Array.isArray(data[key])) {
                            data[key].forEach((item) => {
                                formData.append(`${key}`, item);
                            });
                        } else {
                            formData.append(key, data[key]);
                        }
                    }
                    requestOptions.body = formData;
                }
            }
        }

        return fetch(url, requestOptions).then(handleResponse);
    }
}

function authHeader(url) {
    const { token } = useAuthStore();
    if (typeof token !== "undefined") {
        return { Authorization: `Bearer ${token}` };
    } else {
        return {};
    }
}

function handleResponse(response) {
    return response.text().then(text => {
        const data = text && text !== "Unauthorized" && JSON.parse(text);
        if (!response.ok) {
            const { user, logout } = useAuthStore();
            if ([401, 403].includes(response.status) && user) {
                logout();
            }

            const error = (data && data.message) || response.statusText;
            return Promise.reject(error);
        }

        return data;
    });
}