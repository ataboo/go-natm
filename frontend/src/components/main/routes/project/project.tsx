import React, { Component } from 'react';
import './project.scss';
import { Column } from './column';
import { ProjectDetails as ProjectModel, cloneProject } from '../../../../models/project';
import { ServiceContext } from '../../../../context/service';
import { ICardActions } from './icardactions';
import { IProjectService } from '../../../../services/interface/iproject-service';
import { AddStatus } from './addstatus/addstatus';
import { StatusCreate } from '../../../../models/status';
import { TaskUpdate } from '../../../../models/task';

interface IProjectState {
  projectData: ProjectModel;
  draggedCardId: string;
  activeTaskId: string;
}

type ProjectProps ={
  id: string
};

export class Project extends Component<ProjectProps, IProjectState> {
  static contextType = ServiceContext;

  async componentDidMount() {
    const id = this.props.id;
    const projectData = await this.context.projectService.getProject(id);
    const activeTaskId = await this.context.projectService.getActiveTaskId();

    this.setState({
      draggedCardId: "",
      projectData: projectData,
      activeTaskId: activeTaskId
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

  setActiveTaskId() {
    return (id: string) => {
      const projectService: IProjectService = this.context.projectService;
      projectService.setActiveTaskId(id)
        .then((success) => {
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
          activeTaskId: this.state.activeTaskId,
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
      getDraggedCardId: () => this.state.draggedCardId,
      moveCardToStatus: this.moveCardToStatus(),
      setDraggedCardId: (cardId: string) => this.setState({
        projectData: this.state.projectData,
        draggedCardId: cardId,
        activeTaskId: this.state.activeTaskId
      }),
      swapCards: this.moveCards(),
      saveProject: () => this.context.projectService.saveProject(this.state.projectData),
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
      setActiveTaskId: this.setActiveTaskId(),
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
      }
    }

    return this.state.projectData.statuses.map((status, i) => (<Column 
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
                <AddStatus createStatus={this.addStatus()}/>
              </div>
            </div>)
  }
}
