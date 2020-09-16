export interface StatusRead extends StatusCreate {
    id: string;
}

export interface StatusCreate {
    projectId: string;
    name: string;
}