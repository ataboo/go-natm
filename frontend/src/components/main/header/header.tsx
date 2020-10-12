import React from "react";
import "./header.scss";
import { AuthContext } from "../../../context/auth";
import { UserPanel } from "./userpanel";
import { Nav, Navbar } from "react-bootstrap";

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
     
                return (<Navbar variant="light">
                            <Nav className="mr-auto">
                                <Navbar.Brand href="/">Go NATM</Navbar.Brand>
                                <Nav.Link href="/">Projects</Nav.Link>
                            </Nav>
                            {(hasUser ? <UserPanel/> : signInButton())}
                        </Navbar>)
            }}
                
        </AuthContext.Consumer>
    )
};
