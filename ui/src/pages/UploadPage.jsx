import React from 'react';
import { useParams } from 'react-router-dom'; // Import useParams for route parameters
import FileUploadForm from '../components/FileUploadForm';
import { getToken } from '../utils/jwt';

const UploadPage = () => {
    const { link } = useParams(); // Use useParams to access route parameters
    const token = getToken(); // Retrieve the JWT token

    return (
        <div>
            <h1>Upload your file</h1>
            <FileUploadForm link={link} token={token} />
        </div>
    );
};

export default UploadPage;
