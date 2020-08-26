import React from 'react';
import './column.scss';
import { Card } from './card';
import { DropZone } from './drop-zone';
import { TaskRead, TaskCreate } from '../../../../../models/task';
import { StatusRead } from '../../../../../models/status';
import { ICardActions } from '../icardactions';
import { Dropdown } from 'react-bootstrap';
import { ThreeDotsVertical } from 'react-bootstrap-icons';
import { AddTask } from './addtask';

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
    
        const cards = tasks.map(t => (<Card 
            key={t.id}
            task={t}  
            statusId={t.statusId}
            cardActions={cardActions}
        />))

        cards.push((<AddTask createTask={(data) => cardActions.createTask(data)} statusId={status.id} key="add-task" />));

        return cards;
    }

    const archiveStatus = (event: React.MouseEvent) => {
        if (!window.confirm("Are you sure you want to archive this status?")) {
            return
        }

        cardActions.archiveStatus(status.id);
    }

    return (
    <div className="drag-column">
        <div className="column-header-group">
            <div className="column-header-text">{status.name}</div>
            <Dropdown>
                <Dropdown.Toggle variant="outline-default">
                    <ThreeDotsVertical/>
                </Dropdown.Toggle>

                <Dropdown.Menu>
                    <Dropdown.Item onClick={archiveStatus}>Archive</Dropdown.Item>
                </Dropdown.Menu>
            </Dropdown>
        </div>
        {renderCards()}
    </div>)
};