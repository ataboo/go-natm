import React from 'react';
import {PlayFill, StopFill} from 'react-bootstrap-icons';
import './cardheader.scss';
import { TaskRead } from '../../../../../../../models/task';
import { colorForTaskTag } from '../../../../../../../services/implementation/stringhelpers';
import { UpdateTask } from '../../updatetask';
import { ICardActions } from '../../../icardactions';

type CardHeaderProps = {
    cardActions: ICardActions
    task: TaskRead
}

export const CardHeader = ({cardActions, task} : CardHeaderProps) => {
    const cardActive = cardActions.getActiveTaskId() === task.id;
    
    const activateButton = () => {
        if (cardActive) {
            return (
                <button onClick={() => cardActions.setActiveTaskId("")} className="btn p-1">
                    <StopFill />
                </button>
            );
        }

        return (
            <button onClick={() => cardActions.setActiveTaskId(task.id)} className="btn p-1" >
                <PlayFill />
            </button>
        );
    } 
    
    return (<div className="drag-card-header" onDoubleClick={() => {}} style={{backgroundColor: colorForTaskTag(task)}}>
        <div className="task-title" title={task.title}>{task.identifier} - {task.title}</div>
        <div className="card-btn-group">

            <UpdateTask task={task} updateTask={cardActions.updateTask}/>
            <div className="card-btn-divider"></div>
            {activateButton()}
        </div>
    </div>)
}