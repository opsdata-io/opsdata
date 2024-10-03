import React, { useState, useEffect } from 'react';
import FileDownloadList from '../components/FileDownloadList';

const DownloadPage = () => {
    const [token, setToken] = useState('');

    useEffect(() => {
        // Get the token from local storage and set it to state
        const jwtToken = localStorage.getItem('jwtToken'); // Use jwtToken for consistency
        setToken(jwtToken);
    }, []); // Empty dependency array means this runs once on mount

    return (
        <div>
            <h1>Download Files</h1>
            <FileDownloadList token={token} />
        </div>
    );
};

export default DownloadPage;
