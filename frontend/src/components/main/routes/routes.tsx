import React from "react";
import "./routes.scss";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom"
import {ProjectList} from "./projectlist";
import {Project} from "./project";
import { User } from "../../../models/user";
import { IMainActions } from "../imainactions";

type RouteProps = {
    mainActions: IMainActions
}

export function Routes({mainActions}: RouteProps) {
    return (
        <Router>
            <Switch>
                <Route exact path="/">
                    <ProjectList currentUser={mainActions.currentUser!}/>
                </Route>
                <Route exact path="/project/:id" render={(props) => {
                    return (<Project id={props.match.params.id} mainActions={mainActions}/>)
                }}/>
            </Switch>
        </Router>
    );
};

