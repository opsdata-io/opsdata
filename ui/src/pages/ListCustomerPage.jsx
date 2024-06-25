// pages/ListCustomerPage.jsx

import React, { useState, useEffect } from 'react';

const ListCustomerPage = () => {
    const [customers, setCustomers] = useState([]);

    useEffect(() => {
        fetchCustomers();
    }, []);

    const fetchCustomers = async () => {
        const token = localStorage.getItem('jwtToken'); // Retrieve the JWT token from local storage
        try {
            const response = await fetch(`/api/customers`, {
                headers: {
                    'Authorization': `Bearer ${token}`, // Add the token to the request headers
                },
            });
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
