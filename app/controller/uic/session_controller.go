package uic

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/masato25/owl_backend/app/helper"
	"github.com/masato25/owl_backend/app/model/uic"
	"github.com/masato25/owl_backend/app/utils"
)

func Login(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	password := c.DefaultQuery("password", "")

	if name == "" || password == "" {
		c.JSON(400, gin.H{
			"error": "name or password is blank",
		})
		return
	}
	var user uic.User
	db.Uic.Table("user").Where("name = ?", name).Scan(&user)
	switch {
	case user.Name == "":
		c.JSON(400, gin.H{
			"error": "no such user",
		})
		return
	case user.Passwd != utils.HashIt(password):
		c.JSON(400, gin.H{
			"error": "password error",
		})
		return
	}
	var session uic.Session
	response := map[string]string{}
	s := db.Uic.Table("session").Where("uid = ?", user.ID).Scan(&session)
	if s.Error != nil && s.Error.Error() != "record not found" {
		c.JSON(400, gin.H{
			"error": s.Error.Error(),
		})
		return
	} else if session.ID == 0 {
		session.Sig = utils.GenerateUUID()
		session.Expired = int(time.Now().Unix()) + 3600*24*30
		session.Uid = user.ID
		db.Uic.Create(&session)
	}
	log.Infof("%v", session)
	response["sig"] = session.Sig
	response["name"] = user.Name
	c.JSON(200, response)
	return
}

func Logout(c *gin.Context) {
	wsession, err := helper.GetSession(c)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	var session uic.Session
	var user uic.User
	db.Uic.Table("user").Where(uic.User{Name: wsession.Name}).Scan(&user)
	db.Uic.Table("session").Where("sig = ? AND uid = ?", wsession.Sig, user.ID).Scan(&session)

	if session.ID == 0 {
		c.JSON(400, gin.H{
			"error": "not found this kind of session in database.",
		})
		return
	} else {
		r := db.Uic.Table("session").Delete(&session)
		if r.Error != nil {
			c.JSON(400, gin.H{
				"error": r.Error.Error(),
			})
		}
		c.JSON(200, gin.H{
			"message": "logout successful",
		})
	}
	return
}

func AuthSession(c *gin.Context) {
	auth, err := helper.SessionChecking(c)
	if err != nil || auth != true {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "session is vaild!",
	})
	return
}

func CreateRoot(c *gin.Context) {
	password := c.DefaultQuery("password", "")
	if password == "" {
		c.JSON(400, gin.H{
			"error": "password is empty, please check it",
		})
		return
	}
	password = utils.HashIt(password)
	user := uic.User{
		Name:   "root",
		Passwd: password,
	}
	dt := db.Uic.Table("user").Save(&user)
	if dt.Error != nil {
		c.JSON(400, gin.H{
			"error": dt.Error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "root created!",
	})
	return
}
