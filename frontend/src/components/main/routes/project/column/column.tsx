import React from 'react';
import './column.scss';
import { Card } from './card';
import { DropZone } from './drop-zone';
import { Task } from '../../../../../models/task';
import { Status } from '../../../../../models/status';

type ColumnProps = {
    tasks: Task[],
    status: Status,
    moveCardToStatus: (draggedCardId: string, statusId: string) => void,
    moveCards: (draggedCardId: string, droppedStatusId: string, droppedOrdinal: number) => void,
    draggedCardId: string,
    setDraggedCardId: (draggedCardId: string) => void,
    onDrop: () => void
}

export function Column({tasks, status, moveCards, moveCardToStatus, draggedCardId, setDraggedCardId, onDrop}: ColumnProps) {
    const renderCards = () => {
        if (tasks.length === 0) {
            return (<DropZone status={status} moveCardToStatus={moveCardToStatus} draggedCardId={draggedCardId}/>)
        }
    
        return tasks.map(c => (<Card 
            key={c.id} 
            id={c.id} 
            name={c.name} 
            statusId={c.statusId} 
            ordinal={c.ordinal} 
            moveCards={moveCards}
            draggedCardId={draggedCardId}
            setDraggedCardId={setDraggedCardId}
            onDrop={onDrop}
        />))    
    }

    return (
    <div 
        className="drag-column droppable" 
        style={{background: 'blue'}}
    >
        <div>{status.name} - {tasks.length}</div>
        {renderCards()}
    </div>)
};