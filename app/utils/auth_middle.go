package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	h "github.com/masato25/owl_backend/app/helper"
)

func AuthSessionMidd(c *gin.Context) {
	auth, err := h.SessionChecking(c)
	if err != nil || auth != true {
		c.Set("auth", auth)
		h.JSONR(c, http.StatusUnauthorized, err)
		c.Abort()
		return
	}
	c.Set("auth", auth)
}
