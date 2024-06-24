// Dashboard.jsx

import React, { useState, useEffect } from 'react';
import { getVersion } from '../utils/api'; // Removed unused imports

const Dashboard = ({ token }) => {
    const [versionInfo, setVersionInfo] = useState('');

    useEffect(() => {
        const fetchVersion = async () => {
            try {
                const response = await getVersion(); // Assuming getVersion fetches the version info
                setVersionInfo(response.data.version); // Assuming response.data.version contains the version info
            } catch (error) {
                console.error('Error fetching version:', error);
            }
        };
        fetchVersion();
    }, []);

    return (
        <div>
            <h1>Dashboard</h1>
            <div style={{ position: 'absolute', bottom: 10, left: 10 }}>{versionInfo}</div> {/* Corrected style syntax */}
        </div>
    );
};

export default Dashboard;
