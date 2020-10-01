import React, { FormEvent, useState } from 'react';
import { Button, ButtonGroup, Card, Form, Modal } from 'react-bootstrap';
import { CommentRead } from '../../../../../../models/comment';
import { TaskRead } from '../../../../../../models/task';
import { ICardActions } from '../../icardactions';
import './taskdetailmodal.scss';
import {DateTime} from "luxon";
import { Trash } from 'react-bootstrap-icons';
import { TaskComment } from './taskcomment';

type TaskDetailProps = {
    taskData: TaskRead,
    show: boolean,
    setShow: (show: boolean) => void,
    cardActions: ICardActions
}

export const TaskDetailModal = ({taskData, show, setShow, cardActions}: TaskDetailProps) => {
    const [comments, setComments] = useState<CommentRead[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [showCommentForm, setShowCommentForm] = useState<boolean>(false);

    const renderCommentsSection = () => {
        if (loading) {
            return (<div>Loading Comments...</div>)
        }

        return comments.map(c => <TaskComment key={c.id} comment={c} deleteComment={onDeleteComment} currentUser={cardActions.getCurrentUser()}/>);
    }

    const renderCommentForm = () => {
        if (loading) {
            return;
        }
        
        if (showCommentForm) {
            return (<Form onSubmit={handleSubmit} id="comment-form">
                <Form.Group controlId="message-control">
                    <Form.Control as="textarea" rows={2} name="message"></Form.Control>
                </Form.Group>

                <ButtonGroup>
                    <Button className="btn btn-light" onClick={handleCancelComment}>Cancel</Button>
                    <Button type="submit" className="btn btn-success">Save</Button>
                </ButtonGroup>
            </Form>)
        } else {
            return (<Button className="btn-success" onClick={()=>setShowCommentForm(true)}>Comment</Button>);
        }
    }

    const handleOnShow = async () => {
        await loadComments();
    }

    const handleCancelComment = () => {
        if (!window.confirm("Are you sure you want to cancel creating a comment?")) {
            return;
        }

        setShowCommentForm(false)
    }

    const loadComments = async () => {
        const loadedComments = await cardActions.loadComments(taskData.id);
        setComments(loadedComments);
        setLoading(false);
    }

    const handleSubmit = (e: FormEvent) => {
        const formData = new FormData(e.currentTarget as HTMLFormElement);
        setLoading(true);
        cardActions.addComment({
            message: formData.get("message") as string,
            taskID: taskData.id,
        }).then(async () => {
            setShowCommentForm(false);
            await loadComments();
        });
    };

    const onDeleteComment = (id: string) => {
        if(!window.confirm("Are you sure you want to permenantly delete this comment?")) {
            return;
        }

        setLoading(true);
        cardActions.deleteComment(id).then(async () => {
            await loadComments();
        });
    };
    
    const taskDetailContent = () => (<div>
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

            {renderCommentsSection()}

            {renderCommentForm()}
        </div>);

    return (<>
            <Modal show={show} onShow={handleOnShow} onClose={() => setShow(false)} autoFocus={false} onHide={() => {setShow(false)}} size="lg">
                <Modal.Header closeButton={true} onClick={() => setShow(false)}>
                    <Modal.Title>Task Details: {taskData.identifier}</Modal.Title>
                </Modal.Header>

                <Modal.Body>
                    {taskDetailContent()}
                </Modal.Body>

                <Modal.Footer>
                    <Button variant="secondary" onClick={() => setShow(false)}>Close</Button>
                </Modal.Footer>
            </Modal>
        </>);
};
