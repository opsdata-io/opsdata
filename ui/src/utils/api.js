import axios from 'axios';
import { getToken } from './jwt';  // Ensure this path is correct based on your project structure

// Setup axios instance
const api = axios.create({
    baseURL: '/api'
});

// Request interceptor to add the auth token for every request
api.interceptors.request.use(
    async (config) => {
        const token = getToken();
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

// Simplified API functions
export const login = (credentials) => {
    return api.post(`/login`, credentials);
};

export const createUploadLink = () => {
    return api.post(`/create-link`);
};

export const uploadFile = (link, file) => {
    const formData = new FormData();
    formData.append('file', file);
    return api.post(`/upload/${link}`, formData, {
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    });
};

export const downloadFiles = () => {
    return api.get(`/files`);
};

export const getCustomers = () => {
    return api.get(`/customers`);
};

export const createCustomer = (customerData) => {
    return api.post(`/customers`, customerData);
};

export const updateCustomer = (customerId, customerData) => {
    return api.put(`/customers/${customerId}`, customerData);
};

export const deleteCustomer = (customerId) => {
    return api.delete(`/customers/${customerId}`);
};

export const getVersion = () => {
    return api.get(`/version`);
};

export const postUploadLink = (data) => {
    return api.post(`/create-link`, data);
};
