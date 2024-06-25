// pages/AddCustomerPage.jsx

import React, { useState } from 'react';
import { useHistory } from 'react-router-dom';
import { getToken } from '../utils/jwt'; // Import getToken from jwt utility

const AddCustomerPage = ({ token }) => {
    const history = useHistory();
    const [customerData, setCustomerData] = useState({
        companyName: '',
        address: '',
        notes: '',
        subscriptionStatus: '',
    });
    const [error, setError] = useState(null); // State to hold error messages

    const handleChange = (e) => {
        const { name, value } = e.target;
        setCustomerData({ ...customerData, [name]: value });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        const jwtToken = token || getToken(); // Retrieve the JWT token using the utility function
        try {
            const response = await fetch(`/api/customers`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    Authorization: `Bearer ${jwtToken}`,
                },
                body: JSON.stringify(customerData),
            });
            if (response.ok) {
                const newCustomer = await response.json();
                console.log('New customer:', newCustomer); // Handle success as needed
                // Redirect to CustomerDetailsPage for the new customer
                history.push(`/customer/${newCustomer.id}`); // Assuming newCustomer.id exists
            } else {
                // Handle error response
                const errorResponse = await response.json();
                setError(errorResponse.error); // Set error state with error message
            }
        } catch (error) {
            console.error('Error creating customer:', error); // Log any unexpected errors
            setError('Failed to create customer. Please try again.'); // Generic error message
        }
    };

    return (
        <div>
            <h2>Customer Management</h2>
            <h3>Add Customer</h3>
            {error && <div style={{ color: 'red' }}>{error}</div>} {/* Display error message */}
            <form onSubmit={handleSubmit}>
                <table style={{ width: '100%' }}>
                    <tbody>
                        <tr>
                            <td style={{ textAlign: 'right', paddingRight: '1rem' }}>
                                <label>Company Name:</label>
                            </td>
                            <td>
                                <input type="text" name="companyName" value={customerData.companyName} onChange={handleChange} required />
                            </td>
                        </tr>
                        <tr>
                            <td style={{ textAlign: 'right', paddingRight: '1rem' }}>
                                <label>Address:</label>
                            </td>
                            <td>
                                <input type="text" name="address" value={customerData.address} onChange={handleChange} required />
                            </td>
                        </tr>
                        <tr>
                            <td style={{ textAlign: 'right', paddingRight: '1rem' }}>
                                <label>Subscription Status:</label>
                            </td>
                            <td>
                                <input type="text" name="subscriptionStatus" value={customerData.subscriptionStatus} onChange={handleChange} required />
                            </td>
                        </tr>
                        <tr>
                            <td style={{ textAlign: 'center', paddingRight: '1rem' }} colSpan={2}>
                                <label>Notes:</label>
                            </td>
                        </tr>
                        <tr>
                            <td style={{ textAlign: 'center', paddingRight: '1rem' }} colSpan={2}>
                                <textarea name="notes" value={customerData.notes} onChange={handleChange} rows="4"></textarea>
                            </td>
                        </tr>
                    </tbody>
                </table>
                <button type="submit" style={{ marginTop: '1rem' }}>Submit</button>
            </form>
        </div>
    );
};

export default AddCustomerPage;
