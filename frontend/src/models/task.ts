import { User } from "./user";
import { TaskTiming } from "./tasktiming";
import { TaskType } from "../enums";

export interface TaskRead extends TaskCreate {
    assignee: User|null;
    id: string;
    identifier: string;
    ordinal: number;
    timing: TaskTiming;
}

export interface TaskUpdate extends TaskCreate {
    id: string;
    assigneeEmail: string;
}

export interface TaskCreate {
    statusId: string;
    title: string;
    description: string;
    type: TaskType;
}