import React, { useState } from 'react';
import './card.scss';
import { ICardActions } from '../../icardactions';
import { CardHeader } from './cardheader';
import { TaskRead } from '../../../../../../models/task';
import { userNameAsInitials } from '../../../../../../services/implementation/stringhelpers';
import classNames from 'classnames';

type CardProps = {
    task: TaskRead
    statusId: string;
    cardActions: ICardActions;
};

type DragStartProps = {
    id: string
};

export const Card = ({task, statusId, cardActions} : CardProps) => {
    //until dragend event is fixed in firefox
    const getOpacity = () => cardActions.getDraggedCardId() === task.id ? 0.2 : 1;
    const elementRef = React.useRef<HTMLDivElement>(null);
    
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
        cardActions.saveProject();
    }

    function renderTiming() {
        return (<div className="timing-indicator">2hrs|4hrs</div>)
    }

    const cardClassNames = classNames(
        'drag-card',
        'draggable',
        { 
            'active-task-card': cardActions.getActiveTaskId() === task.id
        }
    )

    return (<div 
                ref={elementRef}
                className={cardClassNames}
                onDragStart = {(event) => onDragStart(event, {id: task.id})}
                draggable
                onDragOver={(event)=>onDragHover(event)}
                style={{opacity: getOpacity(), userSelect: "none"}}
            >
                <CardHeader 
                    onClickActivate={() => cardActions.setActiveTaskId(task.id)} 
                    onClickStop={() => cardActions.setActiveTaskId("")}
                    onClickEdit={() => {console.log("clicked edit!")}} 
                    cardActive={cardActions.getActiveTaskId() === task.id}
                    task={task}
                />
                <div className={"card-body-group"}>
                    <div className="assignee-link">
                        <a href="#">{task.assignee == null ? "" : (userNameAsInitials(task.assignee!))}</a>
                    </div>
                    {renderTiming()}
                </div>
            </div>);
};