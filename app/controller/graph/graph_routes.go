package graph

import (
	"github.com/gin-gonic/gin"
	"github.com/masato25/owl_backend/app/utils"
	"github.com/masato25/owl_backend/config"
)

var db config.DBPool

func Routes(r *gin.Engine) {
	db = config.Con()
	authapi := r.Group("/api/v1")
	authapi.Use(utils.AuthSessionMidd)
	authapi.GET("/graph/endpoint", EndpointRegexpQuery)
	authapi.GET("/graph/endpoint_counter", EndpointCounterRegexpQuery)
}
