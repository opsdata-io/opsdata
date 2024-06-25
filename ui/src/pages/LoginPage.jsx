// pages/LoginPage.jsx

import React, { useState } from 'react';
import { useHistory } from 'react-router-dom';
import { login } from '../utils/api';
import { setToken as saveToken } from '../utils/jwt';

const LoginPage = ({ setToken }) => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState(''); // State to hold error messages
    const history = useHistory();

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await login({ email, password });
            const token = response.data.token;
            saveToken(token);
            setToken(token);
            history.push('/dashboard');  // Redirect to dashboard after successful login
        } catch (error) {
            setError('Login failed. Please check your credentials and try again.'); // Set error state
        }
    };

    return (
        <div>
            <h2>Login</h2>
            {error && <div style={{ color: 'red' }}>{error}</div>} {/* Display error message */}
            <form onSubmit={handleSubmit}>
                <div>
                    <label>Email:</label>
                    <input
                        type="email"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        placeholder="Email"
                        required
                    />
                </div>
                <div>
                    <label>Password:</label>
                    <input
                        type="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        placeholder="Password"
                        required
                    />
                </div>
                <button type="submit">Login</button>
            </form>
        </div>
    );
};

export default LoginPage;
