import { User } from "./user";
import { TaskTiming } from "./tasktiming";
import { TaskType } from "../enums";

export interface TaskRead extends TaskCreate {
    assignee?: User;
    id: string;
    identifier: string;
    ordinal: number;
    timing: TaskTiming;
}

export interface TaskUpdate {
    assigneeEmail: string;
    description: string;
    estimatedTime: string;
    id: string;
    title: string;
    type: TaskType;
}

export interface TaskCreate {
    statusId: string;
    title: string;
    description: string;
    type: TaskType;
}
