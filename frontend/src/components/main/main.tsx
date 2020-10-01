import React, { Component } from "react";
import "./main.scss";
import { Header } from "./header"
import { AuthContext } from '../../context/auth';
import { LoginPage } from "./loginpage";
import { Routes } from "./routes";
import { ServiceContext, defaultServiceContext } from "../../context/service";
import { User } from "../../models/user";
import { IAuthService } from "../../services/interface/iauthservice";
import { AuthService } from "../../services/implementation/authservice";


interface IMainState {
    user: User|null
}

class Main extends Component<any, IMainState> {
    authService: IAuthService;

    constructor(props: any) {
        super(props);

        this.authService = new AuthService();

        this.state = {
            user: null
        };
    }

    async tryAuthenticate(): Promise<User|null> {
        const user = await this.authService.tryAuthenticate();

        this.setState({
            user: user
        });

        return user;
    }
    
    logout() {
        return async () => {
            const success = await this.authService.logout();
            this.setState({
                user: null
            });

            return success;
        };
    }

    render() {
        return (<AuthContext.Provider value={{
                    currentUser: this.state.user,
                    logout: this.logout()
                }}>
                    <ServiceContext.Provider value={defaultServiceContext()} >
                        <Header/>
                        <div className="main-body">
                            {this.state.user != null ? <Routes currentUser={this.state.user} /> : <LoginPage/>}
                        </div>
                    </ServiceContext.Provider>
                </AuthContext.Provider>);
    }

    async componentDidMount() {
        await this.tryAuthenticate();
    }
}

export default Main;