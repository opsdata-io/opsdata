import React, { useState } from 'react';
import { Link, useLocation } from 'react-router-dom';

const SideMenu = () => {
    const location = useLocation();
    const [customersMenuOpen, setCustomersMenuOpen] = useState(false);
    const [usersMenuOpen, setUsersMenuOpen] = useState(false);
    const [fileManagementMenuOpen, setFileManagementMenuOpen] = useState(false);

    const toggleMenu = (menuName) => {
        switch (menuName) {
            case 'customers':
                setCustomersMenuOpen(!customersMenuOpen);
                setUsersMenuOpen(false);
                setFileManagementMenuOpen(false);
                break;
            case 'users':
                setUsersMenuOpen(!usersMenuOpen);
                setCustomersMenuOpen(false);
                setFileManagementMenuOpen(false);
                break;
            case 'fileManagement':
                setFileManagementMenuOpen(!fileManagementMenuOpen);
                setCustomersMenuOpen(false);
                setUsersMenuOpen(false);
                break;
            default:
                setCustomersMenuOpen(false);
                setUsersMenuOpen(false);
                setFileManagementMenuOpen(false);
                break;
        }
    };

    return (
        <div className="bg-light border-right side-menu">
            <div className="sidebar-heading">Menu</div>
            <ul className="list-group list-group-flush">
                <li className={`list-group-item ${location.pathname === '/dashboard' ? 'active' : ''}`}>
                    <Link to="/dashboard" className="nav-link">Dashboard</Link>
                </li>
                <li className={`list-group-item ${location.pathname.includes('/customers') ? 'active' : ''}`} onClick={() => toggleMenu('customers')}>
                    <span className="menu-item">
                        Customers
                        {customersMenuOpen ? <i className="fa fa-caret-up float-right"></i> : <i className="fa fa-caret-down float-right"></i>}
                    </span>
                    {customersMenuOpen && (
                        <ul className="submenu list-group list-group-flush">
                            <li className={`list-group-item ${location.pathname === '/customers/add' ? 'active' : ''}`}>
                                <Link to="/customers/add" className="nav-link">Add</Link>
                            </li>
                            <li className={`list-group-item ${location.pathname === '/customers/search' ? 'active' : ''}`}>
                                <Link to="/customers/search" className="nav-link">Search</Link>
                            </li>
                        </ul>
                    )}
                </li>
                <li className={`list-group-item ${location.pathname.includes('/users') ? 'active' : ''}`} onClick={() => toggleMenu('users')}>
                    <span className="menu-item">
                        Users
                        {usersMenuOpen ? <i className="fa fa-caret-up float-right"></i> : <i className="fa fa-caret-down float-right"></i>}
                    </span>
                    {usersMenuOpen && (
                        <ul className="submenu list-group list-group-flush">
                            <li className={`list-group-item ${location.pathname === '/users/add' ? 'active' : ''}`}>
                                <Link to="/users/add" className="nav-link">Add User</Link>
                            </li>
                            <li className={`list-group-item ${location.pathname === '/users/search' ? 'active' : ''}`}>
                                <Link to="/users/search" className="nav-link">Search Users</Link>
                            </li>
                            <li className={`list-group-item ${location.pathname === '/users/manage' ? 'active' : ''}`}>
                                <Link to="/users/manage" className="nav-link">Manage Users</Link>
                            </li>
                            <li className={`list-group-item ${location.pathname === '/users/permissions' ? 'active' : ''}`}>
                                <Link to="/users/permissions" className="nav-link">User Permissions</Link>
                            </li>
                        </ul>
                    )}
                </li>
                <li className={`list-group-item ${location.pathname.includes('/file-management') ? 'active' : ''}`} onClick={() => toggleMenu('fileManagement')}>
                    <span className="menu-item">
                        File Management
                        {fileManagementMenuOpen ? <i className="fa fa-caret-up float-right"></i> : <i className="fa fa-caret-down float-right"></i>}
                    </span>
                    {fileManagementMenuOpen && (
                        <ul className="submenu list-group list-group-flush">
                            <li className={`list-group-item ${location.pathname === '/file-management/request' ? 'active' : ''}`}>
                                <Link to="/file-management/request" className="nav-link">Request a File</Link>
                            </li>
                            <li className={`list-group-item ${location.pathname === '/file-management/uploaded' ? 'active' : ''}`}>
                                <Link to="/file-management/uploaded" className="nav-link">Uploaded Files</Link>
                            </li>
                        </ul>
                    )}
                </li>
            </ul>
        </div>
    );
};

export default SideMenu;
