package uic

import (
	"github.com/gin-gonic/gin"
	h "github.com/masato25/owl_backend/app/helper"
	"github.com/masato25/owl_backend/app/model/uic"
)

//this is root action pending
func Teams(c *gin.Context) {
	query := c.DefaultQuery("query", ".+")
	user, err := h.GetUser(c)
	if err != nil {
		h.ErrorRepose(c, 400, err)
		return
	}
	teams := []uic.Team{}
	if user.Role == 2 {
		dt := db.Uic.Table("team").Where("name regexp ?", query).Scan(&teams)
		err = dt.Error
	} else {
		dt := db.Uic.Table("team").Where("name regexp ? AND creator = ?", query, user.ID).Scan(&teams)
		err = dt.Error
	}
	if err != nil {
		h.ErrorRepose(c, 400, err)
		return
	}
	c.JSON(200, teams)
}

func CreateTeam(c *gin.Context) {
	//team_name is uniq column on db, so need check existing
	team_name := c.DefaultQuery("team_name", "")
	if team_name == "" {
		h.ErrorRepose(c, 400, "team_name is empty")
		return
	}
	//allow resume empty
	resume := c.DefaultQuery("resume", "")
	websession, _ := h.GetSession(c)
	user := uic.User{
		Name: websession.Name,
	}
	dt := db.Uic.Where(&user).Find(&user)
	if user.ID == 0 {
		h.ErrorRepose(c, 400, "not found this user")
		return
	} else if dt.Error != nil {
		h.ErrorRepose(c, 400, dt.Error)
		return
	}
	team := uic.Team{
		Name:    team_name,
		Resume:  resume,
		Creator: user.ID,
	}
	dt = db.Uic.Save(&team)
	if dt.Error != nil {
		h.ErrorRepose(c, 400, dt.Error)
		return
	}
	h.OkRepose(c, "team created!")
}
