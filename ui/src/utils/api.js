// api.js

import axios from 'axios';

// Get API_URL from environment variable or default to localhost
export const API_URL = process.env.REACT_APP_BACKEND_API || 'http://localhost:8080';

// Log the API_URL to verify it
console.log("API_URL:", API_URL);

export const login = async (credentials) => {
    return await axios.post(`${API_URL}/login`, credentials);
};

export const createUploadLink = async (token) => {
    return await axios.post(
        `${API_URL}/create-link`,
        {},
        {
            headers: {
                Authorization: `Bearer ${token}`,
            },
        }
    );
};

export const uploadFile = async (link, file, token) => {
    const formData = new FormData();
    formData.append('file', file);

    return await axios.post(`${API_URL}/upload/${link}`, formData, {
        headers: {
            'Content-Type': 'multipart/form-data',
            Authorization: `Bearer ${token}`,
        },
    });
};

export const downloadFiles = async (token) => {
    return await axios.get(`${API_URL}/files`, {
        headers: {
            Authorization: `Bearer ${token}`,
        },
    });
};

export const getCustomers = async (token) => {
    return await axios.get(`${API_URL}/api/customers`, {
        headers: {
            Authorization: `Bearer ${token}`,
        },
    });
};

export const createCustomer = async (customerData, token) => {
    return await axios.post(`${API_URL}/api/customers`, customerData, {
        headers: {
            Authorization: `Bearer ${token}`,
        },
    });
};

export const updateCustomer = async (customerId, customerData, token) => {
    return await axios.put(`${API_URL}/api/customers/${customerId}`, customerData, {
        headers: {
            Authorization: `Bearer ${token}`,
        },
    });
};

export const deleteCustomer = async (customerId, token) => {
    return await axios.delete(`${API_URL}/api/customers/${customerId}`, {
        headers: {
            Authorization: `Bearer ${token}`,
        },
    });
};

export const getVersion = async () => {
    return await axios.get(`${API_URL}/api/version`);
};

export const postUploadLink = async (data) => {
    // Implement the postUploadLink function to send data to the backend
};
