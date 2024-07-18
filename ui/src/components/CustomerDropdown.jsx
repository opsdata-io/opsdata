import React, { useState, useEffect } from 'react';
import { createCustomer, getCustomers } from '../utils/api';
import { getToken } from '../utils/jwt';

const CustomerDropdown = ({ value, onChange }) => {
    const [customers, setCustomers] = useState([]);
    const [newCustomerName, setNewCustomerName] = useState('');
    const [error, setError] = useState('');

    useEffect(() => {
        const fetchCustomers = async () => {
            const token = getToken();
            try {
                const fetchedCustomers = await getCustomers(token);
                setCustomers(fetchedCustomers);
            } catch (error) {
                console.error('Error fetching customers:', error);
                setError('Failed to fetch customers');
            }
        };

        fetchCustomers();
    }, []);

    const handleCreateCustomer = async (e) => {
        e.preventDefault();
        const token = getToken();
        try {
            const newCustomer = await createCustomer({ name: newCustomerName }, token);
            setCustomers([...customers, newCustomer]);  // Update the customer list locally
            setNewCustomerName('');  // Clear input after successful creation
            onChange(newCustomer.id);  // Optionally update the selected customer
            setError('');  // Reset error message
        } catch (error) {
            console.error('Error creating customer:', error);
            setError('Error creating customer. Please try again.');
        }
    };

    return (
        <div>
            {error && <p style={{ color: 'red' }}>{error}</p>}
            <select value={value} onChange={(e) => onChange(e.target.value)}>
                {customers.map(customer => (
                    <option key={customer.id} value={customer.id}>
                        {customer.name}
                    </option>
                ))}
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
