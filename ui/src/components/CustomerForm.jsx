import React, { useState } from 'react';
import { createCustomer } from '../utils/api';
import { getToken } from '../utils/jwt';

const CustomerForm = ({ onSubmit }) => {
    const [name, setName] = useState('');
    const [error, setError] = useState('');
    const [isLoading, setIsLoading] = useState(false);

    const handleSubmit = async (e) => {
        e.preventDefault();
        if (!name.trim()) {
            setError('Customer name cannot be empty');
            return;
        }
        setIsLoading(true);
        const token = getToken(); // Retrieve the JWT token
        try {
            const response = await createCustomer({ name }, token);
            if (response.ok) {
                const customer = await response.json();
                onSubmit(customer); // Pass the new customer up to the parent component
                setName(''); // Clear form on successful creation
                setError(''); // Clear any previous errors
            } else {
                throw new Error('Failed to create customer');
            }
        } catch (error) {
            console.error('Error creating customer:', error);
            setError('Error creating customer. Please try again.');
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <label htmlFor="customerName">Customer Name:</label>
            <input
                id="customerName"
                type="text"
                placeholder="Customer Name"
                value={name}
                onChange={(e) => setName(e.target.value)}
                disabled={isLoading}
            />
            <button type="submit" disabled={isLoading || !name.trim()}>
                {isLoading ? 'Creating...' : 'Create Customer'}
            </button>
            {error && <p style={{ color: 'red' }}>{error}</p>}
        </form>
    );
};

export default CustomerForm;
