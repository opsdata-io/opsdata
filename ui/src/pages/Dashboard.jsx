// pages/Dashboard.jsx

import React, { useState, useEffect } from 'react';
import { getVersion } from '../utils/api';
import { getToken } from '../utils/jwt';

const Dashboard = ({ token }) => {
    const [versionInfo, setVersionInfo] = useState('');

    const jwtToken = token || getToken(); // Retrieve the JWT token from props or local storage

    useEffect(() => {
        const fetchVersion = async () => {
            try {
                const response = await getVersion(jwtToken); // Pass the token to getVersion
                setVersionInfo(response.data.version); // Assuming response.data.version contains the version info
            } catch (error) {
                console.error('Error fetching version:', error);
            }
        };
        fetchVersion();
    }, [jwtToken]);

    return (
        <div>
            <h1>Dashboard</h1>
            <div style={{ position: 'absolute', bottom: 10, left: 10 }}>{versionInfo}</div> {/* Corrected style syntax */}
        </div>
    );
};

export default Dashboard;
