import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';

const CustomerDetailsPage = ({ token }) => {
    const { id } = useParams();
    const [customer, setCustomer] = useState(null);
    const [error, setError] = useState('');

    useEffect(() => {
        const fetchCustomer = async () => {
            try {
                const response = await fetch(`/v1/customers/${id}`, {
                    headers: {
                        'Authorization': `Bearer ${token}` // Add the token to the request headers
                    }
                });
                if (response.ok) {
                    const data = await response.json();
                    setCustomer(data);
                } else {
                    throw new Error('Failed to fetch customer details');
                }
            } catch (error) {
                console.error(error);
                setError(error.message || 'Failed to load customer data.');
            }
        };

        fetchCustomer();
    }, [id, token]);

    if (error) {
        return <p>Error: {error}</p>; // Display errors if any
    }

    if (!customer) {
        return <p>Loading...</p>; // Loading state
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
