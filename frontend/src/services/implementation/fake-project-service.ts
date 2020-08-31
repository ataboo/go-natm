import { ProjectDetails, ProjectCreate, ProjectGrid } from "../../models/project";
import { IProjectService } from "../interface/iproject-service";
import ProjectService from "./project-service";
import { User } from "../../models/user";
import { TaskTiming } from "../../models/tasktiming";
import moment from "moment";
import { TaskType } from "../../enums";
import { StatusCreate } from "../../models/status";
import { TaskCreate, TaskUpdate } from "../../models/task";

export class FakeProjectService implements IProjectService {
    cachedProjectData: ProjectDetails
    realService: IProjectService
    activeTaskId: string

    constructor() {
        const assignedUser: User = {
            id: "1",
            name: "Alex Raboud",
            email: "doubar2001@gmail.com"
        }

        const timing1: TaskTiming = {
            current: moment.duration(3.5, "hours"),
            estimated: moment.duration(2, "days")
        }

        const timing2: TaskTiming = {
            current: moment.duration(1, "week"),
            estimated: null
        }

        const timing3: TaskTiming = {
            current: moment.duration(0),
            estimated: null
        }

        this.activeTaskId = "";
        this.realService = new ProjectService();
        this.cachedProjectData = {
            id: "1",
            name: "My Project",
            abbreviation: "PRJ",
            description: "description here",
            statuses: [
                {
                    id: "1",
                    name: "Back Burner"
                },
                {
                    id: "2",
                    name: "Front Burner"
                }
            ],
            tasks: [
                {
                    id: "1",
                    title: "card a",
                    statusId: "1",
                    ordinal: 0,
                    identifier: "PRJ-1",
                    assignee: assignedUser,
                    timing: timing1,
                    type: TaskType.Admin,
                    description: "description"
                },
                {
                    id: "2",
                    title: "card b",
                    statusId: "1",
                    ordinal: 1,
                    identifier: "PRJ-2",
                    assignee: assignedUser,
                    timing: timing2,
                    type: TaskType.Admin,
                    description: "description"
                },
                {
                    id: "3",
                    title: "card c",
                    statusId: "2",
                    ordinal: 0,
                    identifier: "PRJ-3",
                    assignee: assignedUser,
                    timing: timing3,
                    type: TaskType.Problem,
                    description: "description"
                },
                {
                    id: "4",
                    title: "card d this one has a really long title that goes on and on and on and on",
                    statusId: "2",
                    ordinal: 1,
                    identifier: "PRJ-4",
                    assignee: assignedUser,
                    timing: timing1,
                    type: TaskType.Task,
                    description: "description"
                },
                {
                    id: "5",
                    title: "card e",
                    statusId: "2",
                    ordinal: 2,
                    identifier: "PRJ-5",
                    assignee: assignedUser,
                    timing: timing1,
                    type: TaskType.Task,
                    description: "description"
                },
                {
                    id: "6",
                    title: "card f",
                    statusId: "2",
                    ordinal: 3,
                    identifier: "PRJ-6",
                    assignee: assignedUser,
                    timing: timing1,
                    type: TaskType.Task,
                    description: "description"
                }
            ]
        }
    }

    async createProject(project: ProjectCreate): Promise<boolean> {
        return await this.realService.createProject(project);
    }

    async archiveProject(id: string): Promise<boolean> {
        return await this.realService.archiveProject(id)
    }

    async updateTask(updateData: TaskUpdate): Promise<boolean> {
        const task = this.cachedProjectData.tasks.find(t => t.id === updateData.id);
        if (!task) {
            return false;
        }

        if(updateData.assigneeEmail) {
            task.assignee = {
                email: updateData.assigneeEmail,
                name: "Assigned User",
                id: "23"
            };
        } else {
            task.assignee = null;
        }

        task.timing.estimated = moment.duration(updateData.estimatedTime);

        task.description = updateData.description;
        task.title = updateData.title;
        task.type = updateData.type;

        return true;
    }

    async archiveTask(taskId: string): Promise<boolean> {
        this.cachedProjectData.tasks = this.cachedProjectData.tasks.filter(t => t.id !== taskId);

        return true;
    }

    async archiveStatus(statusId: string): Promise<boolean> {
        this.cachedProjectData.statuses = this.cachedProjectData.statuses.filter(s => s.id !== statusId);

        if (this.cachedProjectData.statuses.length === 0) {
            this.cachedProjectData.tasks = [];
        } else {
            this.cachedProjectData.tasks.filter(t => t.statusId === statusId).forEach(t => t.statusId = this.cachedProjectData.statuses[0].id);
        }

        return true;
    }

    async createTask(data: TaskCreate): Promise<boolean> {
        const maxId = Math.max(...this.cachedProjectData.tasks.map(t => +t.id));
        const maxOrdinal = Math.max(...this.cachedProjectData.tasks.map(t => t.ordinal));
        const maxIdentifier = Math.max(...this.cachedProjectData.tasks.map(t => +(t.identifier.split('-')[1])));
        
        this.cachedProjectData.tasks.push({
            assignee: null,
            description: data.description,
            identifier: 'PRJ-' + (maxIdentifier + 1).toString(),
            id: (maxId + 1).toString(),
            ordinal: maxOrdinal + 1,
            statusId: data.statusId,
            timing: {
                current: moment.duration(0),
                estimated: moment.duration(0)
            },
            title: data.title,
            type: data.type
        });

        return true;
    }
    
    async createTaskStatus(data: StatusCreate): Promise<boolean> {
        const ids = this.cachedProjectData.statuses.map(s => s.id);
        let newId = "";
        for(let i=0; i<1000; i++) {
            newId = i.toString();
            if (!ids.includes(newId)) {
                break;
            }
        }
        
        this.cachedProjectData.statuses.push({id: newId, name: data.name})

        return true;
    }
    async setActiveTaskId(id: string): Promise<boolean> {
        this.activeTaskId = id;

        return true;
    }

    async getActiveTaskId(): Promise<string> {
        return this.activeTaskId;
    }

    emptyProject(): ProjectDetails {
        return this.realService.emptyProject();
    }

    swapTasks(project: ProjectDetails, draggedTaskId: string, droppedTaskStatusId: string, droppedTaskOrdinal: number): boolean {
        return this.realService.swapTasks(project, draggedTaskId, droppedTaskStatusId, droppedTaskOrdinal);
    }

    moveCardToStatus(project: ProjectDetails, draggedTaskId: string, statusId: string): boolean {
        return this.realService.moveCardToStatus(project, draggedTaskId, statusId);
    }

    async getProjectList(): Promise<ProjectGrid[]> {
        return this.realService.getProjectList()
    }

    async getProject(projectId: string) : Promise<ProjectDetails> {
        return this.cachedProjectData;
    }

    async saveProject(project: ProjectDetails): Promise<any> {
        this.cachedProjectData = project;

        console.log("saved project!");
    }
}
