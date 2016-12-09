package uic

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/masato25/owl_backend/app/helper"
	"github.com/masato25/owl_backend/app/model/uic"
	"github.com/masato25/owl_backend/app/utils"
)

func CreateUser(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	password := c.DefaultQuery("password", "")
	cnname := c.DefaultQuery("cnname", "")
	email := c.DefaultQuery("email", "")
	phone := c.DefaultQuery("phone", "")
	im := c.DefaultQuery("im", "")
	qq := c.DefaultQuery("qq", "")

	switch {
	case name == "" || password == "":
		c.JSON(400, gin.H{
			"error": "name or password can not be blank",
		})
		return
	case utils.HasDangerousCharacters(cnname):
		c.JSON(400, gin.H{
			"error": "name pattern is invalid",
		})
		return
	}

	var user uic.User
	db.Uic.Table("user").Where("name = ?", name).Scan(&user)
	if user.ID != 0 {
		c.JSON(400, gin.H{
			"error": "name is already existing",
		})
	}
	password = utils.HashIt(password)
	user = uic.User{
		Name:   name,
		Passwd: password,
		Cnname: cnname,
		Email:  email,
		Phone:  phone,
		IM:     im,
		QQ:     qq,
	}
	dt := db.Uic.Table("user").Create(&user)
	if dt.Error != nil {
		c.JSON(400, gin.H{
			"error": dt.Error.Error(),
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

func UserUpdate(c *gin.Context) {
	cnname := c.DefaultQuery("cnname", "")
	email := c.DefaultQuery("email", "")
	phone := c.DefaultQuery("phone", "")
	im := c.DefaultQuery("im", "")
	qq := c.DefaultQuery("qq", "")
	websession, _ := helper.GetSession(c)
	user := uic.User{
		Name:   websession.Name,
		Cnname: cnname,
		Email:  email,
		Phone:  phone,
		IM:     im,
		QQ:     qq,
	}
	dt := db.Uic.Table("user").Update(&user)
	if dt.Error != nil {
		c.JSON(400, gin.H{
			"error": dt.Error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "user update success!",
	})
	return
}

func ChangePassword(c *gin.Context) {
	oldPassword := c.DefaultQuery("old_password", "")
	newPassword := c.DefaultQuery("new_password", "")
	websession, _ := helper.GetSession(c)
	user := uic.User{Name: websession.Name}
	dt := db.Uic.Where(&user).Scan(&user)
	switch {
	case dt.Error != nil:
		c.JSON(400, gin.H{
			"error": dt.Error.Error(),
		})
		return
	case user.Passwd != utils.HashIt(oldPassword):
		c.JSON(400, gin.H{
			"error": "oldPassword is not match current one",
		})
		return
	}
	user.Passwd = utils.HashIt(newPassword)
	dt = db.Uic.Save(&user)
	if dt.Error != nil {
		c.JSON(400, gin.H{
			"error": dt.Error.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "password updated!",
	})
	return
}

func UserInfo(c *gin.Context) {
	websession, _ := helper.GetSession(c)
	user := uic.User{Name: websession.Name}
	dt := db.Uic.Where(&user).Scan(&user)
	if dt.Error != nil {
		c.JSON(400, gin.H{
			"error": dt.Error.Error(),
		})
		return
	}

	c.JSON(200, user)
	return
}
