package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/masato25/owl_backend/app/model/uic"
	"github.com/masato25/owl_backend/config"
)

type WebSession struct {
	Name string
	Sig  string
}

func GetSession(c *gin.Context) (session WebSession, err error) {
	var name, sig string
	name, err = c.Cookie("name")
	if name == "" {
		name = c.DefaultQuery("name", "")
	}
	if err != nil {
		return
	}
	sig, err = c.Cookie("sig")
	if sig == "" {
		sig = c.DefaultQuery("sig", "")
	}
	if err != nil {
		return
	}
	session = WebSession{name, sig}
	return
}

func SessionChecking(c *gin.Context) (auth bool, err error) {
	auth = false
	var websessio WebSession
	websessio, err = GetSession(c)
	if err != nil {
		return
	}
	db := config.Con().Uic
	var user uic.User
	db.Table("user").Where("name = ?", websessio.Name).Scan(&user)
	if user.ID == 0 {
		err = errors.New("not found this user")
		return
	}
	var session uic.Session
	db.Table("session").Where("sig = ? and uid = ?", websessio.Sig, user.ID).Scan(&session)
	if session.ID == 0 {
		err = errors.New("session not found")
		return
	} else {
		auth = true
	}
	return

}
