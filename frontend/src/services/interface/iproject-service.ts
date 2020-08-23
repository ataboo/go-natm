import { Project } from "../../models/project";

export interface IProjectService {
    swapTasks(project: Project, draggedTaskId: string, droppedTaskStatusId: string, droppedTaskOrdinal: number): boolean

    moveCardToStatus(project: Project, draggedTaskId: string, statusId: string): boolean;

    getProjectList(): Promise<Project[]>;

    getProject(id: string): Promise<Project>

    saveProject(project: Project): Promise<any>;

    emptyProject(): Project;
}