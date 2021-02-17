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
            return "#85e8a3"
        case TaskType.Problem:
            return "#ffd09f";
        case TaskType.Task:
            return "#b6e4ff";
        default:
            throw new Error("Not supported." + task.type);

            // 155, 199, 255
    }
}