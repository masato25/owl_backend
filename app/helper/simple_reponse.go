package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func ErrorRepose(c *gin.Context, wcode int, err interface{}) (werror error) {
	switch err.(type) {
	case error:
		c.JSON(wcode, gin.H{
			"error": err.(error).Error(),
		})
	case string:
		c.JSON(wcode, gin.H{
			"error": err.(string),
		})
	default:
		werror = errors.New("type error")
	}
	return
}

func OkRepose(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"message": msg,
	})
	return
}
