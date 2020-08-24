import React from "react";
import "./header.scss";
import { AuthContext } from "../../../context/auth";
import { UserPanel } from "./userpanel";

export function Header() { 
    const signInButton = () => {
        return (<a href="http://localhost:8080/auth/google">
                    <img src="/btn_google_signin.png" alt="This is where button is"/>
                </a>);
    }

    return (
        <AuthContext.Consumer>
            {context => {
                const hasUser = context.currentUser !== null;
     
                return (
                <nav className="navbar navbar-expand-lg navbar-light bg-light">
                    <a className="navbar-brand" href="/">Go NATM</a>
                    <ul className="navbar-nav mr-auto"></ul>
                    {(hasUser ? <UserPanel/> : signInButton())}
                </nav>
            )
            }}
                
        </AuthContext.Consumer>
    )
};
