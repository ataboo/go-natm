import { StatusCreate, StatusRead } from "../../../../models/status";
import { TaskCreate, TaskRead, TaskUpdate } from "../../../../models/task";

export interface ICardActions {
    moveCardToStatus(draggedId: string, statusId: string): void

    swapCards(draggedId: string, statusId: string, ordinal: number): void

    setDraggedCardId(id: string): void

    getDraggedCardId(): string

    getActiveTaskId(): string

    setActiveTaskId(id: string): void

    saveProject(): void

    createStatus(createData: StatusCreate): Promise<StatusRead>

    createTask(createData: TaskCreate): Promise<boolean>

    archiveTask(taskId: string): Promise<boolean>

    archiveStatus(statusId: string): Promise<boolean>

    updateTask(updateData: TaskUpdate): Promise<boolean>
}