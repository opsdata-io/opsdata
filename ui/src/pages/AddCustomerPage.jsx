import React, { useState } from 'react';
import { useHistory } from 'react-router-dom';
import { API_URL } from '../utils/api';

const AddCustomerPage = ({ token }) => {
    const history = useHistory();
    const [customerData, setCustomerData] = useState({
        companyName: '',
        address: '',
        contactName: '',
        contactTitle: '',
        contactEmail: '',
        contactPhone: '',
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
        try {
            const response = await fetch(`/api/v1/customers`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    Authorization: `Bearer ${token}`,
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
                                <input type="text" name="companyName" value={customerData.companyName} onChange={handleChange} />
                            </td>
                        </tr>
                        <tr>
                            <td style={{ textAlign: 'right', paddingRight: '1rem' }}>
                                <label>Address:</label>
                            </td>
                            <td>
                                <input type="text" name="address" value={customerData.address} onChange={handleChange} />
                            </td>
                        </tr>
                        <tr>
                            <td style={{ textAlign: 'right', paddingRight: '1rem' }}>
                                <label>Contact Name:</label>
                            </td>
                            <td>
                                <input type="text" name="contactName" value={customerData.contactName} onChange={handleChange} />
                            </td>
                        </tr>
                        <tr>
                            <td style={{ textAlign: 'right', paddingRight: '1rem' }}>
                                <label>Contact Title:</label>
                            </td>
                            <td>
                                <input type="text" name="contactTitle" value={customerData.contactTitle} onChange={handleChange} />
                            </td>
                        </tr>
                        <tr>
                            <td style={{ textAlign: 'right', paddingRight: '1rem' }}>
                                <label>Contact Email:</label>
                            </td>
                            <td>
                                <input type="email" name="contactEmail" value={customerData.contactEmail} onChange={handleChange} />
                            </td>
                        </tr>
                        <tr>
                            <td style={{ textAlign: 'right', paddingRight: '1rem' }}>
                                <label>Contact Phone:</label>
                            </td>
                            <td>
                                <input type="tel" name="contactPhone" value={customerData.contactPhone} onChange={handleChange} />
                            </td>
                        </tr>
                        <tr>
                            <td style={{ textAlign: 'right', paddingRight: '1rem' }}>
                                <label>Subscription Status:</label>
                            </td>
                            <td>
                                <input type="text" name="subscriptionStatus" value={customerData.subscriptionStatus} onChange={handleChange} />
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
