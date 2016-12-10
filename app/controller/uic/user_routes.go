package uic

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masato25/owl_backend/app/utils"
	"github.com/masato25/owl_backend/config"
)

var db config.DBPool

const badstatus = http.StatusBadRequest

func Routes(r *gin.Engine) {
	db = config.Con()
	//session
	u := r.Group("/api/v1/user")
	u.GET("/auth_session", AuthSession)
	u.POST("/login", Login)
	u.GET("/logout", Logout)

	//user modify
	u.POST("/create", CreateUser)
	authapi := r.Group("/api/v1/user")
	authapi.Use(utils.AuthSessionMidd)
	authapi.GET("/current", UserInfo)
	authapi.POST("/update", UserUpdate)
	authapi.POST("/cgpasswd", ChangePassword)
	//pending
	// authapi.GET("/", AllUsers)

	//team
	authapi_team := r.Group("/api/v1/team")
	authapi_team.Use(utils.AuthSessionMidd)
	authapi_team.GET("/", Teams)
	authapi_team.POST("/", CreateTeam)
	authapi_team.PUT("/:team_id", UpdateTeam)
	authapi_team.GET("/:team_id", GetTeam)
	authapi_team.DELETE("/:team_id", DeleteTeam)
}
