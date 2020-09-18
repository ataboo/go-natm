import { ProjectDetails, ProjectCreate, ProjectGrid, ProjectTaskOrder } from "../../models/project";
import { StatusCreate } from "../../models/status";
import { TaskCreate, TaskUpdate } from "../../models/task";

export interface IProjectService {
    swapTasks(project: ProjectDetails, draggedTaskId: string, droppedTaskStatusId: string, droppedTaskOrdinal: number): Promise<boolean>

    moveCardToStatus(project: ProjectDetails, draggedTaskId: string, statusId: string): Promise<boolean>;

    getProjectList(): Promise<ProjectGrid[]>;

    getProject(id: string): Promise<ProjectDetails>;

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

    saveTaskOrder(taskOrderData: ProjectTaskOrder): Promise<boolean>
}