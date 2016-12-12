package utils

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	h "github.com/masato25/owl_backend/app/helper"
	"github.com/spf13/viper"
)

func AuthSessionMidd(c *gin.Context) {
	auth, err := h.SessionChecking(c)
	if !viper.GetBool("skip_auth") {
		if err != nil || auth != true {
			log.Debugf("error: %v, auth: %v", err.Error(), auth)
			c.Set("auth", auth)
			h.JSONR(c, http.StatusUnauthorized, err)
			c.Abort()
			return
		}
	}
	c.Set("auth", auth)
}
