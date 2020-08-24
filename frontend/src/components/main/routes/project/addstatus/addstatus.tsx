import React, { useState } from 'react';
import "./addstatus.scss";
import { Modal, Button, FormLabel, FormText, Form } from 'react-bootstrap';
import { StatusCreate } from '../../../../../models/status';
import { Plus } from 'react-bootstrap-icons';

type AddStatusProps = {
    createStatus: (status: StatusCreate) => void
}

export const AddStatus = ({createStatus}: AddStatusProps) => {
    const [show, setShow] = useState(false);

    const handleSubmit = (e: React.FormEvent) => {
        const formData = new FormData(e.currentTarget as HTMLFormElement);
        createStatus({
            name: formData.get("name") as string
        });
        
        setShow(false);

        e.preventDefault();
    };

    const handleShow = () => {
        setShow(true);
    };


    return (
        <div>
            <button className="btn btn-primary col-add-button p-1" onClick={handleShow}><Plus size={24}/></button>
            <Modal show={show} onClose={() => setShow(false)}>
                <Modal.Header closeButton={true}>
                    <Modal.Title>Create Task Status</Modal.Title>
                </Modal.Header>

                <Modal.Body>
                    <Form onSubmit={handleSubmit} id="create-status-form">
                        <Form.Group controlId="name">
                            <Form.Label>Status Name</Form.Label>
                            <Form.Control type="text" name="name" required={true}></Form.Control>
                        </Form.Group>
                    </Form>
                </Modal.Body>

                <Modal.Footer>
                    <Button variant="secondary" onClick={() => setShow(false)}>Close</Button>
                    <Button variant="primary" type="submit" form="create-status-form">Save</Button>
                </Modal.Footer>
            </Modal>
        </div>)
};
