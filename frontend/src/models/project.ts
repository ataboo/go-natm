import { Task } from "./task";
import { Status } from "./status";

export interface Project {
    id: string;
    name: string;
    tasks: Task[];
    statuses: Status[];
}