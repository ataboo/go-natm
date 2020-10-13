package routes

import (
	"net/http"

	"github.com/ataboo/go-natm/pkg/api/data"
	"github.com/gin-gonic/gin"
)

func registerProjectAssociationRoutes(e *gin.RouterGroup) {
	g := e.Group("/projectassociation/")

	g.POST("/", handleCreateProjectAssociation)
	g.POST("/delete", handleDeleteProjectAssociation)
	g.POST("/update", handleUpdateProjectAssociation)
}

func handleCreateProjectAssociation(c *gin.Context) {
	userID := data.MustGetActingUserID(c)

	createData := &data.ProjectAssociationCreate{}
	err := c.BindJSON(&createData)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = projectAssociationRepo.Create(createData, userID)
	if err != nil {
		handleErrorWithStatus(err, c)
		return
	}

	c.Status(200)
}

func handleDeleteProjectAssociation(c *gin.Context) {
	userID := data.MustGetActingUserID(c)
	deleteData := &data.ProjectAssociationDelete{}
	err := c.BindJSON(&deleteData)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = projectAssociationRepo.Delete(deleteData, userID)
	if err != nil {
		handleErrorWithStatus(err, c)
		return
	}

	c.Status(200)
}

func handleUpdateProjectAssociation(c *gin.Context) {
	userID := data.MustGetActingUserID(c)
	updateData := &data.ProjectAssociationUpdate{}
	err := c.BindJSON(&updateData)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = projectAssociationRepo.Update(updateData, userID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(200)
}
