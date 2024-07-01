// components/Banner.jsx

import React from 'react';
import '../styles.css'; // Assuming styles.css imports global styles

const Banner = () => {
    const greenColor = '#009736'; // Define the green color

    return (
        <div className="banner" style={{ backgroundColor: 'white' }}>
            <div className="logo">
                <img src="https://cdn.opsdata.io/logos/small-logo.png" alt="OpsData" />
            </div>
            <div className="banner-text" style={{ fontFamily: 'Neuropol X Free, sans-serif' }}>
                <h1>
                    <span style={{ color: greenColor }}>Ops</span>
                    <span style={{ color: 'black' }}>Data</span>
                </h1>
            </div>
        </div>
    );
};

export default Banner;
