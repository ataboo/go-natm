export interface StatusRead extends StatusCreate {
    id: string;
    ordinal: number;
}

export interface StatusCreate {
    projectId: string;
    name: string;
}