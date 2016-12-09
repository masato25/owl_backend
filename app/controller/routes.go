package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masato25/owl_backend/app/controller/graph"
	"github.com/masato25/owl_backend/app/controller/uic"
)

func StartGin(port string, r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, I'm OWL")
	})
	graph.Routes(r)
	uic.Routes(r)
	r.Run()
}
