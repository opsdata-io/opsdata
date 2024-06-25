// components/CustomerList.jsx

import React, { useState, useEffect } from 'react';
import { getCustomers } from '../utils/api';
import { getToken } from '../utils/jwt';

const CustomerList = () => {
    const [customers, setCustomers] = useState([]);

    useEffect(() => {
        const fetchCustomers = async () => {
            const token = getToken(); // Retrieve the JWT token
            try {
                const customersData = await getCustomers(token); // Use getCustomers function to fetch customers
                setCustomers(customersData); // Assuming customersData contains customer data
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
