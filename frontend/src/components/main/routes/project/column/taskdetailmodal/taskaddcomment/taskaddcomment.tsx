import React, { useRef, useState } from "react";
import { FormEvent } from "react";
import { Button, ButtonGroup, Form } from "react-bootstrap";

type AddCommentProps = {
    addNewComment: (message: string) => Promise<boolean>;
}

export const TaskAddComment = ({addNewComment}: AddCommentProps) => {
    const [showSubmit, setShowSubmit] = useState(false);
    const [lockInput, setLockInput] = useState(false);
    const formRef = useRef<HTMLFormElement>(null);

    const handleNewCommentSubmit = (e: FormEvent) => {
        e.preventDefault();
        setLockInput(true);
        const formData = new FormData(e.currentTarget as HTMLFormElement);
        addNewComment(formData.get("message") as string).then(async success => {
            formRef.current!.reset();
            setShowSubmit(false);
            setLockInput(false);
        });
    }

    const handleCancelComment = (e: React.MouseEvent<HTMLElement, MouseEvent>) => {
        formRef.current!.reset();
        
        setShowSubmit(false);
    }

    const handleFocusCommentControl = (e: React.FocusEvent<HTMLTextAreaElement>) => {
        setShowSubmit(true);
    }

    return (<Form ref={formRef} onSubmit={handleNewCommentSubmit} id="comment-form">
        <Form.Group controlId="message-control">
            <Form.Control disabled={lockInput} onFocus={handleFocusCommentControl} as="textarea" rows={2} name="message" placeholder="Add a new comment..." ></Form.Control>
        </Form.Group>

        {showSubmit ? (<ButtonGroup className="float-right">
            <Button disabled={lockInput} className="btn btn-light" onClick={handleCancelComment}>Cancel</Button>
            <Button disabled={lockInput} type="submit" className="btn btn-success">Comment</Button>
        </ButtonGroup>) : ""}
    </Form>)
};