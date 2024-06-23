import React from 'react';
import FileDownloadList from '../components/FileDownloadList';

const DownloadPage = () => {
    const token = localStorage.getItem('token');

    return (
        <div>
            <h1>Download Files</h1>
            <FileDownloadList token={token} />
        </div>
    );
};

export default DownloadPage;
