import React, { useState, useEffect, useRef } from "react";
import "./projectlist.scss";
import { Link } from "react-router-dom";
import { ProjectGrid } from "../../../../models/project";
import { FakeProjectService } from "../../../../services/implementation/fake-project-service";
import { ModalForm } from "../../modalform";
import { Form, Container, Table, Row, Popover, Card } from "react-bootstrap";
import { Plus, Trash } from "react-bootstrap-icons";
import moment from "moment";

export const ProjectList = () => {
    const [loading, setLoading] = useState(true);
    const [projects, setProjects] = useState<ProjectGrid[]>([]);
    const [showCreate, setShowCreate] = useState(false)
    const nameInput = useRef(null);
    const projectService = new FakeProjectService()

    const renderProjects = (projects: ProjectGrid[]) => {
        if (projects.length == 0) {
            return (<div>No Projects Found</div>)
        }

        const projectLinks = projects.map(p => (
                <tr>
                    <td><Link key={p.id} to={"/project/" + p.id}>{p.name} ({p.abbreviation})</Link></td>
                    <td>{p.associationType}</td>
                    <td>{moment.unix(p.lastUpdated).toString()}</td>
                    <td align="right"><button className="btn m-1 p-1" onClick={() => handleArchiveProject(p.id)}><Trash/></button></td>
                </tr>
        ))
        return (
            <Table className="project-table">
                <thead>
                    <th>Project</th>
                    <th>Association</th>
                    <th colSpan={2}>Last Update</th>
                </thead>
                <tbody>
                    {projectLinks}
                </tbody>
            </Table>);
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

        if (!success) {
            //TODO show error
        }

        setLoading(true)

        const projects = await projectService.getProjectList()
        setProjects(projects);
        setLoading(false);
    };

    const handleArchiveProject = async (projectID: string) => {
        if (!window.confirm("Are you sure you would like to archive this project?")) {
            return;
        }

        const success = await projectService.archiveProject(projectID);
        if (!success) {
            //TODO show error
        }

        setLoading(true)

        const projects = await projectService.getProjectList()
        setProjects(projects);
        setLoading(false);
    }

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
        <Card>
            <Card.Body>
                <Card.Title>Projects<button className="btn btn-outline-primary btn-sm mr-3 float-right" onClick={() => setShowCreate(true)}>Create New</button>
            </Card.Title>
                
            {renderProjects(projects)}
            <ModalForm
                focusElement={nameInput}
                formContent={formContent}
                formElementId="project-create-form"
                onSubmit={handleCreateProject}
                setShow={setShowCreate}
                show={showCreate}
                title="Create New Project"
            />
            </Card.Body>
        </Card>
    </Container>)
};