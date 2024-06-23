import React, { useEffect, useState } from 'react';
import { downloadFiles } from '../utils/api';

const FileDownloadList = ({ token }) => {
    const [files, setFiles] = useState([]);

    useEffect(() => {
        const fetchFiles = async () => {
            try {
                const response = await downloadFiles(token);
                setFiles(response.data);
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
