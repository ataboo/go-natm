import React, { useRef, useState } from 'react';
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
    const [show, setShow] = useState(false);

    const handleSubmit = (formData: FormData) => {
        createTask({
            description: formData.get("description") as string,
            title: formData.get("title") as string,
            statusId: statusId,
            type: formData.get("type") as string as TaskType
        });
    };

    const taskTypeOptions = Object.keys(TaskType).map(k => {return {number: k, label: k}});

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
        <>
            <button className="btn m-2 p-1 btn-primary add-task-button" onClick={() => {setShow(true)}}>
                <div className="d-flex"><Plus size={24}/> <span className="mr-2">Task</span></div>
            </button>
            <ModalForm 
                title="Create New Task" 
                formContent={addTaskContent} 
                focusElement={titleInput}
                formElementId="add-task-form" 
                onSubmit={handleSubmit}
                setShow={setShow}
                show={show}
            />
        </>);
};
