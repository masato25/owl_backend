package host

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masato25/owl_backend/app/utils"
	"github.com/masato25/owl_backend/config"
)

var db config.DBPool

const badstatus = http.StatusBadRequest
const expecstatus = http.StatusExpectationFailed

func Routes(r *gin.Engine) {
	db = config.Con()
	hostr := r.Group("/api/v1/hostgroup")
	hostr.Use(utils.AuthSessionMidd)
	hostr.GET("", GetHostGroups)
	hostr.POST("", CrateHostGroup)
	hostr.POST("/host", BindHostToHostGroup)
	hostr.PUT("/host", UnBindAHostToHostGroup)
	hostr.GET("/:host_group", GetHostGroup)
	hostr.DELETE("/:host_group", DeleteHostGroup)
}
