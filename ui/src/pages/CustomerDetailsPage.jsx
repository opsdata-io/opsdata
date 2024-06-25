import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';

const CustomerDetailsPage = () => {
    const { id } = useParams();
    const [customer, setCustomer] = useState(null);

    useEffect(() => {
        const fetchCustomer = async () => {
            try {
                const response = await fetch(`/api/v1/customers/${id}`);
                if (response.ok) {
                    const data = await response.json();
                    setCustomer(data);
                } else {
                    console.error('Failed to fetch customer details');
                }
            } catch (error) {
                console.error('Failed to fetch customer details', error);
            }
        };

        fetchCustomer();
    }, [id]);

    if (!customer) {
        return <p>Loading...</p>;
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
