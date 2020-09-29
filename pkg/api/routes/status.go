package routes

import (
	"net/http"

	"github.com/ataboo/go-natm/pkg/api/data"
	"github.com/gin-gonic/gin"
)

func registerStatusRoutes(e *gin.RouterGroup) {
	g := e.Group("/statuses")

	g.POST("/", handleCreateStatus)
	g.POST("/archive", handleArchiveStatus)
	g.POST("/stepOrdinal", handleStepOrdinal)
	// g.PUT("/update", handleUpdate)
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
		StatusID string `json:"status_id"`
	}{}

	err := ctx.BindJSON(&archiveData)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = statusRepo.Archive(archiveData.StatusID, userID)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
}

func handleStepOrdinal(ctx *gin.Context) {
	userID := data.MustGetActingUserID(ctx)

	archiveData := struct {
		StatusID string `json:"status_id"`
		Step     int    `json:"step"`
	}{}

	err := ctx.BindJSON(&archiveData)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = statusRepo.StepOrdinal(archiveData.StatusID, userID, archiveData.Step)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
}
