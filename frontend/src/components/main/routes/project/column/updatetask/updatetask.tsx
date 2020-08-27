import React, { useRef } from 'react';
import { ModalForm } from '../../../../modalform';
import { Form } from 'react-bootstrap';
import { TaskRead, TaskUpdate } from '../../../../../../models/task';
import { TaskType } from '../../../../../../enums';

type UpdateTaskProps = {
    task: TaskRead
    updateTask: (taskData: TaskUpdate) => void
    show: boolean
    setShow: (show: boolean) => void
}

export const UpdateTask = ({task, updateTask, show, setShow}: UpdateTaskProps) => {
    const titleInput = useRef(null);
    const taskTypeOptions = Object.keys(TaskType).filter(k => !isNaN(+k)).map(k => {return {number: +k, label: TaskType[+k]}});

    const handleFormSubmit = (formData: FormData) => {
        updateTask({
            assigneeEmail: formData.get("assigneeEmail") as string,
            description: formData.get("description") as string,
            id: task.id,
            statusId: task.statusId,
            title: formData.get("title") as string,
            type: Number(formData.get("type") as string) as TaskType,
            estimatedTime: formData.get("estimatedTime") as string
        });
    }

    const formContent = (<> 
            <Form.Group controlId="title">
                <Form.Label>Title</Form.Label>
                <Form.Control autoFocus={true} type="text" name="title" required={true} ref={titleInput} defaultValue={task.title}></Form.Control>
            </Form.Group>

            <Form.Group controlId="description">
                <Form.Label>Description</Form.Label>
                <Form.Control as="textarea" rows={2} name="description" defaultValue={task.description}></Form.Control>
            </Form.Group>

            <Form.Group controlId="type">
                <Form.Label>Type</Form.Label>
                <Form.Control as="select" name="type" required={true} defaultValue={task.type}>
                    {taskTypeOptions.map(t => (<option value={t.number} key={t.number}>{t.label}</option>))}
                </Form.Control>
            </Form.Group>

            <Form.Group controlId="assigneeEmail">
                <Form.Label>Assigned To</Form.Label>
                <Form.Control type="email" name="assigneeEmail" defaultValue={task.assignee?.email}></Form.Control>
            </Form.Group>

            <Form.Group controlId="estimatedTime">
                <Form.Label>Estimated Time</Form.Label>
                <Form.Control type="text" name="estimatedTime" defaultValue={`${task.timing.estimated?.asHours()}h`}></Form.Control>
            </Form.Group>
        </>);

    return (<ModalForm 
        focusElement={titleInput}
        formContent={formContent}
        formElementId="update-task-form"
        onSubmit={handleFormSubmit}
        title={`Update Task ${task.identifier}`}
        show={show}
        setShow={setShow}
    />);
}
