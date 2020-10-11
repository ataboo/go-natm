import { User } from "./user";
import {DateTime} from "luxon";

export interface CommentCreate {
    message: string,
    taskID: string
};

export interface CommentUpdate extends CommentCreate {
    id: string
}

export interface CommentRead extends CommentCreate {
    id: string,
    author: User,
    createdAt: string,
    updatedAt: string,
};