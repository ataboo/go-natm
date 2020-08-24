import React from 'react';
import {Pencil, PlayFill, StopFill} from 'react-bootstrap-icons';
import './cardheader.scss';
import { TaskRead } from '../../../../../../../models/task';
import { colorForTaskTag } from '../../../../../../../services/implementation/stringhelpers';

type CardHeaderProps = {
    cardActive: boolean,
    onClickActivate: () => void
    onClickStop: () => void
    onClickEdit: () => void
    task: TaskRead
}

export const CardHeader = ({cardActive, onClickActivate, onClickStop, onClickEdit, task} : CardHeaderProps) => {
    const activateButton = () => {
        if (cardActive) {
            return (
                <button onClick={onClickStop} className="btn p-1">
                    <StopFill />
                </button>
            );
        }

        return (
            <button onClick={onClickActivate} className="btn p-1" >
                <PlayFill />
            </button>
        );
    } 
    
    return (<div className="drag-card-header" onDoubleClick={onClickEdit} style={{backgroundColor: colorForTaskTag(task)}}>
        <div className="task-title" title={task.title}>{task.identifier} - {task.title}</div>
        <div className="card-btn-group">
            <button onClick={onClickEdit} className="btn p-1">
                <Pencil />
            </button>
            <div className="card-btn-divider"></div>
            {activateButton()}
        </div>
    </div>)
}