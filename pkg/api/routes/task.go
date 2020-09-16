package routes

import (
	"errors"
	"net/http"

	"github.com/ataboo/go-natm/pkg/api/data"
	"github.com/gin-gonic/gin"
)

func registerTaskRoutes(e *gin.RouterGroup) {
	g := e.Group("/tasks")

	g.GET("/:taskID", handleGetTask)
	g.POST("/", handleCreateTask)
	g.POST("/archive/", handleArchiveTask)
	// g.PUT("/:projectID", handleUpdate)
	// g.DELETE("/:projectID", handleDelete)
}

func handleGetTask(ctx *gin.Context) {
	userID := data.MustGetActingUserID(ctx)
	taskID, fail := ctx.Params.Get("taskID")
	if fail {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	task, err := taskRepo.Find(taskID, userID)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, data.TaskGrid{
		Description: task.Description,
		ID:          task.ID,
		Identifier:  data.TaskIdentifier(task.R.TaskStatus.R.Project.Abbreviation, task.Number),
		StatusID:    task.TaskStatusID,
		Timing: data.TimingGrid{
			Current:  "todo",
			Estimate: "todo",
		},
		Title: task.Title,
		Type:  task.TaskType,
	})
}

func handleCreateTask(ctx *gin.Context) {
	userID := data.MustGetActingUserID(ctx)
	createData := data.TaskCreate{}
	err := ctx.BindJSON(&createData)

	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = taskRepo.Create(&createData, userID)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func handleArchiveTask(ctx *gin.Context) {
	ctx.AbortWithError(http.StatusInternalServerError, errors.New("Not implemented yet"))
}
