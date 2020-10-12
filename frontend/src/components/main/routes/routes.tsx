import React from "react";
import "./routes.scss";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom"
import {ProjectList} from "./projectlist";
import {Project} from "./project";
import { User } from "../../../models/user";

type RouteProps = {
    currentUser: User
}

export function Routes({currentUser}: RouteProps) {
    return (
        <Router>
            <Switch>
                <Route exact path="/">
                    <ProjectList currentUser={currentUser}/>
                </Route>
                <Route exact path="/project/:id" render={(props) => {
                    return (<Project id={props.match.params.id} currentUser={currentUser}/>)
                }}/>
            </Switch>
        </Router>
    );
};

