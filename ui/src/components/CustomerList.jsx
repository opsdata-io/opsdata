import React, { useState, useEffect, useContext } from 'react';
import { getCustomers } from '../utils/api';
import { AuthContext } from '../context/AuthContext';

const CustomerList = () => {
    const [customers, setCustomers] = useState([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState('');
    const { token } = useContext(AuthContext);

    useEffect(() => {
        const fetchCustomers = async () => {
            setLoading(true);
            try {
                const customersData = await getCustomers(token);
                setCustomers(customersData); // Assuming customersData is directly usable
                setError(''); // Clear any previous error messages
            } catch (error) {
                console.error('Error fetching customers:', error);
                setError('Failed to fetch customers.');
            } finally {
                setLoading(false);
            }
        };
        fetchCustomers();
    }, [token]);

    if (loading) {
        return <p>Loading customers...</p>;
    }

    if (error) {
        return <p>Error: {error}</p>;
    }

    return (
        <div>
            <h2>Customer List</h2>
            {customers.length > 0 ? (
                <ul>
                    {customers.map(customer => (
                        <li key={customer.id}>{customer.name}</li>
                    ))}
                </ul>
            ) : (
                <p>No customers found.</p>
            )}
        </div>
    );
};

export default CustomerList;
