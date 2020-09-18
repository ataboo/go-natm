package routes

import (
	"net/http"
	"strconv"

	"github.com/ataboo/go-natm/pkg/api/data"
	"github.com/gin-gonic/gin"
)

func registerProjectRoutes(e *gin.RouterGroup) {
	g := e.Group("/projects")

	g.GET("/", handleGetProjects)
	g.GET("/:projectID", handleGetProject)
	g.POST("/", handleCreateProject)
	g.POST("/archive", handleArchiveProject)
	g.POST("/setTaskOrder", handleSetTaskOrder)
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
			estimateStr := ""
			if t.Estimate.Valid {
				estimateStr = strconv.Itoa(t.Estimate.Int)
			}

			var assignee *data.UserRead = nil
			if t.R.Assignee != nil {
				assignee = &data.UserRead{
					Email: t.R.Assignee.Email,
					ID:    t.R.Assignee.ID,
					Name:  t.R.Assignee.Name,
				}
			}

			statusTasks[j] = data.TaskGrid{
				Description: t.Description,
				ID:          t.ID,
				Identifier:  association.R.Project.Abbreviation + "-" + strconv.Itoa(t.Number),
				Timing: &data.TimingGrid{
					Estimate: estimateStr,
					Current:  "",
				},
				StatusID: t.TaskStatusID,
				Type:     t.TaskType,
				Title:    t.Title,
				Ordinal:  t.Ordinal,
				Assignee: assignee,
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

	c.Status(200)
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
		handleErrorWithStatus(err, c)
		return
	}

	c.Status(200)
}

func handleSetTaskOrder(c *gin.Context) {
	userID := data.MustGetActingUserID(c)
	taskOrder := data.ProjectTaskOrder{}
	err := c.BindJSON(&taskOrder)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = projectRepo.SetTaskOrder(taskOrder, userID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.Status(200)
}
