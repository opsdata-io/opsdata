// components/SideMenu.jsx

import React, { useState } from 'react';
import { Link, useLocation } from 'react-router-dom';

const SideMenu = () => {
    const location = useLocation();
    const [menuState, setMenuState] = useState({
        customers: false,
        users: false,
        fileManagement: false
    });

    const toggleMenu = (menuName) => {
        setMenuState((prevState) => ({
            ...prevState,
            [menuName]: !prevState[menuName]
        }));
    };

    const renderMenu = (menuName, displayName, submenuItems) => (
        <li className={`list-group-item ${location.pathname.includes(`/${menuName}`) ? 'active' : ''}`} onClick={() => toggleMenu(menuName)}>
            <span className="menu-item">
                {displayName}
                {menuState[menuName] ? <i className="fa fa-caret-up float-right"></i> : <i className="fa fa-caret-down float-right"></i>}
            </span>
            {menuState[menuName] && (
                <ul className="submenu list-group list-group-flush">
                    {submenuItems.map((item) => (
                        <li key={item.path} className={`list-group-item ${location.pathname === item.path ? 'active' : ''}`}>
                            <Link to={item.path} className="nav-link">{item.label}</Link>
                        </li>
                    ))}
                </ul>
            )}
        </li>
    );

    return (
        <div className="bg-light border-right side-menu">
            <div className="sidebar-heading">Menu</div>
            <ul className="list-group list-group-flush">
                <li className={`list-group-item ${location.pathname === '/dashboard' ? 'active' : ''}`}>
                    <Link to="/dashboard" className="nav-link">Dashboard</Link>
                </li>
                {renderMenu('customers', 'Customers', [
                    { path: '/customers/add', label: 'Add' },
                    { path: '/customers/search', label: 'Search' }
                ])}
                {renderMenu('users', 'Users', [
                    { path: '/users/add', label: 'Add User' },
                    { path: '/users/search', label: 'Search Users' },
                    { path: '/users/manage', label: 'Manage Users' },
                    { path: '/users/permissions', label: 'User Permissions' }
                ])}
                {renderMenu('fileManagement', 'File Management', [
                    { path: '/file-management/request', label: 'Request a File' },
                    { path: '/file-management/uploaded', label: 'Uploaded Files' }
                ])}
            </ul>
        </div>
    );
};

export default SideMenu;
