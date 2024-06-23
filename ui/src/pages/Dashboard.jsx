// Dashboard.jsx

import React, { useState, useEffect } from 'react';
import { getCustomers, createCustomer, updateCustomer, deleteCustomer, getVersion } from '../utils/api'; // Import getVersion
import CustomerForm from '../components/CustomerForm';

const Dashboard = ({ token }) => {
    const [customers, setCustomers] = useState([]);
    const [showCustomerForm, setShowCustomerForm] = useState(false);
    const [versionInfo, setVersionInfo] = useState('');

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

    useEffect(() => {
        const fetchVersion = async () => {
            try {
                const response = await getVersion(); // Assuming getVersion fetches the version info
                setVersionInfo(response.data.version); // Assuming response.data.version contains the version info
            } catch (error) {
                console.error('Error fetching version:', error);
            }
        };
        fetchVersion();
    }, []);

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
            <h1>Dashboard</h1>
            <div style={{ position: 'absolute', bottom: 10, left: 10 }}>{versionInfo}</div> {/* Corrected style syntax */}
        </div>
    );
};

export default Dashboard;
