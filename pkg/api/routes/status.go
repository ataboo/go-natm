package routes

import (
	"net/http"

	"github.com/ataboo/go-natm/pkg/api/data"
	"github.com/gin-gonic/gin"
)

func registerStatusRoutes(e *gin.RouterGroup) {
	g := e.Group("/statuses")

	g.POST("/", handleCreateStatus)
	g.POST("/archive/", handleArchiveStatus)
	// g.PUT("/update", handleUpdate)
	// g.PUT("/:projectID", handleUpdate)
	// g.DELETE("/:projectID", handleDelete)
}

func handleCreateStatus(ctx *gin.Context) {
	userID := data.MustGetActingUserID(ctx)

	createData := data.StatusCreate{}
	err := ctx.BindJSON(&createData)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = statusRepo.Create(&createData, userID)
	if err != nil {
		handleErrorWithStatus(err, ctx)
		return
	}

	ctx.Status(200)
}

func handleArchiveStatus(ctx *gin.Context) {
	userID := data.MustGetActingUserID(ctx)

	archiveData := struct {
		StatusID `json:"status_id"`
	}{}

	err := ctx.BindJSON(&archiveData)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = statusRepo.Archive(archiveData.StatusID, userID)
}
