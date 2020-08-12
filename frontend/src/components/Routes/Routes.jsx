import React from "react";
import "./Routes.scss";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom"
import ProjectList from "components/ProjectList";
import ProjectDetail from "components/ProjectDetail";

const cols = [
    {'name': "Column 1"},
    {'name': "Column 2"},
    {'name': "Column 3"},
    {'name': "Column 4"}
];

export default function Routes() {
    return (
        <Router>
            <Switch>
                <Route exact path="/">
                    <ProjectList columns={cols}/>
                </Route>
                <Route exact path="/project/:id" children={<ProjectDetail/>}/>
            </Switch>
        </Router>
    );
};

