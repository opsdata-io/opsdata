// CustomerPage.jsx

import React, { useState, useEffect } from 'react';
import CustomerList from '../components/CustomerList';
import CustomerForm from '../components/CustomerForm';
import { getCustomers, createCustomer, updateCustomer, deleteCustomer } from '../utils/api';

const CustomerPage = ({ token }) => {
    const [customers, setCustomers] = useState([]);
    const [showCustomerForm, setShowCustomerForm] = useState(false);

    useEffect(() => {
        const fetchCustomers = async () => {
            try {
                const customersData = await getCustomers(token);
                setCustomers(customersData);
            } catch (error) {
                console.error('Error fetching customers:', error);
            }
        };
        fetchCustomers();
    }, [token]);

    const handleCreateCustomer = async (customerData) => {
        try {
            const newCustomer = await createCustomer(customerData, token);
            setCustomers([...customers, newCustomer]);
            setShowCustomerForm(false); // Hide the form after successful creation
        } catch (error) {
            console.error('Error creating customer:', error);
        }
    };

    const handleUpdateCustomer = async (customerId, customerData) => {
        try {
            await updateCustomer(customerId, customerData, token);
            const updatedCustomers = customers.map(cust => (cust.id === customerId ? { ...cust, ...customerData } : cust));
            setCustomers(updatedCustomers);
            // Optionally, update state or show a notification for successful update
        } catch (error) {
            console.error('Error updating customer:', error);
        }
    };

    const handleDeleteCustomer = async (customerId) => {
        try {
            await deleteCustomer(customerId, token);
            const updatedCustomers = customers.filter(cust => cust.id !== customerId);
            setCustomers(updatedCustomers);
            // Optionally, update state or show a notification for successful deletion
        } catch (error) {
            console.error('Error deleting customer:', error);
        }
    };

    return (
        <div>
            <h1>Customer Management</h1>
            <button onClick={() => setShowCustomerForm(true)}>Add Customer</button>
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
