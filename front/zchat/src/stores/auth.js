// src/store/auth.js
import { defineStore } from 'pinia';

export const useAuthStore = defineStore('auth', {
    state: () => ({
        token: null,
        expire: null,
    }),
    actions: {
        setToken(token, expire) {
            this.token = token;
            this.expire = expire;
        },
        clearToken() {
            this.token = null;
            this.expire = null;
        },
    },
});
