import React from "react";
import "./userpanel.scss";
import { AuthContext } from "../../../../context/auth";

export function UserPanel() { 
    return (
        <AuthContext.Consumer>
            {context => (
                <ul className="navbar-nav">
                    <li className="align-items-center">
                        <div>{context.currentUser?.name ?? "No User"}</div>
                    </li>
                    <li className="align-items-center ml-3">
                        <button className="btn btn-secondary" onClick={context.logout}>Logout</button>
                    </li>
                </ul>
            )}
        </AuthContext.Consumer>
        
    );
};
