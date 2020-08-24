import { TaskRead } from "./task";
import { StatusRead } from "./status";

export interface Project {
    id: string;
    name: string;
    identifier: string;
    tasks: TaskRead[];
    statuses: StatusRead[];
}

export const cloneProject = (oldProject: Project) => {
    return {
        id: oldProject.id,
        name: oldProject.name,
        identifier: oldProject.identifier,
        tasks: [...oldProject.tasks],
        statuses: [...oldProject.statuses]
    }
}