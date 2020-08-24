import React, { useState, useEffect } from "react";
import "./projectlist.scss";
import { Link } from "react-router-dom";
import { Project } from "../../../../models/project";
import { FakeProjectService } from "../../../../services/implementation/fake-project-service";

export const ProjectList = () => {
    const [loading, setLoading] = useState(true);
    const [projects, setProjects] = useState<Project[]>([]);

    const renderProjects = (projects: Project[]) => {
        const projectLinks = projects.map(p => (<Link key={p.id} to={"/project/" + p.id}>Project: {p.name}</Link>))

        return (<ul>{projectLinks}</ul>);
    };

    useEffect(() => {
        const fetchProjects = async() => {
            const projectService = new FakeProjectService();
            const projects = await projectService.getProjectList()
            setProjects(projects);
            setLoading(false);
        }

        fetchProjects();
    }, [])

    if (loading) {
        return (<div>Loading...</div>);
    }

    return renderProjects(projects);
};