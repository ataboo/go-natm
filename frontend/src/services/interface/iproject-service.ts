import { ProjectDetails, ProjectCreate } from "../../models/project";
import { StatusCreate } from "../../models/status";
import { TaskCreate, TaskUpdate } from "../../models/task";

export interface IProjectService {
    swapTasks(project: ProjectDetails, draggedTaskId: string, droppedTaskStatusId: string, droppedTaskOrdinal: number): boolean

    moveCardToStatus(project: ProjectDetails, draggedTaskId: string, statusId: string): boolean;

    getProjectList(): Promise<ProjectDetails[]>;

    getProject(id: string): Promise<ProjectDetails>

    saveProject(project: ProjectDetails): Promise<any>;

    createProject(project: ProjectCreate): Promise<boolean>

    emptyProject(): ProjectDetails;

    setActiveTaskId(id: string): Promise<boolean>;

    getActiveTaskId(): Promise<string>;

    createTaskStatus(data: StatusCreate): Promise<boolean>;

    createTask(data: TaskCreate): Promise<boolean>;

    archiveTask(taskId: string): Promise<boolean>;

    archiveStatus(statusId: string): Promise<boolean>;
    
    updateTask(updateData: TaskUpdate): Promise<boolean>
}