import React, { useState } from 'react';
import { postUploadLink } from '../utils/api';
import { getToken } from '../utils/jwt'; // Import getToken to retrieve JWT token
import CustomerDropdown from './CustomerDropdown';

const FileUploadForm = () => {
    const [customer, setCustomer] = useState('');
    const [caseNumber, setCaseNumber] = useState('');
    const [subject, setSubject] = useState('');
    const [notes, setNotes] = useState('');
    const [error, setError] = useState('');
    const [success, setSuccess] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError('');
        setSuccess('');
        try {
            const data = { customer, caseNumber, subject, notes };
            const token = getToken(); // Retrieve the JWT token
            const response = await postUploadLink(data, token);
            console.log('Upload link created:', response);
            setSuccess('Upload link created successfully.'); // Set success message
            // Optionally reset the form here or navigate to confirmation page
            setCustomer('');
            setCaseNumber('');
            setSubject('');
            setNotes('');
        } catch (error) {
            console.error('Error creating upload link:', error);
            setError('Failed to create upload link. Please try again.'); // Set error message
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <label htmlFor="customer">Customer:</label>
            <CustomerDropdown value={customer} onChange={setCustomer} />
            <br />
            <label htmlFor="caseNumber">Case Number:</label>
            <input id="caseNumber" type="text" value={caseNumber} onChange={(e) => setCaseNumber(e.target.value)} required />
            <br />
            <label htmlFor="subject">Subject:</label>
            <input id="subject" type="text" value={subject} onChange={(e) => setSubject(e.target.value)} required />
            <br />
            <label htmlFor="notes">Notes:</label>
            <textarea id="notes" value={notes} onChange={(e) => setNotes(e.target.value)} required />
            <br />
            <button type="submit">Create Upload Link</button>
            {error && <div style={{ color: 'red' }}>{error}</div>}
            {success && <div style={{ color: 'green' }}>{success}</div>}
        </form>
    );
};

export default FileUploadForm;
