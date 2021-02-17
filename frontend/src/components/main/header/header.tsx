import React from "react";
import "./header.scss";
import { UserPanel } from "./userpanel";
import { Nav, Navbar } from "react-bootstrap";
import { User } from "../../../models/user";
import { IMainActions } from "../imainactions";

type HeaderProps = {
    mainActions: IMainActions
}

export const Header = ({mainActions} : HeaderProps) => { 
    const signInButton = () => {
        return (<a href="http://localhost:8080/auth/google">
                    <img src="/btn_google_signin.png" alt="This is where button is"/>
                </a>);
    }

    const hasUser = mainActions.currentUser !== null;

    return (<Navbar variant="light">
                <Nav className="mr-auto">
                    <Navbar.Brand href="/">Go NATM</Navbar.Brand>
                    <Nav.Link href="/">Projects</Nav.Link>
                </Nav>
                {(hasUser ? <UserPanel mainActions={mainActions} /> : signInButton())}
            </Navbar>)
};
