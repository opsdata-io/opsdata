// components/Notification.jsx

import React from 'react';
import PropTypes from 'prop-types';
import './Notification.css'; // Import the CSS file for styling

const Notification = ({ message, type }) => {
    if (!message) return null;

    return (
        <div className={`notification ${type}`}>
            {message}
        </div>
    );
};

Notification.propTypes = {
    message: PropTypes.string,
    type: PropTypes.oneOf(['success', 'error', 'info']).isRequired
};

Notification.defaultProps = {
    type: 'info'
};

export default Notification;
