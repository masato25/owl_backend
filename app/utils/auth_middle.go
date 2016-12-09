package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/masato25/owl_backend/app/helper"
)

func AuthSessionMidd(c *gin.Context) {
	auth, err := helper.SessionChecking(c)
	if err != nil || auth != true {
		c.Set("auth", auth)
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}
	c.Set("auth", auth)
}
