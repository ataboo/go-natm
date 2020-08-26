import React, { useRef } from 'react';
import "./addstatus.scss";
import { Form } from 'react-bootstrap';
import { StatusCreate } from '../../../../../models/status';
import { Plus } from 'react-bootstrap-icons';
import { ModalForm } from '../../../modalform';

type AddStatusProps = {
    createStatus: (status: StatusCreate) => void
}

export const AddStatus = ({createStatus}: AddStatusProps) => {
    const nameInput = useRef<HTMLInputElement>(null);

    const handleSubmit = (formData: FormData) => {
        createStatus({
            name: formData.get("name") as string
        });    
    };

    const buttonContent = (<Plus size={24}/>);

    const formContent = (<Form.Group controlId="name">
                            <Form.Label>Status Name</Form.Label>
                            <Form.Control type="text" name="name" required={true} autoFocus={true} ref={nameInput}></Form.Control>
                        </Form.Group>);

    return (<ModalForm 
        focusElement={nameInput} 
        formContent={formContent} 
        formElementId="add-status-form" 
        onSubmit={handleSubmit} 
        showButtonContent={buttonContent} title="Create New Status"
        buttonClasses={["btn-primary", "p-1", "m-2"]} 
    />)
};
