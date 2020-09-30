import React, { useState } from 'react';
import { Button, Modal } from 'react-bootstrap';
import { TaskRead } from '../../../../../../models/task';
import './taskdetailmodal.scss';

type TaskDetailProps = {
    taskData: TaskRead,
    show: boolean,
    setShow: (show: boolean) => void,
}

export const TaskDetailModal = ({taskData, show, setShow}: TaskDetailProps) => {
    const taskDetailContent = (<>
            <div className="form-group d-flex">
                <label className="col-sm-2 col-form-label">Title</label>
                <div className="col-sm-10 pt-2">{taskData.title}</div>
            </div>
            <div className="form-group d-flex">
                <label className="col-sm-2 col-form-label">Description</label>
                <div className="col-sm-10 pt-2">{taskData.description}</div>
            </div>
            <div className="form-group d-flex">
                <label className="col-sm-2 col-form-label">Type</label>
                <div className="col-sm-10 pt-2">{taskData.type}</div>
            </div>
            <div className="form-group d-flex">
                <label className="col-sm-2 col-form-label">Assigned To</label>
                <div className="col-sm-10 pt-2">{taskData.assignee?.name}</div>
            </div>
        </>);

    return (<>
            <Modal show={show} onClose={() => setShow(false)} autoFocus={false} onHide={() => {setShow(false)}} size="lg">
                <Modal.Header closeButton={true} onClick={() => setShow(false)}>
                    <Modal.Title>Task Details: {taskData.identifier}</Modal.Title>
                </Modal.Header>

                <Modal.Body>
                    {taskDetailContent}
                </Modal.Body>

                <Modal.Footer>
                    <Button variant="secondary" onClick={() => setShow(false)}>Close</Button>
                </Modal.Footer>
            </Modal>
        </>);
};
