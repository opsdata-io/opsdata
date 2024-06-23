// components/Banner.jsx

import React from 'react';
import logo from '../assets/avatar-centered.png';


const Banner = () => {
    return (
        <div className="banner">
            <div className="logo">
                <img src={logo} alt="OpsData" />
            </div>
            <div className="banner-text">
                <h1>OpsData</h1>
            </div>
        </div>
    );
};

export default Banner;