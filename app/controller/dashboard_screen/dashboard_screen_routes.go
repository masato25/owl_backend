package dashboard_screen

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masato25/go_email/app/utils"
	"github.com/masato25/owl_backend/config"
)

var db config.DBPool

const badstatus = http.StatusBadRequest
const expecstatus = http.StatusExpectationFailed

func Routes(r *gin.Engine) {
	db = config.Con()
	r.GET("/api/v1/dashboard/screen_all", GetScreenList)
	expr := r.Group("/api/v1/dashboard/screen")
	expr.Use(utils.AuthSessionMidd)
	expr.POST("", CreateScreen)
	expr.GET("/:sid", GetScreen)
	expr.PUT("", UpdateScreen)
	expr.DELETE("/:sid", DeleteScreen)
}
