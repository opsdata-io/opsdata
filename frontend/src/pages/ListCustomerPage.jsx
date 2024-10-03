import React, { useState, useEffect, useContext } from 'react';
import { AuthContext } from '../context/AuthContext';  // Assuming you have a context for auth

const ListCustomerPage = () => {
    const [customers, setCustomers] = useState([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);
    const { token } = useContext(AuthContext);  // Retrieve token from AuthContext

    useEffect(() => {
        fetchCustomers();
    }, []); // Dependency array left empty to mimic componentDidMount

    const fetchCustomers = async () => {
        setLoading(true);
        try {
            const response = await fetch(`/api/customers`, {
                headers: {
                    'Authorization': `Bearer ${token}`, // Use context token
                },
            });
            if (response.ok) {
                const data = await response.json();
                setCustomers(data);
                setError(null); // Clear any previous errors
            } else {
                throw new Error('Failed to fetch customers');
            }
        } catch (error) {
            console.error('Failed to fetch customers', error);
            setError(error.message);
        } finally {
            setLoading(false); // Ensure loading is false after fetch
        }
    };

    if (loading) {
        return <p>Loading...</p>;
    }

    if (error) {
        return <p>Error: {error}</p>;
    }

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
