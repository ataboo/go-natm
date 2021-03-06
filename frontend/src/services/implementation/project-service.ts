import { ProjectDetails, ProjectCreate, ProjectGrid, ProjectTaskOrder } from "../../models/project";
import { IProjectService } from "../interface/iproject-service";
import { StatusCreate } from "../../models/status";
import { TaskCreate, TaskUpdate } from "../../models/task";
import axios, { AxiosInstance } from "axios";
import { CommentCreate, CommentRead, CommentUpdate } from "../../models/comment";
import { ProjectAssociationCreate, ProjectAssociationDelete, ProjectAssociationUpdate } from "../../models/projectassociation";

export default class ProjectService implements IProjectService {
    client: AxiosInstance
    hostUri: string;
    activeTaskId: string;

    constructor() {
        this.client = axios.create({
            withCredentials: true,
        });

        this.activeTaskId = "";

        this.hostUri = "http://localhost:8080/api/v1/";
    }

    async updateTask(updateData: TaskUpdate): Promise<boolean> {
        const response = await this.client.post(`${this.hostUri}task/update`, JSON.stringify(updateData));
        
        return response.status === 200;
    }
    
    async archiveTask(taskID: string): Promise<boolean> {
        const response = await this.client.post(`${this.hostUri}task/archive`, JSON.stringify({task_id: taskID}));

        return response.status === 200;
    }

    async archiveStatus(statusID: string): Promise<boolean> {
        const response = await this.client.post(`${this.hostUri}status/archive`, JSON.stringify({status_id: statusID}));

        return response.status === 200;
    }

    async createTask(data: TaskCreate): Promise<boolean> {
        const response = await this.client.post(`${this.hostUri}task/create`, JSON.stringify(data));

        return response.status === 200;
    }

    async createTaskStatus(data: StatusCreate): Promise<boolean> {
        const response = await this.client.post(`${this.hostUri}status/`, JSON.stringify(data));
        
        return response.status === 200;
    }
    
    async startLoggingWork(id: string): Promise<boolean> {
        const response = await this.client.post(`${this.hostUri}task/startLoggingWork`, JSON.stringify({ task_id: id }));
        if (response.status === 200) {
            this.activeTaskId = id;

            return true;
        }
 
        return false;
    }

    async stopLoggingWork(): Promise<boolean> {
        const response = await this.client.post(`${this.hostUri}task/stopLoggingWork`);
        if (response.status === 200) {
            this.activeTaskId = "";
            return true;
        }

        return false;
    }

    async swapTasks(project: ProjectDetails, draggedTaskId: string, droppedTaskStatusId: string, droppedTaskOrdinal: number): Promise<boolean> {
        let draggedTask = project.tasks.find(t => t.id === draggedTaskId);
        if (draggedTask === undefined) {
            throw new Error("Failed to find task: " + draggedTaskId);
        }

        if (draggedTask.ordinal === droppedTaskOrdinal && draggedTask.statusId === droppedTaskStatusId) {
            return false;
        }

        let newStatusTasks = project.tasks.filter(c => c.statusId === droppedTaskStatusId && c.id !== draggedTaskId);
        newStatusTasks.sort((a, b) => a.ordinal - b.ordinal);
    
        if (draggedTask.statusId === droppedTaskStatusId) {
            let droppedTask = newStatusTasks.find(c => c.ordinal === droppedTaskOrdinal);
            if (droppedTask !== undefined) {
                droppedTask.ordinal = draggedTask.ordinal;
            }
        } else {   
            let oldStatusTasks = project.tasks.filter(c => c.statusId === draggedTask!.statusId && c.id !== draggedTaskId);
            oldStatusTasks.sort((a, b) => a.ordinal - b.ordinal);
            for(let i=droppedTaskOrdinal; i<newStatusTasks.length; i++) {
                newStatusTasks[i].ordinal = i+1;
            }
    
            for(let i=0; i<oldStatusTasks.length; i++) {
                oldStatusTasks[i].ordinal = i; 
            }
        }
    
        draggedTask.ordinal = droppedTaskOrdinal;
        draggedTask.statusId = droppedTaskStatusId;

        return true;
    };

    async moveCardToStatus(project: ProjectDetails, draggedTaskId: string, statusId: string): Promise<boolean> {
        let draggedTask = project.tasks.find(c => c.id === draggedTaskId);
        if (draggedTask === undefined) {
            throw new Error("Failed to find task: " + draggedTaskId);
        }

        if (draggedTask.statusId === statusId) {
            return false;
        }
        
        var oldStatusTasks = project.tasks.filter(c => c.statusId === draggedTask!.statusId && c.id !== draggedTaskId);

        for(let i=0; i<oldStatusTasks.length; i++) {
            oldStatusTasks[i].ordinal = i;
        }

        draggedTask.statusId = statusId;
        draggedTask.ordinal = 0;

        return true;
    }
    
    async getProjectList(): Promise<ProjectGrid[]> {
        try {
            const response = await this.client.get(this.hostUri + "project/")

            return response.data as ProjectGrid[];
        } catch(e) {
            throw e;
        }
    }

    async getProject(projectId: string): Promise<ProjectDetails> {
        const response = await this.client.get(`${this.hostUri}project/${projectId}`);
        return response.data as ProjectDetails;
    };

    async createProject(project: ProjectCreate): Promise<boolean> {
        const response = await this.client.post(this.hostUri + "project/", JSON.stringify(project));

        return response.status === 200;
    }

    async archiveProject(id: string): Promise<boolean> {
        const response = await this.client.post(this.hostUri + "project/archive/", JSON.stringify({
            projectID: id
        }));

        return response.status === 200;
    }

    async saveTaskOrder(taskOrder: ProjectTaskOrder): Promise<boolean> {
        var response = await this.client.post(`${this.hostUri}project/setTaskOrder`, JSON.stringify(taskOrder));

        return response.status === 200;
    }

    async stepStatusOrdinal(statusID: string, step: number): Promise<boolean> {
        var response = await this.client.post(`${this.hostUri}status/stepOrdinal`, JSON.stringify({status_id: statusID, step: step}));

        return response.status === 200;
    }

    async getComments(taskID: string): Promise<CommentRead[]> {
        var response = await this.client.get(`${this.hostUri}task/${taskID}/comment`);

        return response.data as CommentRead[];
    }

    async addComment(createData: CommentCreate): Promise<CommentRead> {
        var response = await this.client.post(`${this.hostUri}task/comment`, JSON.stringify(createData));

        return response.data as CommentRead;
    }

    async deleteComment(commentID: string): Promise<boolean> {
        var response = await this.client.post(`${this.hostUri}task/comment/delete`, JSON.stringify({commentID: commentID}));

        return response.status === 200;
    }

    async updateComment(data: CommentUpdate): Promise<CommentRead> {
        return (await this.client.post(`${this.hostUri}task/comment/update`, JSON.stringify(data))).data as CommentRead;
    }

    async updateProjectAssociation(data: ProjectAssociationUpdate): Promise<boolean> {
        return (await this.client.post(`${this.hostUri}projectassociation/update`, JSON.stringify(data))).status === 200;
    }
    
    async deleteProjectAssociation(data: ProjectAssociationDelete): Promise<boolean> {
        return (await this.client.post(`${this.hostUri}projectassociation/delete`, JSON.stringify(data))).status === 200;
    }

    async createProjectAssociation(data: ProjectAssociationCreate): Promise<boolean> {
        return (await this.client.post(`${this.hostUri}projectassociation/`, JSON.stringify(data))).status === 200;
    }
}
