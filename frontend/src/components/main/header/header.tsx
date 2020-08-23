import React from "react";
import "./header.scss";
import { AuthContext } from "../../../context/auth";
import { UserPanel } from "./userpanel";

export function Header() { 
    const signInButton = () => {
        return (<a href="http://localhost:8080/auth/google">
                    <img src="btn_google_signin.png" alt=""/>
                </a>);
    }

    return (
        <AuthContext.Consumer>
            {context => (
                <nav className="navbar navbar-expand-lg navbar-light bg-light">
                    <a className="navbar-brand" href="/">Go NATM</a>
                    <ul className="navbar-nav mr-auto"></ul>
                    {context.currentUser != null ? <UserPanel/> : signInButton()}
                </nav>
            )}
        </AuthContext.Consumer>
    )
};
