package utils

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    400,
		"message": message,
		"status":  "Bad Request",
	})
}

func InternalServerError(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
		"code":    500,
		"message": message,
		"status":  "Internal Server Error",
	})
}
