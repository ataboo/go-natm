import React from 'react';
import './drop-zone.scss';
import { Status } from '../../../../../../models/status';

type DropZoneProps = {
    status: Status,
    moveCardToStatus: (draggedCardId: string, statusId: string) => void,
    draggedCardId: string
};

export function DropZone({status, moveCardToStatus, draggedCardId}: DropZoneProps) {
    const onDragOver = (event: React.DragEvent) => {
        moveCardToStatus(draggedCardId, status.id)
    };

    return (
    <div className="col-drop-zone"
        onDragOver={onDragOver}
        >
        {'Drop Zone Here'}
    </div>)
}