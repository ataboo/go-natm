package data

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MustGetActingUserID(c *gin.Context) string {
	id, ok := c.Get("acting_user_id")
	if !ok {
		log.Fatal("No acting user")
	}

	return id.(string)
}

func TaskIdentifier(abbreviation string, taskNumber int) string {
	return abbreviation + "-" + strconv.Itoa(taskNumber)
}
