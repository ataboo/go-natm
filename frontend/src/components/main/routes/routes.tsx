import React from "react";
import "./routes.scss";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom"
import {ProjectList} from "./projectlist";
import {Project} from "./project";


export function Routes() {
    return (
        <Router>
            <Switch>
                <Route exact path="/">
                    <ProjectList />
                </Route>
                <Route exact path="/project/:id" children={<Project />}/>
            </Switch>
        </Router>
    );
};

