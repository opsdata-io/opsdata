import React, { useState, useEffect } from 'react';
import { API_URL } from '../utils/api';

const ListCustomerPage = () => {
    const [customers, setCustomers] = useState([]);

    useEffect(() => {
        fetchCustomers();
    }, []);

    const fetchCustomers = async () => {
        try {
            const response = await fetch(`/api/v1/customers`);
            if (response.ok) {
                const data = await response.json();
                setCustomers(data); // Update customers state with fetched data
            } else {
                console.error('Failed to fetch customers');
            }
        } catch (error) {
            console.error('Failed to fetch customers', error);
        }
    };

    return (
        <div>
            <h1>List Customers</h1>
            <table>
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Name</th>
                        <th>Email</th>
                    </tr>
                </thead>
                <tbody>
                    {customers.map(customer => (
                        <tr key={customer.id}>
                            <td>{customer.id}</td>
                            <td>{customer.name}</td>
                            <td>{customer.email}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
};

export default ListCustomerPage;
