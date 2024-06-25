// pages/UploadPage.jsx

import React from 'react';
import FileUploadForm from '../components/FileUploadForm';
import { getToken } from '../utils/jwt';

const UploadPage = ({ match }) => {
    const { link } = match.params;
    const token = getToken(); // Use getToken to retrieve the token

    return (
        <div>
            <h1>Upload your file</h1>
            <FileUploadForm link={link} token={token} />
        </div>
    );
};

export default UploadPage;
