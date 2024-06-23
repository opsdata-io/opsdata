// SearchCustomerPage.jsx

import React, { useState } from 'react';
import { Link } from 'react-router-dom'; // Import Link from React Router
import { API_URL } from '../utils/api';

const SearchCustomerPage = () => {
    const [searchQuery, setSearchQuery] = useState('');
    const [searchResults, setSearchResults] = useState([]);

    const handleSearch = async (e) => {
        e.preventDefault();
        try {
            const response = await fetch(`${API_URL}/api/customers/search?q=${encodeURIComponent(searchQuery)}`);
            if (response.ok) {
                const data = await response.json();
                setSearchResults(data); // Update search results state with fetched data
            } else {
                console.error('Failed to fetch search results');
            }
        } catch (error) {
            console.error('Failed to fetch search results', error);
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
