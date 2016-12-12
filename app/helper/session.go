package helper

import (
	"errors"

	log "github.com/Sirupsen/logrus"

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
	s := c.Request.Cookies()
	log.Debugf("%v", s)
	if name == "" {
		name = c.DefaultQuery("name", "")
	}
	log.Debugf("session got name: %s", name)
	if err != nil {
		return
	}
	sig, err = c.Cookie("sig")
	if sig == "" {
		sig = c.DefaultQuery("sig", "")
	}
	log.Debugf("session got sig: %s", sig)
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

func GetUser(c *gin.Context) (user uic.User, err error) {
	db := config.Con().Uic
	websession, _ := GetSession(c)
	user = uic.User{
		Name: websession.Name,
	}
	dt := db.Where(&user).Find(&user)
	err = dt.Error
	return
}
