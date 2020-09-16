package routes

import (
	"net/http"
	"strconv"

	"github.com/ataboo/go-natm/pkg/api/data"
	"github.com/ataboo/go-natm/pkg/common"
	"github.com/gin-gonic/gin"
)

func registerProjectRoutes(e *gin.RouterGroup) {
	g := e.Group("/projects")

	g.GET("/", handleGetProjects)
	g.GET("/:projectID", handleGetProject)
	g.POST("/", handleCreateProject)
	g.POST("/archive/", handleArchiveProject)
	// g.PUT("/:projectID", handleUpdate)
	// g.DELETE("/:projectID", handleDelete)
}

func handleGetProjects(c *gin.Context) {
	userID := data.MustGetActingUserID(c)
	associations, err := projectRepo.ListAssociated(userID)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		logger.Println("Failed to get projects", err.Error())
		return
	}

	projDatas := make([]data.ProjectGrid, 0)

	for _, assoc := range associations {
		if !assoc.R.Project.Active {
			continue
		}

		projDatas = append(projDatas, data.ProjectGrid{
			ID:              assoc.R.Project.ID,
			AssociationType: assoc.Association,
			Abbreviation:    assoc.R.Project.Abbreviation,
			Description:     assoc.R.Project.Description,
			Name:            assoc.R.Project.Name,
			LastUpdated:     assoc.R.Project.UpdatedAt.Unix(),
		})
	}

	c.JSON(200, projDatas)
}

func handleGetProject(c *gin.Context) {
	userID := data.MustGetActingUserID(c)
	projectID := c.Param("projectID")

	association, err := projectRepo.Find(projectID, userID)
	if err != nil || association == nil {
		c.AbortWithStatus(http.StatusNotFound)
		logger.Println("Failed to find proj with id", projectID)
		return
	}

	taskStatuses := association.R.Project.R.TaskStatuses
	statusDatas := make([]data.StatusRead, len(taskStatuses))
	taskDatas := make([]data.TaskGrid, 0)
	for i, s := range taskStatuses {
		statusTasks := make([]data.TaskGrid, len(s.R.Tasks))
		for j, t := range s.R.Tasks {
			statusTasks[j] = data.TaskGrid{
				Description: t.Description,
				ID:          t.ID,
				Identifier:  association.R.Project.Abbreviation + "-" + strconv.Itoa(t.Number),
				Timing: data.TimingGrid{
					Estimate: "",
					Current:  "",
				},
				StatusID: t.TaskStatusID,
				Type:     t.TaskType,
				Title:    t.Title,
			}
		}

		taskDatas = append(taskDatas, statusTasks...)

		statusDatas[i] = data.StatusRead{
			ID:   s.ID,
			Name: s.Name,
		}
	}

	projData := data.ProjectDetail{
		AssociationType: association.Association,
		ID:              association.R.Project.ID,
		Abbreviation:    association.R.Project.Abbreviation,
		Description:     association.R.Project.Description,
		Name:            association.R.Project.Name,
		Statuses:        statusDatas,
		Tasks:           taskDatas,
	}

	c.JSON(http.StatusOK, projData)
}

func handleCreateProject(c *gin.Context) {
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

func handleArchiveProject(c *gin.Context) {
	userID := data.MustGetActingUserID(c)

	archiveData := data.ProjectArchive{}
	err := c.BindJSON(&archiveData)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = projectRepo.Archive(archiveData.ProjectID, userID)
	if err != nil {
		statusErr, ok := err.(*common.ErrorWithStatus)
		if !ok {
			c.AbortWithError(http.StatusInternalServerError, err)
		} else {
			c.AbortWithStatus(statusErr.Code)
		}

		return
	}

	c.Status(200)
}
