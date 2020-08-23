import React, { Component } from 'react';
import './project.scss';
import { Column } from './column';
import { Project as ProjectModel } from '../../../../models/project';
import { ServiceContext } from '../../../../context/service';

interface IProjectState {
  projectData: ProjectModel;
  draggedCardId: string;
}

export class Project extends Component<any, IProjectState> {
  static contextType = ServiceContext;

  async componentDidMount() {
    const id = this.props.match.id;
    const projectData = await this.context.getProject(id);

    this.setState({
      draggedCardId: "",
      projectData: projectData
    });
  }

  moveCardToStatus(draggedId: string, statusId: string) {
    let changed = this.context.projectService.moveCardToStatus(this.state.projectData, draggedId, statusId)

    const newProjectData = {
      id: this.state.projectData.id,
      name: this.state.projectData.name,
      tasks: [...this.state.projectData.tasks],
      statuses: [...this.state.projectData.statuses],
    };

    if (changed) {
      this.context.projectService.saveProject(newProjectData);
    }
    
    this.setState({
      projectData: newProjectData,
      draggedCardId: this.state.draggedCardId
    });
  }

  moveCards(draggedId: string, statusId: string, ordinal: number) {
    var changed = this.context.projectService.swapTasks(this.state.projectData, draggedId, statusId, ordinal);

    const newProjectData = {
      id: this.state.projectData.id,
      name: this.state.projectData.name,
      tasks: [...this.state.projectData.tasks],
      statuses: [...this.state.projectData.statuses]
    };

    if(changed) {
      this.context.projectService.saveProject(newProjectData);
    }

    this.setState({
      projectData: newProjectData,
      draggedCardId: this.state.draggedCardId
    });
  };

  renderColumns() {
    if (this.state.projectData === undefined) {
      return (<div>Loading...</div>)
    }

    let sortedCards = (statusId: string) => this.state.projectData.tasks.filter(t => t.statusId === statusId).sort((a, b) => a.ordinal - b.ordinal);

    return this.state.projectData.statuses.map((status, i) => (<Column 
      key={status.id} 
      tasks={sortedCards(status.id)} 
      status={status} 
      moveCards={this.moveCards} 
      moveCardToStatus={this.moveCardToStatus}
      draggedCardId={this.state.draggedCardId}
      setDraggedCardId={(cardId) => this.setState({
        projectData: this.state.projectData,
        draggedCardId: cardId
      })}
      onDrop={() => this.context.projectService.saveProject(this.state.projectData)}
    />))

  }

  render() {
    return (<div className="project">
              <div className="col-container">
                {this.renderColumns()}
              </div>
            </div>)
  }
}
