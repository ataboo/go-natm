import React, { useState, useEffect, useRef } from "react";
import "./projectlist.scss";
import { Link } from "react-router-dom";
import { ProjectDetails } from "../../../../models/project";
import { FakeProjectService } from "../../../../services/implementation/fake-project-service";
import { ModalForm } from "../../modalform";
import { Form, Container } from "react-bootstrap";
import { Plus } from "react-bootstrap-icons";

export const ProjectList = () => {
    const [loading, setLoading] = useState(true);
    const [projects, setProjects] = useState<ProjectDetails[]>([]);
    const [showCreate, setShowCreate] = useState(false)
    const nameInput = useRef(null);
    const projectService = new FakeProjectService()

    const renderProjects = (projects: ProjectDetails[]) => {
        if (projects.length == 0) {
            return (<div>No Projects Found</div>)
        }

        const projectLinks = projects.map(p => (<li><Link key={p.id} to={"/project/" + p.id}>{p.name}</Link></li>))
        return (<ul>{projectLinks}</ul>);
    };

    const formContent = (<>
            <Form.Group controlId="title">
                <Form.Label>Project Name</Form.Label>
                <Form.Control autoFocus={true} type="text" name="name" required={true} ref={nameInput}></Form.Control>
            </Form.Group>

            <Form.Group controlId="abbreviation">
                <Form.Label>Abbreviation</Form.Label>
                <Form.Control type="text" name="abbreviation" required={true}></Form.Control>
            </Form.Group>

            <Form.Group controlId="description">
                <Form.Label>Description</Form.Label>
                <Form.Control as="textarea" rows={2} name="description"></Form.Control>
            </Form.Group>
        </>);

    const handleCreateProject = async (formData: FormData) => {
        const success = await projectService.createProject({
            abbreviation: formData.get("abbreviation") as string,
            description: formData.get("description") as string,
            name: formData.get("name") as string,
        });

        setLoading(true)

        const projects = await projectService.getProjectList()
        setProjects(projects);
        setLoading(false);
    };

    useEffect(() => {
        const fetchProjects = async() => {
            const projects = await projectService.getProjectList()
            setProjects(projects);
            setLoading(false);
        }

        fetchProjects();
    }, [])

    if (loading) {
        return (<div>Loading...</div>);
    }
    
    return (<Container>
        {renderProjects(projects)}
        <br/>
        <button className="btn btn-primary m-2 p-1" onClick={() => setShowCreate(true)}><Plus size={24} /> <span className="mr-1">Create Project</span></button>
        <ModalForm
            focusElement={nameInput}
            formContent={formContent}
            formElementId="project-create-form"
            onSubmit={handleCreateProject}
            setShow={setShowCreate}
            show={showCreate}
            title="Create New Project"
        />
    </Container>)
};