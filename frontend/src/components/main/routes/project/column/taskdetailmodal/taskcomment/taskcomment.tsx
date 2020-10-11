import { DateTime } from 'luxon';
import React, { FormEvent, useState } from 'react';
import { Button, ButtonGroup, Card, Form } from 'react-bootstrap';
import { CheckSquare, PencilSquare, Trash, XSquare } from 'react-bootstrap-icons';
import { CommentRead, CommentUpdate } from '../../../../../../../models/comment';
import { User } from '../../../../../../../models/user';

type TaskCommentProps = {
    comment: CommentRead;
    deleteComment: (id: string) => void;
    updateComment: (data: CommentUpdate) => Promise<void>;
    currentUser: User;
}

export const TaskComment = ({comment, deleteComment, updateComment, currentUser}: TaskCommentProps) => {
    const [editing, setEditing] = useState(false);
    const [submitting, setSubmitting] = useState(false);

    const renderButtons = () => {
        if (currentUser.id === comment.author.id) {
            return (<div>
                <Button type="button" className="btn-light p-1" onClick={() => setEditing(true)}><PencilSquare/></Button>
                <Button type="button" className="btn-light p-1" onClick={() => deleteComment(comment.id)}><Trash/></Button>
            </div>);
        }
    }

    const renderEditingForm = () => {
        return (
            <Form onSubmit={onEditFormSubmit} id="comment-form">
                <Form.Group controlId="message-control">
                    <Form.Control disabled={submitting} as="textarea" rows={2} name="message" defaultValue={comment.message} ></Form.Control>
                </Form.Group>

                <ButtonGroup>
                    <Button className="btn btn-light" onClick={() => setEditing(false)}>Cancel</Button>
                    <Button disabled={submitting} type="submit" className="btn btn-success">Save</Button>
                </ButtonGroup>
            </Form>);
    }

    const renderEditingReadonly = () => {
        return (<span>{comment.message}</span>);
    }
    
    const onEditFormSubmit = (e: FormEvent) => {
        e.preventDefault();
        const formData = new FormData(e.currentTarget as HTMLFormElement);

        setSubmitting(true);
        updateComment({
            id: comment.id,
            message: formData.get("message") as string,
            taskID: comment.taskID
        }).then((success) => {
            setSubmitting(false);
            setEditing(false);
        });
    }
    
    return (
    <Card className="mb-2">
        <Card.Header className="d-flex justify-content-between align-items-center py-1">
                <span>{comment.author.name}@{DateTime.fromISO(comment.createdAt).toLocaleString(DateTime.DATETIME_SHORT)}</span>
                {renderButtons()}
        </Card.Header>
        <Card.Body>
            <Card.Text>{editing ? renderEditingForm() : renderEditingReadonly()}</Card.Text>
        </Card.Body>
    </Card>);
};
