package data

import (
	"log"

	"github.com/gin-gonic/gin"
)

func MustGetActingUserID(c *gin.Context) string {
	id, ok := c.Get("acting_user_id")
	if !ok {
		log.Fatal("No acting user")
	}

	return id.(string)
}
