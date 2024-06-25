// utils/api.js

import axios from 'axios';
import { getToken } from './jwt';

export const login = async (credentials) => {
    return await axios.post(`/api/login`, credentials);
};

export const createUploadLink = async (token) => {
    return await axios.post(
        `/api/create-link`,
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

    return await axios.post(`/api/upload/${link}`, formData, {
        headers: {
            'Content-Type': 'multipart/form-data',
            Authorization: `Bearer ${token}`,
        },
    });
};

export const downloadFiles = async () => {
    const token = getToken();
    return await axios.get(`/api/files`, {
        headers: {
            Authorization: `Bearer ${token}`,
        },
    });
};

export const getCustomers = async () => {
    const token = getToken();
    return await axios.get(`/api/customers`, {
        headers: {
            Authorization: `Bearer ${token}`,
        },
    });
};

export const createCustomer = async (customerData) => {
    const token = getToken();
    return await axios.post(`/api/customers`, customerData, {
        headers: {
            Authorization: `Bearer ${token}`,
        },
    });
};

export const updateCustomer = async (customerId, customerData) => {
    const token = getToken();
    return await axios.put(`/api/customers/${customerId}`, customerData, {
        headers: {
            Authorization: `Bearer ${token}`,
        },
    });
};

export const deleteCustomer = async (customerId) => {
    const token = getToken();
    return await axios.delete(`/api/customers/${customerId}`, {
        headers: {
            Authorization: `Bearer ${token}`,
        },
    });
};

export const getVersion = async () => {
    return await axios.get(`/api/version`);
};

export const postUploadLink = async (data) => {
    const token = getToken();
    return await axios.post(
        `/api/create-link`,
        data,
        {
            headers: {
                Authorization: `Bearer ${token}`,
            },
        }
    );
};
