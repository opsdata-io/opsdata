import React, { useState, useEffect } from 'react';
import CustomerList from '../components/CustomerList';
import CustomerForm from '../components/CustomerForm';
import { getCustomers, createCustomer, updateCustomer, deleteCustomer } from '../utils/api';

const CustomerPage = () => {
    const [customers, setCustomers] = useState([]);
    const [showCustomerForm, setShowCustomerForm] = useState(false);
    const [error, setError] = useState(null);

    useEffect(() => {
        const fetchCustomers = async () => {
            try {
                const customersData = await getCustomers();
                setCustomers(customersData);
            } catch (error) {
                console.error('Error fetching customers:', error);
                setError('Error fetching customers. Please try again.');
            }
        };
        fetchCustomers();
    }, []);

    const handleCreateCustomer = async (customerData) => {
        try {
            const newCustomer = await createCustomer(customerData);
            setCustomers(prev => [...prev, newCustomer]);
            setShowCustomerForm(false);
            setError(null); // Reset error on successful operation
        } catch (error) {
            console.error('Error creating customer:', error);
            setError('Error creating customer. Please try again.');
        }
    };

    const handleUpdateCustomer = async (customerId, customerData) => {
        try {
            await updateCustomer(customerId, customerData);
            const updatedCustomers = customers.map(cust => (cust.id === customerId ? { ...cust, ...customerData } : cust));
            setCustomers(updatedCustomers);
            setError(null); // Reset error on successful operation
        } catch (error) {
            console.error('Error updating customer:', error);
            setError('Error updating customer. Please try again.');
        }
    };

    const handleDeleteCustomer = async (customerId) => {
        try {
            await deleteCustomer(customerId);
            const updatedCustomers = customers.filter(cust => cust.id !== customerId);
            setCustomers(updatedCustomers);
            setError(null); // Reset error on successful operation
        } catch (error) {
            console.error('Error deleting customer:', error);
            setError('Error deleting customer. Please try again.');
        }
    };

    return (
        <div>
            <h1>Customer Management</h1>
            {error && <div style={{ color: 'red' }}>{error}</div>}
            <button onClick={() => setShowCustomerForm(!showCustomerForm)}>
                {showCustomerForm ? 'Cancel' : 'Add Customer'}
            </button>
            {showCustomerForm && <CustomerForm onSubmit={handleCreateCustomer} />}
            <CustomerList
                customers={customers}
                onUpdate={handleUpdateCustomer}
                onDelete={handleDeleteCustomer}
            />
        </div>
    );
};

export default CustomerPage;
