import { DateTime } from 'luxon';
import React, { FormEvent, useState } from 'react';
import { Button, Card } from 'react-bootstrap';
import { PencilSquare, Trash } from 'react-bootstrap-icons';
import { CommentRead } from '../../../../../../../models/comment';
import { User } from '../../../../../../../models/user';

type TaskCommentProps = {
    comment: CommentRead;
    deleteComment: (id: string) => void;
    currentUser: User;
}

export const TaskComment = ({comment, deleteComment, editComment, currentUser}: TaskCommentProps) => {
    const [editing, setEditing] = useState(false);
    
    const renderButtons = () => {
        if (currentUser.id === comment.author.id) {
            return (<div>
                <Button type="button" className="btn-light p-1" onClick={() => setEditing(true)}><PencilSquare/></Button>
                <Button type="button" className="btn-light p-1" onClick={() => deleteComment(comment.id)}><Trash/></Button>
            </div>);
        }
    }
    
    const onEditFormSubmit = (e: FormEvent) => {

    }
    
    return (
    <Card className="mb-2">
        <Card.Header className="d-flex justify-content-between align-items-center py-1">
                <span>{comment.author.name}@{DateTime.fromISO(comment.createdAt).toLocaleString(DateTime.DATETIME_SHORT)}</span>
                {renderButtons()}
        </Card.Header>
        <Card.Body>
            <Card.Text>{{if(editing) {
                    return (<span>{comment.message}</span>);
                }}}</Card.Text>
        </Card.Body>
    </Card>);
};
