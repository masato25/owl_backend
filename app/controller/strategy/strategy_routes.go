package strategy

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
	strr := r.Group("/api/v1/strategy")
	strr.Use(utils.AuthSessionMidd)
	strr.GET("", GetStrategy)
	strr.POST("", CreateStrategy)
	met := r.Group("/api/v1/metric")
	met.Use(utils.AuthSessionMidd)
	met.GET("query", MetricQuery)
}
