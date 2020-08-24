import React from 'react';
import './column.scss';
import { Card } from './card';
import { DropZone } from './drop-zone';
import { TaskRead } from '../../../../../models/task';
import { StatusRead } from '../../../../../models/status';
import { ICardActions } from '../icardactions';
import { Dropdown, DropdownButton } from 'react-bootstrap';
import DropdownMenu from 'react-bootstrap/esm/DropdownMenu';
import { Dot, ThreeDotsVertical } from 'react-bootstrap-icons';

type ColumnProps = {
    tasks: TaskRead[],
    status: StatusRead,
    cardActions: ICardActions
}

export function Column({tasks, status, cardActions}: ColumnProps) {
    const renderCards = () => {
        if (tasks.length === 0) {
            return (<DropZone status={status} cardActions={cardActions} />)
        }
    
        return tasks.map(t => (<Card 
            key={t.id}
            task={t}  
            statusId={t.statusId}
            cardActions={cardActions}
        />))    
    }

    return (
    <div 
        className="drag-column droppable" 
    >
        <div className="column-header-group">
            <div className="column-header-text">{status.name}</div>
            <Dropdown>
                <Dropdown.Toggle variant="outline-default">
                    <ThreeDotsVertical/>
                </Dropdown.Toggle>

                <Dropdown.Menu>
                    <Dropdown.Item onClick={() => console.log("tadaa")}>Do thing</Dropdown.Item>
                </Dropdown.Menu>
            </Dropdown>
        </div>
        {renderCards()}
    </div>)
};