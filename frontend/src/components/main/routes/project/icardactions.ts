import { StatusCreate, StatusRead } from "../../../../models/status";
import { TaskCreate, TaskRead } from "../../../../models/task";

export interface ICardActions {
    moveCardToStatus(draggedId: string, statusId: string): void

    swapCards(draggedId: string, statusId: string, ordinal: number): void

    setDraggedCardId(id: string): void

    getDraggedCardId(): string

    getActiveTaskId(): string

    setActiveTaskId(id: string): void

    saveProject(): void

    createStatus(createData: StatusCreate): Promise<StatusRead>

    createTask(createData: TaskCreate): Promise<TaskRead>

    
}