import React, { useState, useEffect } from 'react';
import './card.scss';
import { ICardActions } from '../../icardactions';
import { CardHeader } from './cardheader';
import { TaskRead } from '../../../../../../models/task';
import { userNameAsInitials } from '../../../../../../services/implementation/stringhelpers';
import classNames from 'classnames';
import { formatHMSDuration, formatReadibleDuration } from '../../../../../../constants';

type CardProps = {
    task: TaskRead
    statusId: string;
    cardActions: ICardActions;
};

type DragStartProps = {
    id: string
};

export const Card = ({task, statusId, cardActions} : CardProps) => {
    const getOpacity = () => cardActions.getDraggedCardId() === task.id ? 0.2 : 1;
    const elementRef = React.useRef<HTMLDivElement>(null);

    const [currentTime, setCurrentTime] = useState(task.timing.current);

    useEffect(() => {
        if (cardActions.getActiveTaskId() === task.id) {
            setTimeout(activeTaskTick, 1000);
        }
    });

    const activeTaskTick = () => {
        setCurrentTime(currentTime + 1);
    }
    
    function onDragStart(event: React.DragEvent<HTMLDivElement>, {id} : DragStartProps) {
        elementRef!.current!.addEventListener('dragend', onDragEnd, false);
        event.dataTransfer!.dropEffect = 'move';

        cardActions.setDraggedCardId(id);
    }

    function onDragHover(event: React.DragEvent<HTMLDivElement>) {
        var draggedCardId = cardActions.getDraggedCardId();
        if (draggedCardId && draggedCardId !== task.id) {
            cardActions.swapCards(draggedCardId, statusId, task.ordinal);
        }

        event.preventDefault();
    }

    function onDragEnd(event: DragEvent) {
        if(elementRef.current) {
            elementRef.current.removeEventListener('dragend', onDragEnd, false);
        }
        cardActions.setDraggedCardId("");
        cardActions.saveTaskOrder();
    }

    function renderTiming() {
        return (<div className="timing-indicator">{formatHMSDuration(currentTime)} | {formatReadibleDuration(task.timing.estimate)}</div>)
    }

    const activeTaskId = cardActions.getActiveTaskId();

    const cardClassNames = classNames(
        'drag-card',
        'draggable',
        { 
            'active-task-card': activeTaskId === task.id,
            'innactive-task-card': activeTaskId && activeTaskId !== task.id
        }
    )

    return (<div ref={elementRef}
                 className={cardClassNames}
                 onDragStart = {(event) => onDragStart(event, {id: task.id})}
                 draggable
                 onDragOver={(event)=>onDragHover(event)}
                 style={{opacity: getOpacity(), userSelect: "none"}}>
                <CardHeader cardActions={cardActions} task={task} currentTime={currentTime}/>
                <div className={"card-body-group"}>
                    <div className="assignee-link">
                        <a href={"mailto:"+task.assignee?.email}>{task.assignee == null ? "" : (userNameAsInitials(task.assignee!))}</a>
                    </div>
                    {renderTiming()}
                </div>
            </div>);
};