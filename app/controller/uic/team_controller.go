package uic

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	h "github.com/masato25/owl_backend/app/helper"
	"github.com/masato25/owl_backend/app/model/uic"
)

//support root as admin
func Teams(c *gin.Context) {
	query := c.DefaultQuery("query", ".+")
	user, err := h.GetUser(c)
	if err != nil {
		h.JSONR(c, 400, err)
		return
	}
	teams := []uic.Team{}
	if user.IsAdmin() {
		dt := db.Uic.Table("team").Where("name regexp ?", query).Scan(&teams)
		err = dt.Error
	} else {
		dt := db.Uic.Table("team").Where("name regexp ? AND creator = ?", query, user.ID).Scan(&teams)
		err = dt.Error
	}
	if err != nil {
		h.JSONR(c, badstatus, err)
		return
	}
	h.JSONR(c, teams)
	return
}

func CreateTeam(c *gin.Context) {
	//team_name is uniq column on db, so need check existing
	team_name := c.DefaultQuery("team_name", "")
	if team_name == "" {
		h.JSONR(c, badstatus, "team_name is empty")
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
		h.JSONR(c, badstatus, "not found this user")
		return
	} else if dt.Error != nil {
		h.JSONR(c, badstatus, dt.Error)
		return
	}
	team := uic.Team{
		Name:    team_name,
		Resume:  resume,
		Creator: user.ID,
	}
	dt = db.Uic.Save(&team)
	if dt.Error != nil {
		h.JSONR(c, badstatus, dt.Error)
		return
	}
	h.JSONR(c, "team created!")
	return
}

func DeleteTeam(c *gin.Context) {
	team_name := c.DefaultQuery("team_name", "")
	if team_name == "" {
		h.JSONR(c, badstatus, "team_name is empty")
		return
	}
	var err error
	user, _ := h.GetUser(c)
	if user.IsAdmin() {
		dt := db.Uic.Table("team").Delete("name = ?", team_name)
		err = dt.Error
	} else {
		team := uic.Team{
			Name:    team_name,
			Creator: user.ID,
		}
		dt := db.Uic.Where(&team).Find(&team)
		if team.ID == 0 {
			err = errors.New("You don't have permission")
		} else if dt.Error != nil {
			err = dt.Error
		} else {
			db.Uic.Delete(&team)
		}
	}
	if err != nil {
		h.JSONR(c, badstatus, err)
		return
	}
	h.JSONR(c, fmt.Sprintf("team %s is deleted.", team_name))
	return
}

func GetTeam(c *gin.Context) {
	team_id_str := c.Params.ByName("team_id")
	team_id, err := strconv.Atoi(team_id_str)
	if team_id == 0 {
		h.JSONR(c, badstatus, "team_id is empty")
		return
	} else if err != nil {
		h.JSONR(c, badstatus, err)
		return
	}
	team := uic.Team{ID: int64(team_id)}
	dt := db.Uic.Where(&team).Find(&team)
	if dt.Error != nil {
		h.JSONR(c, badstatus, dt.Error)
		return
	}
	h.JSONR(c, team)
	return
}

func UpdateTeam(c *gin.Context) {
	team_id_str := c.Params.ByName("team_id")
	team_id, err := strconv.Atoi(team_id_str)
	//team_name is uniq column on db, so need check existing
	team_name := c.DefaultQuery("team_name", "")
	if team_name == "" || team_id == 0 {
		h.JSONR(c, badstatus, "team_name or team_id is empty")
		return
	} else if err != nil {
		h.JSONR(c, badstatus, err.Error())
		return
	}
	//allow resume empty
	resume := c.DefaultQuery("resume", "")
	websession, _ := h.GetSession(c)
	user := uic.User{
		Name: websession.Name,
	}
	team := uic.Team{
		Name:   team_name,
		Resume: resume,
	}
	if user.IsAdmin() {
		dt := db.Uic.Table("team").Where("id = ?", team_id).Update(&team)
		err = dt.Error
	} else {
		dt := db.Uic.Table("name").Where("creator = ?", user.ID).Update(&team)
		err = dt.Error
	}
	if err != nil {
		h.JSONR(c, badstatus, err)
		return
	}
	h.JSONR(c, "team updated!")
}
