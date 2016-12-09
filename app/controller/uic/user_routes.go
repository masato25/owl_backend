package uic

import (
	"github.com/gin-gonic/gin"
	"github.com/masato25/owl_backend/app/utils"
	"github.com/masato25/owl_backend/config"
)

var db config.DBPool

func Routes(r *gin.Engine) {
	db = config.Con()
	r.GET("/api/v1/user/auth_session", AuthSession)
	r.POST("/api/v1/user/login", Login)
	r.GET("/api/v1/user/logout", Logout)
	r.POST("/api/v1/user/create", CreateUser)
	authapi := r.Group("/api/v1/user")
	authapi.Use(utils.AuthSessionMidd)
	authapi.POST("/update", UserUpdate)
	authapi.POST("/cgpasswd", ChangePassword)
	authapi.GET("/current", UserInfo)
	authapi_team := r.Group("/api/v1/team")
	authapi_team.Use(utils.AuthSessionMidd)
	authapi_team.POST("/create", CreateTeam)
}
