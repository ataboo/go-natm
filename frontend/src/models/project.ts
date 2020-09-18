import { TaskRead } from "./task";
import { StatusRead } from "./status";

export interface ProjectDetails extends ProjectCreate {
    id: string;
    tasks: TaskRead[];
    statuses: StatusRead[];
}

export interface ProjectCreate {
    name: string;
    abbreviation: string;
    description: string;
}

export interface ProjectGrid {
    id: string;
    name: string;
    abbreviation: string;
    associationType: ProjectAssociation,
    lastUpdated: number
}

export interface ProjectTaskOrder {
    id: string,
    tasks: TaskOrder[],
}

export interface TaskOrder {
    id: string,
    statusId: string,
    ordinal: number
}

export const cloneProject = (oldProject: ProjectDetails) => {
    return {
        id: oldProject.id,
        name: oldProject.name,
        abbreviation: oldProject.abbreviation,
        description: oldProject.description,
        tasks: [...oldProject.tasks],
        statuses: [...oldProject.statuses]
    }
}

export enum ProjectAssociation {
    Owner,
    Writer,
    Reader
};
