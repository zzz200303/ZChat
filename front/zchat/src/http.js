import axios from 'axios';
import { useAuthStore } from '@/store/auth';

const http = axios.create({
    baseURL: 'http://localhost:9000',
});

http.interceptors.request.use(config => {
    const authStore = useAuthStore();
    if (authStore.token) {
        config.headers.Authorization = `Bearer ${authStore.token}`;
    }
    return config;
}, error => {
    return Promise.reject(error);
});

export default http;
