import React, { useEffect, useState } from 'react';
import { downloadFiles } from '../utils/api';
import { useAuth } from '../context/AuthContext'; // Import useAuth hook

const FileDownloadList = () => {
    const [files, setFiles] = useState([]);
    const [error, setError] = useState('');
    const [loading, setLoading] = useState(false);
    const { token } = useAuth(); // Use token from AuthContext via useAuth hook

    useEffect(() => {
        const fetchFiles = async () => {
            setLoading(true);
            try {
                const filesData = await downloadFiles(token); // Ensure downloadFiles uses the token properly
                setFiles(filesData.data); // Assuming filesData comes in a data attribute
                setError(''); // Clear previous errors
            } catch (error) {
                console.error('Error fetching files:', error);
                setError('Failed to load files. Please try again later.');
            } finally {
                setLoading(false);
            }
        };
        fetchFiles();
    }, [token]);

    return (
        <div>
            <h2>Files</h2>
            {loading ? (
                <p>Loading files...</p>
            ) : error ? (
                <p>{error}</p>
            ) : (
                <ul>
                    {files.map(file => (
                        <li key={file.id}>
                            <a href={file.downloadLink} target="_blank" rel="noopener noreferrer">
                                {file.fileName}
                            </a>
                        </li>
                    ))}
                </ul>
            )}
        </div>
    );
};

export default FileDownloadList;
