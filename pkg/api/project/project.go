package project

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/ataboo/go-natm/v4/pkg/api/data"
	"github.com/ataboo/go-natm/v4/pkg/storage"
	"github.com/gin-gonic/gin"
)

var projectRepo *storage.ProjectRepository
var ctx context.Context
var logger = log.New(os.Stdout, "go-natm", 0)

func RegisterRoutes(e *gin.RouterGroup, pr *storage.ProjectRepository) {
	projectRepo = pr
	ctx = context.Background()

	g := e.Group("/projects")

	g.GET("/", handleGetList)
	g.GET("/:projectID", handleGet)
	g.POST("/", handleCreate)
	// g.PUT("/:projectID", handleUpdate)
	// g.DELETE("/:projectID", handleDelete)
}

func handleGetList(c *gin.Context) {
	userID := data.MustGetActingUserID(c)
	associations, err := projectRepo.ListAssociated(userID)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		logger.Println("Failed to get projects", err.Error())
		return
	}

	projDatas := make([]data.ProjectGrid, len(associations))

	for i, assoc := range associations {
		projDatas[i] = data.ProjectGrid{
			ID:              assoc.R.Project.ID,
			AssociationType: assoc.Association,
			Abbreviation:    assoc.R.Project.Abbreviation,
			Description:     assoc.R.Project.Description,
			Name:            assoc.R.Project.Name,
		}
	}

	c.JSON(200, projDatas)
}

func handleGet(c *gin.Context) {
	userID := data.MustGetActingUserID(c)
	projectID, ok := c.Params.Get("projectID")
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	association, err := projectRepo.Find(projectID, userID)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		logger.Println("Failed to find proj with id", projectID)
		return
	}

	taskModels := association.R.Project.R.Tasks
	tasks := make([]data.TaskGrid, len(taskModels))
	for i, t := range taskModels {
		tasks[i] = data.TaskGrid{
			ID:         t.ID,
			Identifier: t.ID,
			Name:       t.Title,
			StatusID:   t.TaskStatusID,
			Type:       t.TaskType,
			Timing: data.TimingGrid{
				Current:  "TODO query",
				Estimate: "TODO query",
			},
		}
	}

	projData := data.ProjectDetail{
		AssociationType: association.Association,
		ID:              association.R.Project.ID,
		Abbreviation:    association.R.Project.Abbreviation,
		Description:     association.R.Project.Description,
		Name:            association.R.Project.Name,
		Tasks:           tasks,
	}

	c.JSON(http.StatusOK, projData)
}

func handleCreate(c *gin.Context) {
	userID := data.MustGetActingUserID(c)

	createData := data.ProjectCreate{}
	err := c.BindJSON(&createData)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = projectRepo.Create(&createData, userID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

}
