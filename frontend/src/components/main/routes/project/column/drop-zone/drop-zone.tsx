import React from 'react';
import './drop-zone.scss';
import { StatusRead } from '../../../../../../models/status';
import { ICardActions } from '../../icardactions';

type DropZoneProps = {
    status: StatusRead,
    cardActions: ICardActions
};

export function DropZone({status, cardActions}: DropZoneProps) {
    const onDragOver = (event: React.DragEvent) => {
        cardActions.moveCardToStatus(cardActions.getDraggedCardId(), status.id)
    };

    return (
    <div className="col-drop-zone"
        onDragOver={onDragOver}
        >
        {'Drop Zone Here'}
    </div>)
}