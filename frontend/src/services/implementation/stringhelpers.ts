import { User } from "../../models/user";
import { TaskRead } from "../../models/task";
import { TaskType } from "../../enums";

export const userNameAsInitials = (user: User): string => {
    const names = user.name.split(" ");

    return names.map(n => n[0]).join(".");
}

export const colorForTaskTag = (task: TaskRead): string => {
    switch (task.type) {
        case TaskType.Admin:
            return "#00ff0088"
        case TaskType.Problem:
            return "#ff000088";
        case TaskType.Task:
            return "#0000ff88";
        default:
            throw new Error("Not supported.");
    }
}