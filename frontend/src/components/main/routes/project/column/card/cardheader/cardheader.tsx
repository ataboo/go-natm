import React, { useState } from 'react';
import {PlayFill, StopFill, PencilSquare, Trash} from 'react-bootstrap-icons';
import './cardheader.scss';
import { TaskRead } from '../../../../../../../models/task';
import { colorForTaskTag } from '../../../../../../../services/implementation/stringhelpers';
import { UpdateTask } from '../../updatetask';
import { ICardActions } from '../../../icardactions';
import { TaskDetailModal } from '../../taskdetailmodal';

type CardHeaderProps = {
    cardActions: ICardActions
    task: TaskRead
}

export const CardHeader = ({cardActions, task} : CardHeaderProps) => {
    const [showEdit, setShowEdit] = useState(false);
    const [showDetail, setShowDetail] = useState(false);
    const cardActive = cardActions.getActiveTaskId() === task.id;
    
    const activateButton = () => {
        if (cardActive) {
            return (
                <button onClick={() => cardActions.stopLoggingWork()} className="btn p-1">
                    <StopFill />
                </button>
            );
        }

        return (
            <button onClick={() => cardActions.startLoggingWork(task.id)} className="btn p-1" >
                <PlayFill />
            </button>
        );
    }

    const handleDeleteTask = async () => {
        if (!window.confirm("Are you sure you would like to archive this task?")) {
            return;
        }

        cardActions.archiveTask(task.id);
    }
    
    const editButton = () => (
        <button className="btn m-1 p-1" onClick={() => setShowEdit(true)}><PencilSquare /></button>
    );

    return (<div className="drag-card-header" onDoubleClick={() => {setShowDetail(true)}} style={{backgroundColor: colorForTaskTag(task)}}>
                <div className="task-title" title={task.title}>{task.identifier} - {task.title}</div>
                <div className="card-btn-group">
                    <TaskDetailModal taskData={task} show={showDetail} setShow={setShowDetail} />
                    <UpdateTask task={task} show={showEdit} setShow={setShowEdit} updateTask={cardActions.updateTask} />
                    <button className="btn m-1 p-1" onClick={handleDeleteTask}><Trash/></button>
                    <div className="card-btn-divider"></div>
                    {editButton()}
                    <div className="card-btn-divider"></div>
                    {activateButton()}
                </div>
            </div>)
}