import React, { Component } from "react";
import "./Main.scss";
import Header from "../Header/Header"
import { AuthContext } from 'context/auth';

class Main extends Component {
    render() {
        return (
            <React.Fragment>
                <Header />
                <AuthContext.Consumer>
                    {(value) => (<div>User: {value}</div>)}
                </AuthContext.Consumer>
            </React.Fragment>
        );
    }
}

export default Main;