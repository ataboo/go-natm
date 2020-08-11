import React, { Component } from "react";
import "./Main.scss";
import Header from "../Header/Header"
import { AuthContext } from 'context/auth';


class Main extends Component {
    static contextType = AuthContext;

    render() {
        let body;
        if (this.context.authenticated) {
            body = (
                <div>
                    <div>User: {this.context.username}</div>
                    <button onClick={this.context.logout}>Logout</button>
                </div>
            )
        } else {
            body = (
                <div>
                    <p>Please sign-in to continue.</p>
                    <a href="http://localhost:8080/auth/google">
                        <img src="btn_google_signin.png" alt=""/>
                    </a>
                </div>
            ) 
        }

        return (
            <div>
                <Header/>
                <div className="container">
                    {body}
                </div>
            </div>
        );
    }

    componentDidMount() {
        this.context.tryAuthenticate();
    }
}

export default Main;