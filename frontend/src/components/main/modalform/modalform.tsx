import { Modal, Button, Form } from 'react-bootstrap';
import React from 'react';
import './modalform.scss';


type ModalFormProps = {
    formElementId: string
    title: string 
    formContent: JSX.Element
    onSubmit: (formData: FormData) => void
    focusElement: React.RefObject<HTMLElement>
    show: boolean
    setShow: (show: boolean) => void
};

export const ModalForm = ({formElementId, title, formContent, onSubmit, focusElement, show, setShow} :ModalFormProps) => {    
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

    return (<>
                <Modal show={show} onClose={() => setShow(false)} autoFocus={false} onHide={() => {setShow(false)}} onEntered={handleOnEntered}>
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
            </>);
};
