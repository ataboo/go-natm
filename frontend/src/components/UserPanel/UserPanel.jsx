import React from "react";
import "./UserPanel.scss";
import { AuthContext } from "context/auth";

export default function UserPanel() { 
    return (
        <AuthContext.Consumer>
            {context => (
                <ul className="navbar-nav">
                    <li className="align-items-center">
                        <div>{context.username}</div>
                    </li>
                    <li className="align-items-center ml-3">
                        <button className="btn btn-secondary" onClick={context.logout}>Logout</button>
                    </li>
                </ul>
            )}
        </AuthContext.Consumer>
        
    );
};
