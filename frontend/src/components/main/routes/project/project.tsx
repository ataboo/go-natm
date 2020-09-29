import React, { Component } from 'react';
import './project.scss';
import { Column } from './column';
import { ProjectDetails as ProjectModel, cloneProject, ProjectDetails } from '../../../../models/project';
import { ServiceContext } from '../../../../context/service';
import { ICardActions } from './icardactions';
import { IProjectService } from '../../../../services/interface/iproject-service';
import { AddStatus } from './addstatus/addstatus';
import { StatusCreate } from '../../../../models/status';
import { TaskUpdate } from '../../../../models/task';

interface IProjectState {
  projectData: ProjectModel;
  draggedCardId: string;
  activeTaskId?: string;
}

type ProjectProps ={
  id: string
};

export class Project extends Component<ProjectProps, IProjectState> {
  static contextType = ServiceContext;

  async componentDidMount() {
    const id = this.props.id;
    const projectData : ProjectDetails = await this.context.projectService.getProject(id);

    this.setState({
      draggedCardId: "",
      projectData: projectData,
      activeTaskId: await projectData.workingTaskID,
    });
  }

  moveCardToStatus() {
    return (draggedId: string, statusId: string) => {
      this.context.projectService.moveCardToStatus(this.state.projectData, draggedId, statusId)
      
      this.setState({
        projectData: cloneProject(this.state.projectData),
        draggedCardId: this.state.draggedCardId
      });
    };
  }

  moveCards() {
    return (draggedId: string, statusId: string, ordinal: number) => {
      this.context.projectService.swapTasks(this.state.projectData, draggedId, statusId, ordinal);

      this.setState({
        projectData: cloneProject(this.state.projectData),
        draggedCardId: this.state.draggedCardId
      });
    };
  };

  startLoggingWork() {
    return (id: string) => {
      const projectService: IProjectService = this.context.projectService;
      projectService.startLoggingWork(id)
        .then(success => {
          if (!success) {
            console.error("failed to activate task");
            return;
          }
          this.setState({
            projectData: this.state.projectData,
            draggedCardId: this.state.draggedCardId,
            activeTaskId: id
          })
        })
    }
  }

  stopLoggingWork() {
    return () => {
      const projectService: IProjectService = this.context.projectService;
      projectService.stopLoggingWork()
        .then(success => {
          if (!success) {
            console.error("failed to stop work");
            return;
          }

          this.setState({
            projectData: this.state.projectData,
            draggedCardId: this.state.draggedCardId,
            activeTaskId: undefined
          })
        });
    }
  }

  addStatus() {
    return (createData: StatusCreate) => {
      const projectService: IProjectService = this.context.projectService;

      const createAndUpdate = async () => {
        const success = await projectService.createTaskStatus(createData);
        if (!success) {
          throw new Error("failed to add status");
        }
        const newData = await projectService.getProject(this.props.id);

        this.setState({
          activeTaskId: newData.workingTaskID,
          draggedCardId: this.state.draggedCardId,
          projectData: newData,
        });
      }

      createAndUpdate();
    }
  }

  renderColumns() {
    if (!this.state || !this.state.projectData) {
      return (<div>Loading...</div>)
    }

    let sortedCards = (statusId: string) => this.state.projectData.tasks.filter(t => t.statusId === statusId).sort((a, b) => a.ordinal - b.ordinal);

    const cardActions: ICardActions = {
      getMaxStatusOrdinal: () => {
        return Math.max.apply(null, this.state.projectData.statuses.map(s => s.ordinal))
      },
      getDraggedCardId: () => this.state.draggedCardId,
      moveCardToStatus: this.moveCardToStatus(),
      setDraggedCardId: (cardId: string) => this.setState({
        projectData: this.state.projectData,
        draggedCardId: cardId,
        activeTaskId: this.state.activeTaskId
      }),
      swapCards: this.moveCards(),
      saveTaskOrder: () => this.context.projectService.saveTaskOrder(this.state.projectData),
      createStatus: async (createData) => {
        throw new Error("Not implemented");
      },
      createTask: async(createData) => {
        const success = await this.context.projectService.createTask(createData);
        if (success) {
          this.setState({
            activeTaskId: this.state.activeTaskId,
            draggedCardId: this.state.draggedCardId,
            projectData: await this.context.projectService.getProject(this.state.projectData.id)
          })
        }

        return success;
      },
      startLoggingWork: this.startLoggingWork(),
      stopLoggingWork: this.stopLoggingWork(),
      getActiveTaskId: () => this.state.activeTaskId,
      archiveStatus: async(statusId: string) => {
        const success = await this.context.projectService.archiveStatus(statusId);
        if (success) {
          this.setState({
            activeTaskId: this.state.activeTaskId,
            draggedCardId: this.state.draggedCardId,
            projectData: await this.context.projectService.getProject(this.state.projectData.id)
          })
        }

        return success;
      },
      archiveTask: async(taskId: string) => {
        const success = await this.context.projectService.archiveTask(taskId);
        if (success) {
          this.setState({
            activeTaskId: this.state.activeTaskId,
            draggedCardId: this.state.draggedCardId,
            projectData: await this.context.projectService.getProject(this.state.projectData.id)
          });
        }

        return success;
      },
      updateTask: async(updateData: TaskUpdate) => {
        const success = await this.context.projectService.updateTask(updateData);
        if (success) {
          this.setState({
            activeTaskId: this.state.activeTaskId,
            draggedCardId: this.state.draggedCardId,
            projectData: await this.context.projectService.getProject(this.state.projectData.id)
          });
        }

        return success;
      },
      stepStatusOrdinal: async(statusID: string, step: number) => {
        const success = await this.context.projectService.stepStatusOrdinal(statusID, step);
        if (success) {
          this.setState({
            activeTaskId: this.state.activeTaskId,
            draggedCardId: this.state.draggedCardId,
            projectData: await this.context.projectService.getProject(this.state.projectData.id)
          });
        }

        return success;
      }
    }

    return this.state.projectData.statuses.sort((a, b) => a.ordinal - b.ordinal).map((status, i) => (<Column 
      key={status.id} 
      tasks={sortedCards(status.id)} 
      status={status} 
      cardActions={cardActions}
    />))

  }

  render() {
    return (<div className="project">
              <div className="col-container">
                {this.renderColumns()}
                <AddStatus projectId={this.props.id} createStatus={this.addStatus()}/>
              </div>
            </div>)
  }
}
