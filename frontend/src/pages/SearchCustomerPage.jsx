import React, { useState } from 'react';
import { Link } from 'react-router-dom';

const SearchCustomerPage = ({ token }) => {
    const [searchQuery, setSearchQuery] = useState('');
    const [searchResults, setSearchResults] = useState([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState('');

    const handleSearch = async (e) => {
        e.preventDefault();  // Prevent the default form submission behavior
        setLoading(true);  // Set loading to true to indicate search is in progress
        setError('');  // Clear any existing errors
        try {
            const response = await fetch(`/v1/customers/search?q=${encodeURIComponent(searchQuery)}`, {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            });
            if (response.ok) {
                const data = await response.json();
                setSearchResults(data);
            } else {
                throw new Error('Failed to fetch search results');  // Throw an error if response is not OK
            }
        } catch (error) {
            console.error('Failed to fetch search results', error);
            setError('Failed to fetch search results. Please try again.');
        } finally {
            setLoading(false);  // Set loading to false after the fetch operation is done
        }
    };

    const handleChange = (e) => {
        setSearchQuery(e.target.value);
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
                    required
                    style={{ marginRight: '0.5rem' }}
                />
                <button type="submit" disabled={loading}>Search</button>
            </form>
            {loading && <p>Searching...</p>}
            {error && <div style={{ color: 'red' }}>{error}</div>}
            <div>
                {searchResults.length === 0 && !loading ? (
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
