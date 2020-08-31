import { ProjectDetails, ProjectCreate, ProjectGrid } from "../../models/project";
import { StatusCreate } from "../../models/status";
import { TaskCreate, TaskUpdate } from "../../models/task";

export interface IProjectService {
    swapTasks(project: ProjectDetails, draggedTaskId: string, droppedTaskStatusId: string, droppedTaskOrdinal: number): boolean

    moveCardToStatus(project: ProjectDetails, draggedTaskId: string, statusId: string): boolean;

    getProjectList(): Promise<ProjectGrid[]>;

    getProject(id: string): Promise<ProjectDetails>;

    saveProject(project: ProjectDetails): Promise<any>;

    createProject(project: ProjectCreate): Promise<boolean>;

    archiveProject(id: string): Promise<boolean>;

    emptyProject(): ProjectDetails;

    setActiveTaskId(id: string): Promise<boolean>;

    getActiveTaskId(): Promise<string>;

    createTaskStatus(data: StatusCreate): Promise<boolean>;

    createTask(data: TaskCreate): Promise<boolean>;

    archiveTask(taskId: string): Promise<boolean>;

    archiveStatus(statusId: string): Promise<boolean>;
    
    updateTask(updateData: TaskUpdate): Promise<boolean>
}