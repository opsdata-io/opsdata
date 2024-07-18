import React, { useEffect, useState } from 'react';
import PropTypes from 'prop-types';
import './Notification.css';

const Notification = ({ message, type }) => {
    const [visible, setVisible] = useState(false);

    useEffect(() => {
        if (message) {
            setVisible(true);
            const timer = setTimeout(() => {
                setVisible(false);
            }, 3000); // Dismiss notification after 3 seconds
            return () => clearTimeout(timer);
        }
    }, [message]);

    if (!visible) return null;

    return (
        <div
            className={`notification ${type}`}
            role="alert"
            aria-live="assertive"
            aria-atomic="true"
        >
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
