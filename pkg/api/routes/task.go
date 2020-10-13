package routes

import (
	"net/http"

	"github.com/ataboo/go-natm/pkg/api/data"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null/v8"
)

func registerTaskRoutes(e *gin.RouterGroup) {
	g := e.Group("/task")

	g.GET("/:taskID", handleGetTask)
	g.POST("/create", handleCreateTask)
	g.POST("/archive", handleArchiveTask)
	g.POST("/update", handleUpdateTask)
	g.POST("/startLoggingWork", handleStartLoggingWork)
	g.POST("/stopLoggingWork", handleStopLoggingWork)
	g.GET("/:taskID/comments", handleGetComments)
	g.POST("/comments", handleCreateComment)
	g.POST("/comments/delete", handleDeleteComment)
	g.POST("/comments/update", handleUpdateComment)
}

func handleGetTask(ctx *gin.Context) {
	userID := data.MustGetActingUserID(ctx)
	taskID, ok := ctx.Params.Get("taskID")
	if !ok {
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
		Timing: &data.TimingGrid{
			Current:  0,
			Estimate: null.IntFrom(0),
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

	ctx.Status(200)
}

func handleArchiveTask(ctx *gin.Context) {
	userID := data.MustGetActingUserID(ctx)
	archiveData := struct {
		TaskID string `json:"task_id"`
	}{}

	err := ctx.BindJSON(&archiveData)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = taskRepo.Archive(archiveData.TaskID, userID)
	if err != nil {
		handleErrorWithStatus(err, ctx)
		return
	}

	ctx.Status(200)
}

func handleUpdateTask(ctx *gin.Context) {
	userID := data.MustGetActingUserID(ctx)
	updateData := data.TaskUpdate{}
	err := ctx.BindJSON(&updateData)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = taskRepo.Update(&updateData, userID)
	if err != nil {
		handleErrorWithStatus(err, ctx)
		return
	}

	ctx.Status(200)
}

func handleStartLoggingWork(ctx *gin.Context) {
	userID := data.MustGetActingUserID(ctx)
	taskData := struct {
		TaskID string `json:"task_id"`
	}{}
	err := ctx.BindJSON(&taskData)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = taskRepo.StartLoggingWork(userID, taskData.TaskID)
	if err != nil {
		handleErrorWithStatus(err, ctx)
		return
	}

	ctx.Status(200)
}

func handleStopLoggingWork(ctx *gin.Context) {
	userID := data.MustGetActingUserID(ctx)

	err := taskRepo.StopLoggingWork(userID)
	if err != nil {
		handleErrorWithStatus(err, ctx)
		return
	}

	ctx.Status(200)
}

func handleGetComments(ctx *gin.Context) {
	userID := data.MustGetActingUserID(ctx)
	taskID, ok := ctx.Params.Get("taskID")
	if !ok {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	comments, err := taskRepo.GetComments(userID, taskID)
	if err != nil {
		handleErrorWithStatus(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

func handleCreateComment(ctx *gin.Context) {
	userID := data.MustGetActingUserID(ctx)
	createData := data.CommentCreate{}
	err := ctx.BindJSON(&createData)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	commentVM, err := taskRepo.AddComment(userID, &createData)
	if err != nil {
		handleErrorWithStatus(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, commentVM)
}

func handleUpdateComment(ctx *gin.Context) {
	userID := data.MustGetActingUserID(ctx)
	updateData := data.CommentUpdate{}
	err := ctx.BindJSON(&updateData)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	outputData, err := taskRepo.UpdateComment(userID, &updateData)
	if err != nil {
		handleErrorWithStatus(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, outputData)
}

func handleDeleteComment(ctx *gin.Context) {
	userID := data.MustGetActingUserID(ctx)
	deleteData := struct {
		CommentID string `json:"commentID"`
	}{}

	err := ctx.BindJSON(&deleteData)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = taskRepo.DeleteComment(userID, deleteData.CommentID)
	if err != nil {
		handleErrorWithStatus(err, ctx)
		return
	}

	ctx.Status(http.StatusOK)
}
