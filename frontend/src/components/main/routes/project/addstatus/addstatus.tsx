import React, { useRef, useState } from 'react';
import "./addstatus.scss";
import { Form } from 'react-bootstrap';
import { StatusCreate } from '../../../../../models/status';
import { Plus } from 'react-bootstrap-icons';
import { ModalForm } from '../../../modalform';

type AddStatusProps = {
    projectId: string
    createStatus: (status: StatusCreate) => void
}

export const AddStatus = ({projectId, createStatus}: AddStatusProps) => {
    const nameInput = useRef<HTMLInputElement>(null);
    const [show, setShow] = useState(false);

    const handleSubmit = (formData: FormData) => {
        createStatus({
            projectId: projectId,
            name: formData.get("name") as string
        });    
    };

    const formContent = (<Form.Group controlId="name">
                            <Form.Label>Status Name</Form.Label>
                            <Form.Control type="text" name="name" required={true} autoFocus={true} ref={nameInput}></Form.Control>
                        </Form.Group>);

    return (
        <div>
            <button className="btn m-2 p-1 btn-primary add-status-button" onClick={() => {setShow(true)}}>
                <div className="d-flex"><Plus size={24}/><span className="mr-2">Status</span></div>
            </button>
            <ModalForm 
                focusElement={nameInput} 
                formContent={formContent} 
                formElementId="add-status-form" 
                onSubmit={handleSubmit} 
                title="Create New Status"
                setShow={setShow}
                show={show} 
            />
        </div>)
};
