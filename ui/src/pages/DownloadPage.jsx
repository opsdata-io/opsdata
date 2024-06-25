// pages/DownloadPage.jsx

import React from 'react';
import FileDownloadList from '../components/FileDownloadList';

const DownloadPage = () => {
    const token = localStorage.getItem('jwtToken'); // Use jwtToken for consistency

    return (
        <div>
            <h1>Download Files</h1>
            <FileDownloadList token={token} />
        </div>
    );
};

export default DownloadPage;
