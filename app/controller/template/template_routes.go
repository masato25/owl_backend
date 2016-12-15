package template

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
	tmpr := r.Group("/api/v1/template")
	tmpr.Use(utils.AuthSessionMidd)
	tmpr.GET("", GetTemplates)
	tmpr.POST("", CreateTemplate)
	tmpr.GET("/:tpl_id", GetATemplate)
	tmpr.PUT("", UpdateTemplate)
	tmpr.DELETE("/:tpl_id", DeleteTemplate)
	tmpr.POST("/action", CreateActionToTmplate)
	tmpr.PUT("/action", UpdateActionToTmplate)
}
