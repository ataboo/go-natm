import React, { Component } from "react";
import "./UnauthenticatedApp.scss";
import Header from "../Header/Header"

class UnauthenticatedApp extends Component {
    render() {
        return (
            <div>
                <Header />
                <div>Not Authenticated</div>
            </div>
        );
    }
}

export default UnauthenticatedApp;