import React, { createRef, useEffect, useState } from 'react';
import { ButtonGroup, Col, Form, Modal, Row } from 'react-bootstrap';
import { CheckSquare, PencilSquare, TrashFill, XSquare } from 'react-bootstrap-icons';
import { AssociationType, ProjectAssociation, ProjectDetails, ProjectGrid } from '../../../../../models/project';
import { User } from '../../../../../models/user';

type ProjectShareProps = {
    projectGridData: ProjectGrid,
    showShare: boolean,
    currentUser: User,
    loadProjectDetail: (projectId: string) => Promise<ProjectDetails>;
    onClose: () => void;
}

type AssociationMapEntry = {
    editing: boolean,
    formRef: React.RefObject<HTMLFormElement> | undefined
};

export const ProjectShare = ({projectGridData, showShare, currentUser, loadProjectDetail, onClose}: ProjectShareProps) => {
    const [loading, setLoading] = useState(true);
    const [projectData, setProjectData] = useState<ProjectDetails | undefined>(undefined);
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

    useEffect(() => {
        const loadDetails = async () => {
            loadProjectDetail(projectGridData.id).then(details => {
                setProjectData(details);
                const newAssocMap = details.associations.reduce((m, v) => {
                    return m.set(v.user.id, {
                        editing: false,
                        formRef: undefined
                    }
                )}, new Map<string, AssociationMapEntry>());
                setAssociationMap(newAssocMap);

                setLoading(false);
            });
        };

        loadDetails();
    }, []);
    
    const associationTypeOptions = Object.keys(AssociationType).map(k => <option value={k} key={k}>{k}</option>);
    const renderInteractiveRow = (association: ProjectAssociation) => {
        return (<Form.Group>
            <input type="hidden" name="userId"/>
            <Form.Control as="select" name="type" defaultValue={association.type}>
                {associationTypeOptions}
            </Form.Control>
            <button><CheckSquare/></button>
            <button onClick={() => {
                updateEntryEditing(association.user.id, false);

                associationMap.get(association.user.id)!.formRef!.current!.reset();
            }}><XSquare/></button>
        </Form.Group>)
    };

    const renderReadonlyRow = (association: ProjectAssociation, canUpdate: boolean) => {
        return (<div>
            <div>{association.type}</div> 
            {canUpdate ? <ButtonGroup>
                <button onClick={() => updateEntryEditing(association.user.id, true)}><PencilSquare/></button>
                <button onClick={() => console.log("tadaa!")}><TrashFill/></button>
            </ButtonGroup>  : ""}
        </div>)
    };

    const onAssociationFormSubmit = (e: React.FormEvent) => {
        const formData = new FormData(e.target as HTMLFormElement);
        const type = formData.get('type') as string;
        const userId = formData.get('userId') as string;

        console.log("tadaa!");
    }

    const renderAssociations = () => {
        const hasUpdateRole = projectGridData.associationType === AssociationType.Owner || projectGridData.associationType === AssociationType.Writer;

        if (loading) {
            return (<div>Loading...</div>);
        }

        return projectData!.associations.filter(a => a.user.id !== currentUser.id).map(a => {
            const newRef = createRef<HTMLFormElement>();
            updateEntryRef(a.user.id, newRef);
            const canUpdate = hasUpdateRole && a.type !== AssociationType.Owner;
            return (
                <Form onSubmit={onAssociationFormSubmit} key={a.user.id} ref={newRef}>
                    <Row>
                        <Col>
                            <Form.Label>{a.user.name} ({a.user.email})</Form.Label>
                            {canUpdate && associationMap.get(a.user.id)?.editing ? renderInteractiveRow(a) : renderReadonlyRow(a, canUpdate)}
                        </Col>
                    </Row>
                </Form>
            );
        });
    };

    return (
        <Modal show={showShare} size='lg' onHide={onClose}>
            <Modal.Header>
                <Modal.Title>Share Project: {projectGridData.name}</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                <Form>
                {renderAssociations()}
                </Form>
            </Modal.Body>
            <Modal.Footer>

            </Modal.Footer>
        </Modal>);
}