import React, { useState } from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import LoginPage from './pages/LoginPage';
import Dashboard from './pages/Dashboard';
import UploadPage from './pages/UploadPage';
import ConfirmationPage from './pages/ConfirmationPage';
import DownloadPage from './pages/DownloadPage';
import SideMenu from './components/SideMenu';
import Banner from './components/Banner';
import AddCustomerPage from './pages/AddCustomerPage';
import SearchCustomerPage from './pages/SearchCustomerPage';
import CustomerDetailsPage from './pages/CustomerDetailsPage';
import AddUserPage from './pages/AddUserPage';
import './App.css';

function App() {
    const [token, setToken] = useState(localStorage.getItem('token'));

    const handleSetToken = (token) => {
        localStorage.setItem('token', token);
        setToken(token);
    };

    return (
        <Router>
            <div className="app">
                <Banner />
                <div className="main-content">
                    <SideMenu />
                    <Routes>
                        <Route path="/login" element={<LoginPage setToken={handleSetToken} />} />
                        <Route path="/dashboard" element={token ? <Dashboard token={token} /> : <Navigate to="/login" />} />
                        <Route path="/customers/add" element={token ? <AddCustomerPage token={token} /> : <Navigate to="/login" />} />
                        <Route path="/customers/search" element={token ? <SearchCustomerPage token={token} /> : <Navigate to="/login" />} />
                        <Route path="/customer/:id" element={token ? <CustomerDetailsPage token={token} /> : <Navigate to="/login" />} />
                        <Route path="/upload/:link" element={token ? <UploadPage token={token} /> : <Navigate to="/login" />} />
                        <Route path="/confirmation" element={<ConfirmationPage />} />
                        <Route path="/downloads" element={token ? <DownloadPage token={token} /> : <Navigate to="/login" />} />
                        <Route path="/users/add" element={token ? <AddUserPage token={token} /> : <Navigate to="/login" />} />
                        <Route exact path="/" element={<Navigate to={token ? "/dashboard" : "/login"} />} />
                    </Routes>
                </div>
            </div>
        </Router>
    );
}

export default App;
