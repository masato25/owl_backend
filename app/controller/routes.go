package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/masato25/owl_backend/app/helper"
	"github.com/masato25/owl_backend/config"
)

var db config.DBPool

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

func StartGin() {
	r := gin.Default()
	db = config.Con()
	authapi := r.Group("/api/v1")
	authapi.Use(AuthSessionMidd)
	authapi.GET("/graph/endpoint", EndpointRegexpQuery)
	authapi.GET("/graph/endpoint_counter", EndpointCounterRegexpQuery)
	r.POST("/api/v1/user/login", Login)
	r.GET("/api/v1/user/logout", Logout)
	r.GET("/api/v1/user/auth_session", AuthSession)
	r.Run()
}
