import { User } from "../../models/user";


export interface IMainActions {
    // const [currentTime, setCurrentTime] = useState(task.timing.current);

    // useEffect(() => {
    //     if (cardActions.getActiveTaskId() === task.id) {
    //         setTimeout(activeTaskTick, 1000);
    //     }
    // });

    // const activeTaskTick = () => {
    //     setCurrentTime(currentTime + 1);
    // }

    startLoggingTask(taskId: string): void;

    stopLoggingTask(): void;

    logout(): Promise<boolean>;

    activeTaskId?: string;

    currentUser: User|null;

    refreshTimedTask(): void;
}

/*
// start logging work
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
*/

/*
//stop logging work

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
    */