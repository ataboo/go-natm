import React, { useRef } from 'react';
import './addtask.scss';
import { Form } from 'react-bootstrap';
import { TaskCreate } from '../../../../../../models/task';
import { Plus } from 'react-bootstrap-icons';
import { TaskType } from '../../../../../../enums';
import { ModalForm } from '../../../../modalform';

type AddTaskProps = {
    createTask: (status: TaskCreate) => void,
    statusId: string,
}

export const AddTask = ({createTask, statusId}: AddTaskProps) => {
    const titleInput = useRef<HTMLInputElement>(null);

    const handleSubmit = (formData: FormData) => {
        createTask({
            description: formData.get("description") as string,
            title: formData.get("title") as string,
            statusId: statusId,
            type: Number(formData.get("type") as string) as TaskType
        });
    };

    const taskTypeOptions = Object.keys(TaskType).filter(k => !isNaN(+k)).map(k => {return {number: +k, label: TaskType[+k]}});


    const addTaskContent = (<> 
            <Form.Group controlId="title">
                <Form.Label>Title</Form.Label>
                <Form.Control autoFocus={true} type="text" name="title" required={true} ref={titleInput}></Form.Control>
            </Form.Group>

            <Form.Group controlId="description">
                <Form.Label>Description</Form.Label>
                <Form.Control as="textarea" rows={2} name="description"></Form.Control>
            </Form.Group>

            <Form.Group controlId="type">
                <Form.Label>Type</Form.Label>
                <Form.Control as="select" name="type" required={true}>
                    {taskTypeOptions.map(t => (<option value={t.number} key={t.number}>{t.label}</option>))}
                </Form.Control>
            </Form.Group>
        </>);

    return (
        <ModalForm 
            title="Create New Task" 
            formContent={addTaskContent} 
            focusElement={titleInput}
            formElementId="add-task-form" 
            onSubmit={handleSubmit}
            showButtonContent={(<Plus size={24}/>)}
            buttonClasses={["btn-primary", "p-1", "m-2"]}
        />);
};
