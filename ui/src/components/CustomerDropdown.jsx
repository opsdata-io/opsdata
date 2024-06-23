// CustomerDropdown.jsx

import React, { useState } from 'react';
import { createCustomer } from '../utils/api';

const CustomerDropdown = ({ value, onChange }) => {
    const [newCustomerName, setNewCustomerName] = useState('');

    const handleCreateCustomer = async (e) => {
        e.preventDefault();
        try {
            const newCustomer = await createCustomer({ name: newCustomerName });
            // Optionally update dropdown list or perform other actions
            console.log('New customer created:', newCustomer);
            setNewCustomerName(''); // Clear input after successful creation
        } catch (error) {
            console.error('Error creating customer:', error);
            // Handle error
        }
    };

    return (
        <div>
            <select value={value} onChange={(e) => onChange(e.target.value)}>
                {/* Render customer options here */}
            </select>
            <form onSubmit={handleCreateCustomer}>
                <input
                    type="text"
                    value={newCustomerName}
                    onChange={(e) => setNewCustomerName(e.target.value)}
                    placeholder="New Customer Name"
                />
                <button type="submit">Create Customer</button>
            </form>
        </div>
    );
};

export default CustomerDropdown;
