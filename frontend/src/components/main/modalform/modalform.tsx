import { Modal, Button, FormLabel, FormText, Form } from 'react-bootstrap';
import React, { useState, useRef } from 'react';
import './modalform.scss';
import { ClassValue } from 'classnames/types';
import classNames from 'classnames';


type ModalFormProps = {
    formElementId: string
    title: string 
    formContent: JSX.Element
    showButtonContent: JSX.Element
    onSubmit: (formData: FormData) => void
    focusElement: React.RefObject<HTMLElement>
    buttonClasses: ClassValue[]
};

export const ModalForm = ({formElementId, title, formContent, showButtonContent, onSubmit, focusElement, buttonClasses} :ModalFormProps) => {
    const [show, setShow] = useState(false);

    const handleShow = () => {
        setShow(true);
    };
    
    const handleSubmit = (e: React.FormEvent) => {
        const formData = new FormData(e.currentTarget as HTMLFormElement);
        onSubmit(formData);
        setShow(false);

        e.preventDefault();
    };

    const handleOnEntered = (e: React.FormEvent) => {
        if (focusElement.current) {
            focusElement.current!.focus();
        }
    }

    const buttonClassNames = classNames(...["btn",...buttonClasses])

    return (<div>
                <button className={buttonClassNames} onClick={handleShow}>{showButtonContent}</button>
                <Modal show={show} onClose={() => setShow(false)} autoFocus={false} onEntered={handleOnEntered}>
                    <Modal.Header closeButton={true} onClick={() => setShow(false)}>
                        <Modal.Title>{title}</Modal.Title>
                    </Modal.Header>

                    <Modal.Body>
                        <Form onSubmit={handleSubmit} id={formElementId}>
                            {formContent}
                        </Form>
                    </Modal.Body>

                    <Modal.Footer>
                        <Button variant="secondary" onClick={() => setShow(false)}>Close</Button>
                        <Button variant="primary" type="submit" form={formElementId} autoFocus={true}>Save</Button>
                    </Modal.Footer>
                </Modal>
            </div>);
};
