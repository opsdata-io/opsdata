// src/context/AuthContext.js
import React, { createContext, useContext, useState, useEffect } from 'react';
import { getToken, setToken, removeToken } from '../utils/jwt';

const AuthContext = createContext();

export const useAuth = () => useContext(AuthContext);

export const AuthProvider = ({ children }) => {
    const [token, setAuthToken] = useState(getToken());

    const login = (newToken) => {
        setToken(newToken);
        setAuthToken(newToken);
    };

    const logout = () => {
        removeToken();
        setAuthToken(null);
    };

    useEffect(() => {
        // Optionally check token validity or perform refresh here
        setAuthToken(getToken());
    }, []);

    return (
        <AuthContext.Provider value={{ token, login, logout }}>
            {children}
        </AuthContext.Provider>
    );
};
