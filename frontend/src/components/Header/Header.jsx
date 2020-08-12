import React from "react";
import "./Header.scss";
import { AuthContext } from "context/auth";
import UserPanel from "components/UserPanel";

export default function Header() { 
    return (
        <AuthContext.Consumer>
            {context => (
                <nav class="navbar navbar-expand-lg navbar-light bg-light">
                    <a class="navbar-brand" href="#">Go NATM</a>
                    <ul className="navbar-nav mr-auto"></ul>
                    {context.authenticated ? <UserPanel/> : SignInButton()}
                </nav>
            )}
        
        </AuthContext.Consumer>
    
    )
};

function SignInButton() {
    return (<a href="http://localhost:8080/auth/google">
        <img src="btn_google_signin.png" alt=""/>
    </a>);
}