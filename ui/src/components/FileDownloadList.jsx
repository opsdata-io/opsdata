// components/FileDownloadList.jsx

import React, { useEffect, useState } from 'react';
import { downloadFiles } from '../utils/api';
import { getToken } from '../utils/jwt';

const FileDownloadList = () => {
    const [files, setFiles] = useState([]);
    const token = getToken(); // Retrieve the JWT token

    useEffect(() => {
        const fetchFiles = async () => {
            try {
                const filesData = await downloadFiles(token);
                setFiles(filesData);
            } catch (error) {
                console.error('Error fetching files:', error);
            }
        };
        fetchFiles();
    }, [token]);

    return (
        <div>
            <h2>Files</h2>
            <ul>
                {files.map(file => (
                    <li key={file.id}>
                        <a href={file.downloadLink} target="_blank" rel="noopener noreferrer">
                            {file.fileName}
                        </a>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default FileDownloadList;
