import React from 'react';
import { useNavigate } from 'react-router-dom'; // Import useNavigate instead of useHistory

const ConfirmationPage = () => {
    const navigate = useNavigate(); // useNavigate hook for navigation

    const handleGoBack = () => {
        navigate('/upload'); // navigate replaces history.push
    };

    const handleGoHome = () => {
        navigate('/'); // navigate replaces history.push
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
