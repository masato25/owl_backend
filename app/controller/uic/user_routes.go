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
	authapi.POST("/api/v1/user/update", UserUpdate)
	authapi.POST("/api/v1/user/cgpasswd", ChangePassword)
	authapi.GET("/api/v1/user/current", UserInfo)
}
