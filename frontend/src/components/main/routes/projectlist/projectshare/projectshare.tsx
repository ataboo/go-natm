import React, { createRef, RefObject, useEffect, useState } from 'react';
import { Button, ButtonGroup, Col, Form, Modal, Row } from 'react-bootstrap';
import { Check, Pencil, Trash, X } from 'react-bootstrap-icons';
import { AssociationType } from '../../../../../enums';
import { ProjectDetails, ProjectGrid } from '../../../../../models/project';
import { ProjectAssociationCreate, ProjectAssociationDelete, ProjectAssociationDetail, ProjectAssociationUpdate } from '../../../../../models/projectassociation';
import { User } from '../../../../../models/user';

type ProjectShareProps = {
    projectGridData: ProjectGrid,
    showShare: boolean,
    currentUser: User,
    loadProjectDetail: (projectId: string) => Promise<ProjectDetails>;
    onClose: () => void;
    createProjectAssociation(data: ProjectAssociationCreate): Promise<boolean>
    updateProjectAssociation(data: ProjectAssociationUpdate): Promise<boolean>
    deleteProjectAssociation(data: ProjectAssociationDelete): Promise<boolean>
}

type AssociationMapEntry = {
    editing: boolean,
    formRef: React.RefObject<HTMLFormElement> | undefined
};

export const ProjectShare = ({projectGridData, showShare, currentUser, loadProjectDetail, onClose, createProjectAssociation, updateProjectAssociation, deleteProjectAssociation}: ProjectShareProps) => {
    const [projectData, setProjectData] = useState<ProjectDetails | undefined>(undefined);
    const associationTypeOptions = Object.keys(AssociationType).filter(k => k !== AssociationType.Owner).map(k => <option value={k} key={k}>{k}</option>);
    const [associationMap, setAssociationMap] = useState(new Map<string, AssociationMapEntry>());
    const [newShareEmailError, setNewShareError] = useState<string | undefined>(undefined);

    const updateEntryEditing = (id: string, editing: boolean) => {
        const entry = associationMap.get(id)!;
        entry.editing = editing;
        setAssociationMap(new Map(associationMap.set(id, entry)));
    };

    const updateEntryRef = (id: string, ref: React.RefObject<HTMLFormElement>) => {
        const entry = associationMap.get(id)!;
        entry.formRef = ref;
        setAssociationMap(new Map(associationMap.set(id, entry)));
    }

    useEffect(() => {
        (async () => {
            const detail = await loadProjectDetail(projectGridData.id);
            setProjectData(detail);

            const newAssocMap = detail.associations.reduce((m, v) => {
                return m.set(v.id, {
                    editing: false,
                    formRef: undefined
                }
            )}, new Map<string, AssociationMapEntry>());

            setAssociationMap(newAssocMap);
        })();            
    }, [setAssociationMap, setProjectData, loadProjectDetail, projectGridData.id]);

    const reloadProjectData = async () => {
        setProjectData(undefined);
        const detail = await loadProjectDetail(projectGridData.id);

        const newAssocMap = detail.associations.reduce((m, v) => {
            return m.set(v.id, {
                editing: false,
                formRef: undefined
            }
        )}, new Map<string, AssociationMapEntry>());

        setAssociationMap(newAssocMap);
        setProjectData(detail);
    }
    
    const renderInteractiveRow = (association: ProjectAssociationDetail, ref: RefObject<HTMLFormElement>) => {
        return (
            <Form onSubmit={onExistingAssociationSubmit} ref={ref} key={association.id} className='mb-2'>    
                <Row>
                    <Col lg='5' sm='12'>
                        <div className='pt-2'>{association.email}</div>
                    </Col>
                    <Col lg='5' sm='6'>
                        <input type="hidden" name="id" value={association.id}/>
                        <Form.Control as="select" name="type" defaultValue={association.type}>
                            {associationTypeOptions}
                        </Form.Control>                
                    </Col>
                    <Col lg='2' sm='6'>
                        <ButtonGroup>
                            <Button variant='success' type='submit'><Check/></Button>
                            <Button variant='light' type='button' onClick={() => {
                                associationMap.get(association.id)!.formRef!.current!.reset();
                                updateEntryEditing(association.id, false);

                            }}><X/></Button>
                        </ButtonGroup>
                    </Col>
                </Row>
            </Form>);
    };

    const renderReadonlyRow = (association: ProjectAssociationDetail, canUpdate: boolean) => {
        return (
            <Row key={association.id} className='mb-2'>
                <Col lg='5' sm='12'>
                    <div className='pt-2'>{association.email}</div>
                </Col>
                <Col lg='5' sm='6'>    
                    <div className='pt-2'>{association.type}</div>            
                </Col>
                <Col lg='2' sm='6'>
                    {canUpdate ? (
                        <ButtonGroup>
                            <Button type="button" variant='light' onClick={() => updateEntryEditing(association.id, true)}><Pencil/></Button>
                            <Button type="button" variant='danger' onClick={() => onClickDeleteAssociation(association.id, association.email)}><Trash/></Button>
                        </ButtonGroup>) : ''}
                </Col>
            </Row>)
    };

    const onExistingAssociationSubmit = (e: React.FormEvent) => {
        e.preventDefault();

        const formData = new FormData(e.target as HTMLFormElement);
        const type = formData.get('type') as string as AssociationType;
        const id = formData.get('id') as string;

        updateProjectAssociation({
            id: id,
            type: type
        }).then(success => {
            reloadProjectData();
        });
    }

    const validateNewShareEmail = (e: React.FocusEvent<HTMLInputElement>) => {
        const element = e.target as HTMLInputElement
        const emailValue = element.value;
        const emailPattern = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/

        if (!emailPattern.test(emailValue)) {
            setNewShareError('Please enter a valid email.');
            return;
        }
        
        if (projectData!.associations.find(a => a.email === emailValue)) {
            setNewShareError('Please enter an email address that has not already been shared.');
            return;
        }
        
        setNewShareError(undefined);
    }

    const onNewAssociationSubmit = (e: React.FormEvent) => {
        e.preventDefault();

        if (newShareEmailError !== undefined) {
            return;
        }

        const formData = new FormData(e.target as HTMLFormElement);
        const type = formData.get('type') as string as AssociationType;
        const email = formData.get('email') as string;

        createProjectAssociation({
            projectId: projectGridData.id,
            email: email,
            type: type
        }).then(success => {
            reloadProjectData();
        });
    }

    const onClickDeleteAssociation = (id: string, email: string) => {
        if (!window.confirm(`Are you sure you want to stop sharing this project with ${email}`)) {
            return;
        }

        deleteProjectAssociation({
            id: id
        }).then(success => {
            reloadProjectData();
        })
    }

    const renderAssociations = () => {
        const hasUpdateRole = projectGridData.associationType === AssociationType.Owner || projectGridData.associationType === AssociationType.Writer;

        if (projectData === undefined || associationMap.size === 0) {
            return (<div>Loading...</div>);
        }

        return projectData.associations.map((a, i) => {
            const associationMapEntry = associationMap.get(a.id)!;
            if (associationMapEntry.formRef === undefined) {
                associationMapEntry.formRef = createRef<HTMLFormElement>();
                updateEntryRef(a.id, associationMapEntry.formRef);
            }
            
            const canUpdate = hasUpdateRole && a.type !== AssociationType.Owner && a.email !== currentUser.email;
            return (<>
                {canUpdate && associationMap.get(a.id)!.editing ? renderInteractiveRow(a, associationMapEntry.formRef) : renderReadonlyRow(a, canUpdate)}
                {i < projectData.associations.length - 1 ? <hr/> : ""}
            </>);
        });
    };

    const renderNewAssociation = () => {
        return (
            <Form onSubmit={onNewAssociationSubmit}>
                <Row>
                    <Col>
                        <Form.Group>
                            <Form.Label>User Email</Form.Label>
                            <Form.Control 
                                type="email" 
                                name="email" 
                                placeholder="Enter User Email" 
                                required={true} 
                                isInvalid={newShareEmailError !== undefined} 
                                onBlur={validateNewShareEmail}
                            />
                            <Form.Control.Feedback type="invalid">
                                {newShareEmailError}
                            </Form.Control.Feedback>
                        </Form.Group>
                    </Col>
                    <Col>

                        <Form.Group>
                            <Form.Label>Role</Form.Label>
                            <Form.Control as="select" name="type">
                                {associationTypeOptions}
                            </Form.Control>
                        </Form.Group>
                    </Col>
                </Row>
                <ButtonGroup>
                    <Button type="submit" className="btn-primary">Share</Button>
                </ButtonGroup>
            </Form>);
    }

    return (
        <Modal show={showShare} size='lg' onHide={onClose}>
            <Modal.Header closeButton={true}>
                <Modal.Title>Share Project: {projectGridData.name}</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                {renderAssociations()}
                <hr/>
                {renderNewAssociation()}
            </Modal.Body>
            <Modal.Footer>

            </Modal.Footer>
        </Modal>);
}