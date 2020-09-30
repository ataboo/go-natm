import { User } from "./user";

export interface CommentCreate {
    message: string,
    taskID: string
};

export interface CommentRead extends CommentCreate {
    author: User,
    createdAt: moment.Moment,
    updatedAt: moment.Moment
};