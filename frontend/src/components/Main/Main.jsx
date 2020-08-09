import React, { Component } from "react";
import "./Main.scss";
import Header from "../Header/Header"
import { AuthContext } from 'context/auth';
import {googleLoginPath} from 'api'

class Main extends Component {
    onCode(code, params) {
        console.log("Got code: " + code);
        console.log("Got params: " + params);
    }

    onClose() {
        console.log("closed");
    }

    render() {
        return (
            <React.Fragment>
                <Header />
                <a href={googleLoginPath}>Login with Google</a>
                <AuthContext.Consumer>
                    {(value) => (<div>User: {value}</div>)}
                </AuthContext.Consumer>
            </React.Fragment>
        );
    }
}

export default Main;