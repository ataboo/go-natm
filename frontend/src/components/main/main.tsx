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
import { TaskRead } from "../../models/task";
import { IMainActions } from "./imainactions";
import { IProjectService } from "../../services/interface/iproject-service";


interface IMainState {
    currentUser: User|null;
    activeTaskId?: string;
}

class Main extends Component<any, IMainState> {
    authService: IAuthService;
    static contextType = ServiceContext;

    constructor(props: any) {
        super(props);

        this.authService = new AuthService();

        this.state = {
            currentUser: null,
            activeTaskId: undefined
        };
    }

    async tryAuthenticate(): Promise<User|null> {
        const user = await this.authService.tryAuthenticate();

        this.setState(oldState => ({...oldState, currentUser: user}));

        return user;
    }
    
    logout() {
        return async () => {
            const success = await this.authService.logout();
            this.setState({
                currentUser: null
            });

            return success;
        };
    }

    startTimingTask() {
        return (id: string) => {
            const projectService: IProjectService = this.context.projectService;
            projectService.startLoggingWork(id)
              .then(success => {
                if (!success) {
                  console.error("failed to activate task");
                  return;
                }
                this.setState(oldState => ({
                    currentUser: oldState.currentUser,
                    activeTaskId: id
                }))
              })
          }
    }

    stopTimingTask() {
        return () => {
            const projectService: IProjectService = this.context.projectService;
            projectService.stopLoggingWork()
            .then(success => {
                if (!success) {
                console.error("failed to stop work");
                return;
                }
    
                this.setState(oldState => ({
                    currentUser: oldState.currentUser,
                    activeTaskId: undefined
                }))
            });
        }
    }

    refreshTimedTask() {
        console.warn("todo!");
    }

    render() {
        const mainActions: IMainActions = {
            startLoggingTask: this.startTimingTask(),
            stopLoggingTask: this.stopTimingTask(),
            currentUser: this.state.currentUser,
            activeTaskId: this.state.activeTaskId,
            logout: this.logout(),
            refreshTimedTask: this.refreshTimedTask
        }

        return (<>
                    <Header mainActions={mainActions}/>
                    <div className="main-background">
                        <div className="main-body">
                            {this.state.currentUser != null ? <Routes mainActions={mainActions} /> : <LoginPage/>}
                        </div>
                    </div>
                </>);
    }

    async componentDidMount() {
        await this.tryAuthenticate();
    }
}

export default Main;