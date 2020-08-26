import { Project } from "../../models/project";
import { StatusCreate } from "../../models/status";
import { TaskCreate, TaskUpdate } from "../../models/task";

export interface IProjectService {
    swapTasks(project: Project, draggedTaskId: string, droppedTaskStatusId: string, droppedTaskOrdinal: number): boolean

    moveCardToStatus(project: Project, draggedTaskId: string, statusId: string): boolean;

    getProjectList(): Promise<Project[]>;

    getProject(id: string): Promise<Project>

    saveProject(project: Project): Promise<any>;

    emptyProject(): Project;

    setActiveTaskId(id: string): Promise<boolean>;

    getActiveTaskId(): Promise<string>;

    createTaskStatus(data: StatusCreate): Promise<boolean>;

    createTask(data: TaskCreate): Promise<boolean>;

    archiveTask(taskId: string): Promise<boolean>;

    archiveStatus(statusId: string): Promise<boolean>;
    
    updateTask(updateData: TaskUpdate): Promise<boolean> 
}