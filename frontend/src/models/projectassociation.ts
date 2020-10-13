import { AssociationType } from "../enums";
import { User } from "./user";

export interface ProjectAssociationDetail {
    id: string,
    email: string,
    type: AssociationType,
};

export interface ProjectAssociationCreate {
    projectId: string,
    email: string,
    type: AssociationType
};

export interface ProjectAssociationUpdate {
    id: string,
    type: AssociationType
};

export interface ProjectAssociationDelete {
    id: string
}