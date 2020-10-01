import { DateTime } from 'luxon';
import React from 'react';
import { Button, Card } from 'react-bootstrap';
import { Trash } from 'react-bootstrap-icons';
import { CommentRead } from '../../../../../../../models/comment';
import { User } from '../../../../../../../models/user';

type TaskCommentProps = {
    comment: CommentRead;
    deleteComment: (id: string) => void;
    currentUser: User;
}

export const TaskComment = ({comment, deleteComment, currentUser}: TaskCommentProps) => {
    const renderTrashButton = () => {
        if (currentUser.id === comment.author.id) {
            return <Button type="button" className="btn-light p-1" onClick={() => deleteComment(comment.id)}><Trash/></Button>
        }
    }
    
    
    return (
    <Card className="mb-2">
        <Card.Header className="d-flex justify-content-between align-items-center py-1">
                <span>{comment.author.name}@{DateTime.fromISO(comment.createdAt).toLocaleString(DateTime.DATETIME_SHORT)}</span>
                {renderTrashButton()}
        </Card.Header>
        <Card.Body>
            <Card.Text>{comment.message}</Card.Text>
        </Card.Body>
    </Card>);
};
