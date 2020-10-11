import React, { FormEvent, useState } from 'react';
import { Button, ButtonGroup, Form, Modal } from 'react-bootstrap';
import { CommentRead, CommentUpdate } from '../../../../../../models/comment';
import { TaskRead } from '../../../../../../models/task';
import { ICardActions } from '../../icardactions';
import './taskdetailmodal.scss';
import { TaskComment } from './taskcomment';
import { formatHMSDuration, formatHrDuration } from '../../../../../../constants';
import { setSyntheticTrailingComments } from 'typescript';
import { TaskType } from '../../../../../../enums';
import { DistributeVertical, PencilSquare } from 'react-bootstrap-icons';
import { TaskAddComment } from './taskaddcomment';

type TaskDetailProps = {
    taskData: TaskRead,
    show: boolean,
    setShow: (show: boolean) => void,
    cardActions: ICardActions
}

export const TaskDetailModal = ({ taskData, show, setShow, cardActions }: TaskDetailProps) => {
    const [comments, setComments] = useState<CommentRead[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [editingTask, setEditingTask] = useState(false);
    const taskTypeOptions = Object.keys(TaskType).map(k => {return {number: k, label: k}});

    const updateComment = async (data: CommentUpdate) => {
        const updatedComment = await cardActions.updateComment(data);
        const newComments = [...comments];
        const oldCommentIdx = newComments.findIndex(c => c.id === updatedComment.id);
        if (oldCommentIdx >= 0) {
            newComments[oldCommentIdx] = updatedComment;
            setComments(newComments);
        }
    }

    const renderCommentsSection = () => {
        if (loading) {
            return (<div>Loading Comments...</div>)
        }

        return comments.map(c => <TaskComment key={c.id} comment={c} deleteComment={onDeleteComment} updateComment={updateComment} currentUser={cardActions.getCurrentUser()} />);
    }

    const handleOnShow = async () => {
        await loadComments();
    }

    const handleSubmitTaskEdit = (e: FormEvent) => {
        e.preventDefault();

        setEditingTask(false);
    }

    const handleCancelEditTask = () => {
        setEditingTask(false);
    }

    const loadComments = async () => {
        const loadedComments = await cardActions.loadComments(taskData.id);
        setComments(loadedComments);
        setLoading(false);
    }

    const addNewComment = async (message: string): Promise<boolean> => {
        await cardActions.addComment({
            message: message,
            taskID: taskData.id,
        }).then(async () => {
            await loadComments();
        });

        return true;
    };

    const onDeleteComment = (id: string) => {
        if (!window.confirm("Are you sure you want to permenantly delete this comment?")) {
            return;
        }

        setLoading(true);
        cardActions.deleteComment(id).then(async () => {
            await loadComments();
        });
    };

    const taskDetailContent = () => {
        // const titleElement = editingTask ? () : (<div className="col-sm-10 pt-2">{taskData.title}</div>);

        const titleElement = editingTask ?
            (<Form.Control autoFocus={true} type="text" name="title" required={true} defaultValue={taskData.title}></Form.Control>) :
            (<div>{taskData.title}</div>);

        const descriptionElement = editingTask ?
            (<Form.Control as="textarea" rows={2} name="description" defaultValue={taskData.description}></Form.Control>) :
            (<div>{taskData.description}</div>);

        const typeElement = editingTask ?
            (<Form.Control as="select" name="type" required={true} defaultValue={taskData.type}>
                {taskTypeOptions.map(t => (<option value={t.number} key={t.number}>{t.label}</option>))}
            </Form.Control>) :
            (<div>{taskData.type}</div>);

        const assignedToElement = editingTask ?
            (<Form.Control type="email" name="assigneeEmail" defaultValue={taskData.assignee?.email}></Form.Control>) :
            (<div>{taskData.assignee?.name}</div>);

        const estimateElement = editingTask ?
            (<Form.Control type="text" name="estimatedTime" defaultValue={taskData.timing.estimate == undefined ? "" : (taskData.timing.estimate! / 60.0).toFixed(0) + 'm'}></Form.Control>) :
            (<div>{formatHrDuration(taskData.timing.estimate)}</div>);
        

        return (<div>
            <Form className="mb-3" onSubmit={handleSubmitTaskEdit}>
                <Form.Row>
                    <Form.Group className="col-md-6" controlId="title">
                        <Form.Label>Title</Form.Label>
                        {titleElement}
                    </Form.Group>

                    <Form.Group className="col-md-6" controlId="type">
                        <Form.Label>Type</Form.Label>
                        {typeElement}
                    </Form.Group>
                </Form.Row>

                <Form.Row>
                    <Form.Group className="col-md-6" controlId="description">
                        <Form.Label>Description</Form.Label>
                        {descriptionElement}
                    </Form.Group>
                    <Form.Group className="col-md-6" controlId="assigneeEmail">
                        <Form.Label>Assigned To</Form.Label>
                        {assignedToElement}
                    </Form.Group>
                </Form.Row>

                <Form.Row>
                    <Form.Group className="col-md-6" controlId="loggedTime">
                        <Form.Label>Logged Time</Form.Label>
                        <div>{formatHMSDuration(taskData.timing.current)}</div>
                    </Form.Group>

                    <Form.Group className="col-md-6" controlId="estimatedTime">
                        <Form.Label>Estimated Time</Form.Label>
                        {estimateElement}
                    </Form.Group>
                </Form.Row>
                {() => editingTask ? 
                    (<ButtonGroup>
                        <Button className="btn btn-light" onClick={handleCancelEditTask}>Cancel</Button>
                        <Button type="submit" className="btn btn-success">Save</Button>
                    </ButtonGroup>) : ""}
                
            </Form>

            <hr/>
            <h3>Comments</h3>

            {renderCommentsSection()}

            {loading ? "" : <TaskAddComment addNewComment={addNewComment}></TaskAddComment>}
        </div>)
    }

    return (
        <Modal show={show} onShow={handleOnShow} onClose={() => setShow(false)} autoFocus={false} onHide={() => { setShow(false) }} size="xl">
            <Modal.Header closeButton={true} onClick={() => setShow(false)}>
                <div className="d-flex justify-content-between">
                    <Modal.Title>Task Details: {taskData.identifier}</Modal.Title>
                    <button disabled={editingTask} onClick={(e) => {e.preventDefault(); setEditingTask(true);}} className="btn btn-light ml-3"><PencilSquare/></button>
                </div>
            </Modal.Header>

            <Modal.Body>
                {taskDetailContent()}
            </Modal.Body>

            <Modal.Footer>
                <Button variant="secondary" onClick={() => setShow(false)}>Close</Button>
            </Modal.Footer>
        </Modal>
    );
};
