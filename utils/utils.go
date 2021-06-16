package utils

import (
	"github.com/gin-gonic/gin"
)

func Message(status bool, message string) (map[string]interface{}) {
	return map[string]interface{} {"status" : status, "message" : message}
}

func Respond(c *gin.Context, status int, data map[string] interface{}) {
	c.JSON(status, data)
}