import React from "react";
import "./projectlist.scss";
import {Link} from "react-router-dom";
import { ServiceContext } from "../../../../context/service";

export const ProjectList = () => (
    <ServiceContext.Consumer>
        {async (context) => {
            const projects = await context.projectService.getProjectList()
            const projectLinks = projects.map(p => (<Link to={"/project/" + p.id}>Project: {p.name}</Link>))   
            
            return (<ul>{projectLinks}</ul>)
        }}
    </ServiceContext.Consumer>
);