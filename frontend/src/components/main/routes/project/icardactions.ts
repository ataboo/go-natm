import { CommentCreate, CommentRead, CommentUpdate } from "../../../../models/comment";
import { StatusCreate, StatusRead } from "../../../../models/status";
import { TaskCreate, TaskUpdate } from "../../../../models/task";
import { User } from "../../../../models/user";

export interface ICardActions {
    moveCardToStatus(draggedId: string, statusId: string): void

    swapCards(draggedId: string, statusId: string, ordinal: number): void

    setDraggedCardId(id: string): void

    getDraggedCardId(): string

    getActiveTaskId(): string | undefined

    startLoggingWork(id: string): void

    stopLoggingWork(): void

    saveTaskOrder(): void

    createStatus(createData: StatusCreate): Promise<StatusRead>

    createTask(createData: TaskCreate): Promise<boolean>

    archiveTask(taskId: string): Promise<boolean>

    archiveStatus(statusId: string): Promise<boolean>

    updateTask(updateData: TaskUpdate): Promise<boolean>

    getMaxStatusOrdinal(): number;

    stepStatusOrdinal(statusID: string, step: number): Promise<boolean>;

    loadComments(taskID: string): Promise<CommentRead[]>;

    addComment(data: CommentCreate): Promise<CommentRead>;

    deleteComment(id: string): Promise<boolean>;

    updateComment(data: CommentUpdate): Promise<CommentRead>;

    getCurrentUser(): User;
}