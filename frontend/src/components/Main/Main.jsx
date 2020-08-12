import React, { Component } from "react";
import "./Main.scss";
import Header from "../Header/Header"
import { AuthContext } from 'context/auth';
import LoginPage from "components/LoginPage/LoginPage";
import { Container } from "react-bootstrap";
import Routes from "components/Routes/Routes";


class Main extends Component {
    static contextType = AuthContext;

    render() {
        return (
            <div>
                <Header/>
                <Container>
                    {this.context.authenticated ? <Routes/> : <LoginPage/>}
                </Container>
            </div>
        );
    }

    componentDidMount() {
        this.context.tryAuthenticate();
    }
}

export default Main;