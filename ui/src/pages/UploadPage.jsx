import React from 'react';
import FileUploadForm from '../components/FileUploadForm';

const UploadPage = ({ match }) => {
    const { link } = match.params;
    const token = localStorage.getItem('token');

    return (
        <div>
            <h1>Upload your file</h1>
            <FileUploadForm link={link} token={token} />
        </div>
    );
};

export default UploadPage;
