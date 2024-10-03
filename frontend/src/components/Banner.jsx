import React from 'react';
import '../styles.css';

const Banner = () => {
    return (
        <header className="banner">
            <div className="logo">
                <img src="https://cdn.opsdata.io/logos/small-logo.png" alt="OpsData Logo" />
            </div>
            <div className="banner-text">
                <h1>
                    <span className="ops-color">Ops</span>
                    <span className="data-color">Data</span>
                </h1>
            </div>
        </header>
    );
};

export default Banner;
