// pages/ConfirmationPage.jsx

import React from 'react';
import { useHistory } from 'react-router-dom';

const ConfirmationPage = () => {
    const history = useHistory();

    const handleGoBack = () => {
        history.push('/upload');
    };

    const handleGoHome = () => {
        history.push('/');
    };

    return (
        <div style={{ textAlign: 'center', marginTop: '20px' }}>
            <h1>File Uploaded Successfully</h1>
            <p>Your file has been uploaded successfully.</p>
            <div>
                <button onClick={handleGoBack} style={{ marginRight: '10px' }}>Upload Another File</button>
                <button onClick={handleGoHome}>Go to Home</button>
            </div>
        </div>
    );
};

export default ConfirmationPage;
