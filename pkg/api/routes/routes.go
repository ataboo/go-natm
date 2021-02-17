package routes

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/ataboo/go-natm/pkg/common"
	"github.com/ataboo/go-natm/pkg/storage"
	"github.com/gin-gonic/gin"
)

var logger = log.New(os.Stdout, "go-natm", 0)
var ctx context.Context
var projectRepo *storage.ProjectRepository
var statusRepo *storage.StatusRepository
var taskRepo *storage.TaskRepository
var projectAssociationRepo *storage.ProjectAssociationRepository

func RegisterRoutes(
	e *gin.RouterGroup,
	pr *storage.ProjectRepository,
	sr *storage.StatusRepository,
	tr *storage.TaskRepository,
	par *storage.ProjectAssociationRepository,
) {
	projectRepo = pr
	statusRepo = sr
	taskRepo = tr
	projectAssociationRepo = par

	registerProjectRoutes(e)
	registerStatusRoutes(e)
	registerTaskRoutes(e)
	registerProjectAssociationRoutes(e)
	registerTaskCommentRoutes(e)
}

func handleErrorWithStatus(err error, c *gin.Context) {
	if err == nil {
		return
	}

	statusErr, ok := err.(*common.ErrorWithStatus)
	if !ok {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.AbortWithStatus(statusErr.Code)
	}
}
