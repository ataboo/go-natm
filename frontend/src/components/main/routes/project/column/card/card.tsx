import React from 'react';
import './card.scss';

type CardProps = {
    name: string,
    id: string,
    statusId: string,
    ordinal: number,
    moveCards: (draggedCardId: string, droppedStatusId: string, droppedOrdinal: number) => void,
    draggedCardId: string,
    setDraggedCardId: (draggedCardId: string) => void,
    onDrop: () => void
};

type DragStartProps = {
    id: string
};

export const Card = ({name, id, statusId, ordinal, moveCards, draggedCardId, setDraggedCardId, onDrop} : CardProps) => {
    //until dragend event is fixed in firefox
    const getOpacity = () => draggedCardId === id ? 0.2 : 1;
    const elementRef = React.useRef<HTMLDivElement>(null);
    
    function onDragStart(event: React.DragEvent<HTMLDivElement>, {id} : DragStartProps) {
        elementRef!.current!.addEventListener('dragend', onDragEnd, false);
        event.dataTransfer!.dropEffect = 'move';

        setDraggedCardId(id);
    }

    function onDragHover(event: React.DragEvent<HTMLDivElement>) {
        if (draggedCardId && draggedCardId !== id) {
            moveCards(draggedCardId, statusId, ordinal);
        }

        event.preventDefault();
    }

    function onDragEnd(event: DragEvent) {
        if(elementRef.current) {
            elementRef.current.removeEventListener('dragend', onDragEnd, false);
        }
        setDraggedCardId("");

        onDrop();
    }

    return (<div 
                ref={elementRef}
                className="drag-card draggable"
                onDragStart = {(event) => onDragStart(event, {id: id})}
                draggable
                onDragOver={(event)=>onDragHover(event)}
                style={{opacity: getOpacity(), userSelect: "none"}}
            >
                {name} - {ordinal}
            </div>);
};