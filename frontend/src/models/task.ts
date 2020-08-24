import { User } from "./user";
import { TaskTiming } from "./tasktiming";
import { TaskType } from "../enums";

export interface TaskRead extends TaskCreate {
    id: string;
}

export interface TaskCreate {
    statusId: string;
    title: string;
    identifier: string;
    ordinal: number;
    assignee: User|null;
    timing: TaskTiming;
    type: TaskType;
}