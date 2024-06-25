// pages/SearchCustomerPage.jsx

import React, { useState } from 'react';
import { Link } from 'react-router-dom';

const SearchCustomerPage = () => {
    const [searchQuery, setSearchQuery] = useState('');
    const [searchResults, setSearchResults] = useState([]);
    const [error, setError] = useState(''); // State to hold error messages

    const handleSearch = async (e) => {
        e.preventDefault();
        const token = localStorage.getItem('jwtToken');
        try {
            const response = await fetch(`/api/customers/search?q=${encodeURIComponent(searchQuery)}`, {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            });
            if (response.ok) {
                const data = await response.json();
                setSearchResults(data);
                setError(''); // Clear any previous errors
            } else {
                console.error('Failed to fetch search results');
                setError('Failed to fetch search results'); // Set error state with error message
            }
        } catch (error) {
            console.error('Failed to fetch search results', error);
            setError('Failed to fetch search results'); // Set error state with error message
        }
    };

    const handleChange = (e) => {
        setSearchQuery(e.target.value);
    };

    const handleKeyPress = (e) => {
        if (e.key === 'Enter') {
            handleSearch(e);
        }
    };

    return (
        <div style={{ marginTop: '20px', marginLeft: '20px' }}>
            <h2>Customers</h2>
            {error && <div style={{ color: 'red' }}>{error}</div>} {/* Display error message */}
            <form onSubmit={handleSearch} style={{ display: 'flex', alignItems: 'center', marginBottom: '20px' }}>
                <input
                    type="text"
                    id="searchQuery"
                    value={searchQuery}
                    onChange={handleChange}
                    onKeyPress={handleKeyPress} // Handle Enter key press
                    required
                    style={{ marginRight: '0.5rem' }}
                />
                <button type="submit">Search</button>
            </form>
            <div>
                {searchResults.length === 0 ? (
                    <p>No results found</p>
                ) : (
                    <table style={{ borderCollapse: 'collapse', width: '100%', border: '1px solid #ccc' }}>
                        <thead style={{ backgroundColor: '#f2f2f2' }}>
                            <tr>
                                <th style={{ border: '1px solid #ccc', padding: '8px' }}>Company Name</th>
                            </tr>
                        </thead>
                        <tbody>
                            {searchResults.map((customer) => (
                                <tr key={customer.id} style={{ border: '1px solid #ccc' }}>
                                    <td style={{ border: '1px solid #ccc', padding: '8px' }}>
                                        <Link to={`/customer/${customer.id}`}>{customer.companyName}</Link>
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                )}
            </div>
        </div>
    );
};

export default SearchCustomerPage;
