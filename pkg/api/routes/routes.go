package routes

import (
	"context"
	"log"
	"os"

	"github.com/ataboo/go-natm/pkg/storage"
	"github.com/gin-gonic/gin"
)

var logger = log.New(os.Stdout, "go-natm", 0)
var ctx context.Context
var projectRepo *storage.ProjectRepository
var statusRepo *storage.StatusRepository
var taskRepo *storage.TaskRepository

func RegisterRoutes(e *gin.RouterGroup, pr *storage.ProjectRepository, sr *storage.StatusRepository, tr *storage.TaskRepository) {
	projectRepo = pr
	statusRepo = sr
	taskRepo = tr

	registerProjectRoutes(e)
	registerStatusRoutes(e)
	registerTaskRoutes(e)
}
