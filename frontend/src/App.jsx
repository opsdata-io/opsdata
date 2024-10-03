import React, { lazy, Suspense } from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate, useLocation } from 'react-router-dom';
import { AuthProvider, useAuth } from './context/AuthContext';
import SideMenu from './components/SideMenu';
import Banner from './components/Banner';
import './App.css';

const LoginPage = lazy(() => import('./pages/LoginPage'));
const Dashboard = lazy(() => import('./pages/Dashboard'));
const UploadPage = lazy(() => import('./pages/UploadPage'));
const ConfirmationPage = lazy(() => import('./pages/ConfirmationPage'));
const DownloadPage = lazy(() => import('./pages/DownloadPage'));
const AddCustomerPage = lazy(() => import('./pages/AddCustomerPage'));
const SearchCustomerPage = lazy(() => import('./pages/SearchCustomerPage'));
const CustomerDetailsPage = lazy(() => import('./pages/CustomerDetailsPage'));
const AddUserPage = lazy(() => import('./pages/AddUserPage'));

const AppContent = () => {
    const { token } = useAuth();
    const location = useLocation();

    const showSideMenu = location.pathname !== '/login'; // Hide side menu on login page

    return (
        <div className="app">
            <Banner />
            <div className={`main-content ${showSideMenu ? '' : 'full-width'}`}>
                {showSideMenu && <SideMenu />}
                <Suspense fallback={<div>Loading...</div>}>
                    <Routes>
                        <Route path="/login" element={<LoginPage />} />
                        <Route path="/dashboard" element={token ? <Dashboard /> : <Navigate to="/login" />} />
                        <Route path="/customers/add" element={token ? <AddCustomerPage /> : <Navigate to="/login" />} />
                        <Route path="/customers/search" element={token ? <SearchCustomerPage /> : <Navigate to="/login" />} />
                        <Route path="/customer/:id" element={token ? <CustomerDetailsPage /> : <Navigate to="/login" />} />
                        <Route path="/upload/:link" element={token ? <UploadPage /> : <Navigate to="/login" />} />
                        <Route path="/confirmation" element={<ConfirmationPage />} />
                        <Route path="/downloads" element={token ? <DownloadPage /> : <Navigate to="/login" />} />
                        <Route path="/users/add" element={token ? <AddUserPage /> : <Navigate to="/login" />} />
                        <Route exact path="/" element={<Navigate to={token ? "/dashboard" : "/login"} />} />
                    </Routes>
                </Suspense>
            </div>
        </div>
    );
};

const App = () => (
    <Router>
        <AuthProvider>
            <AppContent />
        </AuthProvider>
    </Router>
);

export default App;
