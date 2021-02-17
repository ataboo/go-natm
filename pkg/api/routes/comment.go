package routes

import (
	"net/http"

	"github.com/ataboo/go-natm/pkg/api/data"
	"github.com/gin-gonic/gin"
)

func registerTaskCommentRoutes(e *gin.RouterGroup) {
	g := e.Group("/taskcomment")

	g.GET("/:taskID", handleGetComments)
	g.POST("/create", handleCreateComment)
	g.POST("/delete", handleDeleteComment)
	g.POST("/update", handleUpdateComment)
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
	deleteData := data.CommentDelete{}

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
