// pages/CustomerDetailsPage.jsx

import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';

const CustomerDetailsPage = () => {
    const { id } = useParams();
    const [customer, setCustomer] = useState(null);
    const [error, setError] = useState(null); // State to hold error messages
    const [loading, setLoading] = useState(true); // State to handle loading

    useEffect(() => {
        const fetchCustomer = async () => {
            const token = localStorage.getItem('jwtToken'); // Retrieve the JWT token from local storage
            try {
                const response = await fetch(`/api/customers/${id}`, {
                    headers: {
                        'Authorization': `Bearer ${token}` // Add the token to the request headers
                    }
                });
                if (response.ok) {
                    const data = await response.json();
                    setCustomer(data);
                } else {
                    setError('Failed to fetch customer details');
                }
            } catch (error) {
                setError('Failed to fetch customer details');
            } finally {
                setLoading(false); // Set loading to false regardless of the outcome
            }
        };

        fetchCustomer();
    }, [id]);

    if (loading) {
        return <p>Loading...</p>;
    }

    if (error) {
        return <p style={{ color: 'red' }}>{error}</p>;
    }

    return (
        <div style={{ marginTop: '20px', marginLeft: '20px' }}>
            <h2>Customer Details</h2>
            <table style={{ width: '100%', borderCollapse: 'collapse' }}>
                <tbody>
                    <tr style={{ borderBottom: '1px solid lightgrey' }}>
                        <td style={{ padding: '8px' }}><strong>Company Name:</strong></td>
                        <td style={{ padding: '8px' }}>{customer.companyName}</td>
                    </tr>
                    <tr style={{ borderBottom: '1px solid lightgrey' }}>
                        <td style={{ padding: '8px' }}><strong>Address:</strong></td>
                        <td style={{ padding: '8px' }}>{customer.address}</td>
                    </tr>
                    <tr style={{ borderBottom: '1px solid lightgrey' }}>
                        <td style={{ padding: '8px' }}><strong>Contact Name:</strong></td>
                        <td style={{ padding: '8px' }}>{customer.contactName}</td>
                    </tr>
                    <tr style={{ borderBottom: '1px solid lightgrey' }}>
                        <td style={{ padding: '8px' }}><strong>Contact Title:</strong></td>
                        <td style={{ padding: '8px' }}>{customer.contactTitle}</td>
                    </tr>
                    <tr style={{ borderBottom: '1px solid lightgrey' }}>
                        <td style={{ padding: '8px' }}><strong>Contact Email:</strong></td>
                        <td style={{ padding: '8px' }}>{customer.contactEmail}</td>
                    </tr>
                    <tr style={{ borderBottom: '1px solid lightgrey' }}>
                        <td style={{ padding: '8px' }}><strong>Contact Phone:</strong></td>
                        <td style={{ padding: '8px' }}>{customer.contactPhone}</td>
                    </tr>
                    <tr style={{ borderBottom: '1px solid lightgrey' }}>
                        <td style={{ padding: '8px' }}><strong>Subscription Status:</strong></td>
                        <td style={{ padding: '8px' }}>{customer.subscriptionStatus}</td>
                    </tr>
                    <tr>
                        <td style={{ padding: '8px' }}><strong>Notes:</strong></td>
                        <td style={{ padding: '8px' }}>{customer.notes}</td>
                    </tr>
                </tbody>
            </table>
        </div>
    );
};

export default CustomerDetailsPage;
