import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom'; // Import useNavigate instead of useHistory
import { getToken } from '../utils/jwt'; // Import getToken from jwt utility

const AddUserPage = ({ token }) => {
    const navigate = useNavigate(); // useNavigate hook for navigation
    const [userData, setUserData] = useState({
        username: '',
        email: '',
        role: '',
        isActive: true,
    });

    const handleChange = (e) => {
        const { name, value } = e.target;
        setUserData({ ...userData, [name]: value });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        const jwtToken = token || getToken(); // Retrieve the JWT token using the utility function
        try {
            const response = await fetch(`/api/users`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    Authorization: `Bearer ${jwtToken}`,
                },
                body: JSON.stringify(userData),
            });
            if (response.ok) {
                const newUser = await response.json();
                console.log('New user:', newUser); // Handle success as needed
                // Optionally, redirect to user list page or show a success message
                navigate('/users'); // Redirect to users page after successful creation using navigate
            } else {
                console.error('Failed to create user:', response.statusText);
            }
        } catch (error) {
            console.error('Error creating user:', error); // Handle error appropriately
        }
    };

    return (
        <div>
            <h2>User Management</h2>
            <h3>Add User</h3>
            <form onSubmit={handleSubmit}>
                <label>Username:</label>
                <input type="text" name="username" value={userData.username} onChange={handleChange} required />
                <br />
                <label>Email:</label>
                <input type="email" name="email" value={userData.email} onChange={handleChange} required />
                <br />
                <label>Role:</label>
                <input type="text" name="role" value={userData.role} onChange={handleChange} required />
                <br />
                <label>Active:</label>
                <input type="checkbox" name="isActive" checked={userData.isActive} onChange={() => setUserData({ ...userData, isActive: !userData.isActive })} />
                <br />
                <button type="submit">Submit</button>
            </form>
        </div>
    );
};

export default AddUserPage;
