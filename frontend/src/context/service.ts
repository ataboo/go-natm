import { createContext } from 'react';
import { IProjectService } from '../services/interface/iproject-service';
import { FakeProjectService } from '../services/implementation/fake-project-service';

type ServiceContextProps = {
    projectService: IProjectService
};

export const ServiceContext = createContext<ServiceContextProps>(defaultServiceContext());

export function defaultServiceContext() {
    return {
        projectService: new FakeProjectService()
    };
}

