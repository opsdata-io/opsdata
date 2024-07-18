import React, { useState, useEffect } from 'react';
import { getVersion } from '../utils/api';
import { getToken } from '../utils/jwt';

const Dashboard = ({ token }) => {
    const [versionInfo, setVersionInfo] = useState('');
    const [error, setError] = useState(null);
    const [loading, setLoading] = useState(true);

    const jwtToken = token || getToken();

    useEffect(() => {
        const fetchVersion = async () => {
            try {
                const response = await getVersion(jwtToken);
                if (response.ok) {
                    const data = await response.json();
                    setVersionInfo(data.version);
                    setError(null);
                } else {
                    throw new Error('Failed to fetch version information');
                }
            } catch (error) {
                console.error('Error fetching version:', error);
                setError('Unable to fetch version information.');
            }
            setLoading(false);
        };
        fetchVersion();
    }, [jwtToken]);

    if (loading) {
        return <p>Loading...</p>;
    }

    return (
        <div>
            <h1>Dashboard</h1>
            {error ? <p>{error}</p> :
                <div style={{ position: 'absolute', bottom: 10, left: 10 }}>
                    {versionInfo || 'Version information not available'}
                </div>}
        </div>
    );
};

export default Dashboard;
