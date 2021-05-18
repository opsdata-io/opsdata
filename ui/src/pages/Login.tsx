import React, {SyntheticEvent, useState} from 'react';
import {Redirect} from "react-router-dom";

const Login = (props: { setName: (name: string) => void}) => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [redirect, setRedirect] = useState(false);
    const [failed, setFailed] = useState(false);

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();
        const response = await fetch('/api/login', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            credentials: 'include',
            body: JSON.stringify({
                email,
                password
            })
        });

        function refreshPage() {
          window.location.reload(false);
        }

        const content = await response.json();
        
        if(response.status === 200){
          setRedirect(true);
          props.setName(content.name);
          refreshPage();
        }else{
          setFailed(true)
          setRedirect(false);
        }
    }

    let message;

    if (failed) {
      message = (
        <p className="alert alert-danger" style={{ textAlign: "center" }}>Login failed, please try again</p>
      )
    }
    else
    {
      message = (
        <p className="alert alert-primary" style={{ textAlign: "center" }}>Please sign in</p>
      )
    }
    if (redirect) {
        return <Redirect to="/"/>;
    }

    return (
        <form onSubmit={submit}>
            {message}
            
            <input type="email" className="form-control" placeholder="Email address" required
                   onChange={e => setEmail(e.target.value)}
            />

            <input type="password" className="form-control" placeholder="Password" required
                   onChange={e => setPassword(e.target.value)}
            />

            <button className="w-100 btn btn-lg btn-primary" type="submit">Sign in</button>
        </form>
    );
};

export default Login;