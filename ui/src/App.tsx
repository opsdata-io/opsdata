import React, {useEffect, useState} from 'react';
import './App.css';
import Login from "./pages/Login";
import Nav from "./components/Nav";
import {BrowserRouter, Route} from "react-router-dom";
import Home from "./pages/Home";
import Register from "./pages/Register";

var apiUrl = process.env.API_URL;
var currentDomain = window.location.hostname;

if(!apiUrl)
{
    apiUrl = 'http://localhost:8000';
}
else
{
    apiUrl = `https://${currentDomain}`;
}

function App() {
    const [name, setName] = useState('');

    useEffect(() => {
        (
            async () => {
                var url = `${apiUrl}/api/user`;
                const response = await fetch(url, {
                    headers: {'Content-Type': 'application/json'},
                    credentials: 'include',
                });

                const content = await response.json();

                setName(content.name);
            }
        )();
    });


    return (
        <div className="App">
            <BrowserRouter>
                <Nav name={name} setName={setName}/>

                <main className="form-signin">
                    <Route path="/" exact component={() => <Home name={name}/>}/>
                    <Route path="/login" component={() => <Login setName={setName}/>}/>
                    <Route path="/register" component={Register}/>
                </main>
            </BrowserRouter>
        </div>
    );
}

export default App;