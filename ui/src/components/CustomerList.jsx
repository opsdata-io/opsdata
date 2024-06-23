// CustomerList.jsx

import React, { useState, useEffect } from 'react';
import { getCustomers } from '../utils/api'; // Adjusted import

const CustomerList = () => {
    const [customers, setCustomers] = useState([]);

    useEffect(() => {
        const fetchCustomers = async () => {
            try {
                const response = await getCustomers(); // Use getCustomers function to fetch customers
                setCustomers(response.data); // Assuming response contains customer data
            } catch (error) {
                console.error('Error fetching customers:', error);
            }
        };
        fetchCustomers();
    }, []);

    return (
        <div>
            <h2>Customer List</h2>
            <ul>
                {customers.map(customer => (
                    <li key={customer.id}>{customer.name}</li>
                ))}
            </ul>
        </div>
    );
};

export default CustomerList;
