// CustomerForm.jsx

import React, { useState } from 'react';
import { createCustomer } from '../utils/api';

const CustomerForm = ({ onSubmit }) => {
    const [name, setName] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await createCustomer({ name });
            onSubmit(response.data); // Handle success, e.g., update dropdown list
            setName(''); // Clear form
        } catch (error) {
            console.error('Error creating customer:', error);
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <input
                type="text"
                placeholder="Customer Name"
                value={name}
                onChange={(e) => setName(e.target.value)}
            />
            <button type="submit">Create Customer</button>
        </form>
    );
};

export default CustomerForm;
