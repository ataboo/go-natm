package storage

import (
	"github.com/ataboo/go-natm/pkg/models"
	"github.com/gin-gonic/gin"
)

type SessionStore interface {
	MustGetUserSession(c *gin.Context) *models.User

	CreateSession(c *gin.Context)
}
