import React, { createRef, MouseEvent, RefObject, useCallback, useEffect, useState } from 'react';
import { Button, ButtonGroup, Col, Form, Modal, Row } from 'react-bootstrap';
import { Check, CheckSquare, PencilSquare, TrashFill, X, XSquare } from 'react-bootstrap-icons';
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

    // const loadAndInitialize = useCallback(() => {
        
    // }, [projectGridData, setProjectData, setAssociationMap, setLoading, loading]);

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
    }, [setAssociationMap, setProjectData]);

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
            <Form onSubmit={onExistingAssociationSubmit} ref={ref}>
                <Form.Row>
                    <input type="hidden" name="id" value={association.id}/>

                    <Form.Control as="select" name="type" defaultValue={association.type}>
                        {associationTypeOptions}
                    </Form.Control>

                    <ButtonGroup>
                        <Button size='sm' className='btn-light' type='submit'><Check/></Button>
                        <Button size='sm' className='btn-light' type='button' onClick={() => {
                            associationMap.get(association.id)!.formRef!.current!.reset();
                            updateEntryEditing(association.id, false);

                        }}><X/></Button>
                    </ButtonGroup>
                </Form.Row>
            </Form>)
    };

    const renderReadonlyRow = (association: ProjectAssociationDetail, canUpdate: boolean) => {
        return (
            <div>
                <div>{association.type}</div> 
                {canUpdate ? <ButtonGroup>
                    <button type="button" onClick={() => updateEntryEditing(association.id, true)}><PencilSquare/></button>
                    <button type="button" onClick={() => onClickDeleteAssociation(association.id)}><TrashFill/></button>
                </ButtonGroup>  : ""}
            </div>)
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

    const onNewAssociationSubmit = (e: React.FormEvent) => {
        e.preventDefault();

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

    const onClickDeleteAssociation = (id: string) => {
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

        const otherAssociations = projectData!.associations.filter(a => a.email !== currentUser.email);

        if (otherAssociations.length === 0) {
            return (<div>This project is not shared with anyone.</div>);
        }

        return otherAssociations.map(a => {
            const associationMapEntry = associationMap.get(a.id)!;
            if (associationMapEntry.formRef === undefined) {
                associationMapEntry.formRef = createRef<HTMLFormElement>();
                updateEntryRef(a.id, associationMapEntry.formRef);
            }
            
            const canUpdate = hasUpdateRole && a.type !== AssociationType.Owner;
            return (
                <Row>
                    <Col>
                        <Form.Label>{a.email}</Form.Label>
                    </Col>
                    <Col>
                        {canUpdate && associationMap.get(a.id)!.editing ? renderInteractiveRow(a, associationMapEntry.formRef) : renderReadonlyRow(a, canUpdate)}
                    </Col>
                </Row>
            );
        });
    };

    const renderNewAssociation = () => {
        return (
            <Form onSubmit={onNewAssociationSubmit}>
                <Row>
                    <Col>
                        <Form.Group>
                            <Form.Label>User Email</Form.Label>
                            <Form.Control type="email" name="email" placeholder="Enter User Email" required={true}/>
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