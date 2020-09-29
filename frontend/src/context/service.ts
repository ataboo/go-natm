import { createContext } from 'react';
import ProjectService from '../services/implementation/project-service';
import { IProjectService } from '../services/interface/iproject-service';

type ServiceContextProps = {
    projectService: IProjectService
};

export const ServiceContext = createContext<ServiceContextProps>(defaultServiceContext());

export function defaultServiceContext() : ServiceContextProps {
    return {
        projectService: new ProjectService()
    };
}

