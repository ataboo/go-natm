import { Project } from "../../models/project";
import { IProjectService } from "../interface/iproject-service";
import ProjectService from "./project-service";
import { User } from "../../models/user";
import { TaskTiming } from "../../models/tasktiming";
import moment from "moment";
import { TaskType } from "../../enums";
import { StatusCreate } from "../../models/status";

export class FakeProjectService implements IProjectService {
    cachedProjectData: Project
    realService: IProjectService
    activeTaskId: string

    constructor() {
        const assignedUser: User = {
            id: "1",
            name: "Alex Raboud"
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
            identifier: "PRJ",
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
                    type: TaskType.Admin
                },
                {
                    id: "2",
                    title: "card b",
                    statusId: "1",
                    ordinal: 1,
                    identifier: "PRJ-2",
                    assignee: assignedUser,
                    timing: timing2,
                    type: TaskType.Admin
                },
                {
                    id: "3",
                    title: "card c",
                    statusId: "2",
                    ordinal: 0,
                    identifier: "PRJ-3",
                    assignee: assignedUser,
                    timing: timing3,
                    type: TaskType.Problem
                },
                {
                    id: "4",
                    title: "card d this one has a really long title that goes on and on and on and on",
                    statusId: "2",
                    ordinal: 1,
                    identifier: "PRJ-4",
                    assignee: assignedUser,
                    timing: timing1,
                    type: TaskType.Task
                },
                {
                    id: "5",
                    title: "card e",
                    statusId: "2",
                    ordinal: 2,
                    identifier: "PRJ-5",
                    assignee: assignedUser,
                    timing: timing1,
                    type: TaskType.Task
                },
                {
                    id: "6",
                    title: "card f",
                    statusId: "2",
                    ordinal: 3,
                    identifier: "PRJ-6",
                    assignee: assignedUser,
                    timing: timing1,
                    type: TaskType.Task
                }
            ]
        }
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
    emptyProject(): Project {
        return this.realService.emptyProject();
    }

    swapTasks(project: Project, draggedTaskId: string, droppedTaskStatusId: string, droppedTaskOrdinal: number): boolean {
        return this.realService.swapTasks(project, draggedTaskId, droppedTaskStatusId, droppedTaskOrdinal);
    }
    moveCardToStatus(project: Project, draggedTaskId: string, statusId: string): boolean {
        return this.realService.moveCardToStatus(project, draggedTaskId, statusId);
    }

    async getProjectList(): Promise<Project[]> {
        return [this.cachedProjectData];
    }

    async getProject(projectId: string) : Promise<Project> {
        return this.cachedProjectData;
    }

    async saveProject(project: Project): Promise<any> {
        this.cachedProjectData = project;

        console.log("saved project!");
    }
}
