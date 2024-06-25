// components/FileUploadForm.jsx

import React, { useState } from 'react';
import { postUploadLink } from '../utils/api';
import { getToken } from '../utils/jwt'; // Import getToken to retrieve JWT token
import CustomerDropdown from './CustomerDropdown';

const FileUploadForm = () => {
    const [customer, setCustomer] = useState('');
    const [caseNumber, setCaseNumber] = useState('');
    const [subject, setSubject] = useState('');
    const [notes, setNotes] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const data = { customer, caseNumber, subject, notes };
            const token = getToken(); // Retrieve the JWT token
            const response = await postUploadLink(data, token);
            console.log('Upload link created:', response);
            // Handle success or navigate to confirmation page
        } catch (error) {
            console.error('Error creating upload link:', error);
            // Handle error
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <label>Customer:</label>
            <CustomerDropdown value={customer} onChange={setCustomer} /> {/* Use CustomerDropdown here */}
            <br />
            <label>Case Number:</label>
            <input type="text" value={caseNumber} onChange={(e) => setCaseNumber(e.target.value)} required />
            <br />
            <label>Subject:</label>
            <input type="text" value={subject} onChange={(e) => setSubject(e.target.value)} required />
            <br />
            <label>Notes:</label>
            <textarea value={notes} onChange={(e) => setNotes(e.target.value)} required />
            <br />
            <button type="submit">Create Upload Link</button>
        </form>
    );
};

export default FileUploadForm;
