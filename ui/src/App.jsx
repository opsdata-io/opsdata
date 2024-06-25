// App.jsx

import React, { useState } from 'react';
import { BrowserRouter as Router, Route, Switch, Redirect } from 'react-router-dom';
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
        localStorage.setItem('token', token); // Save token to local storage
        setToken(token); // Set token in component state
    };

    return (
        <Router>
            <div className="app">
                <Banner />
                <div className="main-content">
                    <SideMenu />
                    <Switch>
                        <Route path="/login">
                            <LoginPage setToken={handleSetToken} />
                        </Route>
                        <Route path="/dashboard">
                            {token ? <Dashboard token={token} /> : <Redirect to="/login" />}
                        </Route>
                        <Route path="/customers/add">
                            {token ? <AddCustomerPage token={token} /> : <Redirect to="/login" />}
                        </Route>
                        <Route path="/customers/search">
                            {token ? <SearchCustomerPage token={token} /> : <Redirect to="/login" />}
                        </Route>
                        <Route path="/customer/:id">
                            {token ? <CustomerDetailsPage token={token} /> : <Redirect to="/login" />}
                        </Route>
                        <Route path="/upload/:link">
                            {token ? <UploadPage token={token} /> : <Redirect to="/login" />}
                        </Route>
                        <Route path="/confirmation" component={ConfirmationPage} />
                        <Route path="/downloads">
                            {token ? <DownloadPage token={token} /> : <Redirect to="/login" />}
                        </Route>
                        <Route path="/users/add">
                            {token ? <AddUserPage token={token} /> : <Redirect to="/login" />}
                        </Route>
                        <Route exact path="/">
                            <Redirect to={token ? "/dashboard" : "/login"} />
                        </Route>
                    </Switch>
                </div>
            </div>
        </Router>
    );
}

export default App;
