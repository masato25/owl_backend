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

type APICreateTeamInput struct {
	Name    string  `json:"team_name" binding:"required"`
	Resume  string  `json:"resume"`
	UserIDs []int64 `json:"users"`
}

func CreateTeam(c *gin.Context) {
	var cteam APICreateTeamInput
	err := c.Bind(&cteam)
	//team_name is uniq column on db, so need check existing
	// team_name := c.DefaultQuery("team_name", "")
	if err != nil {
		h.JSONR(c, badstatus, err)
		return
	}
	user, err := h.GetUser(c)
	if err != nil {
		h.JSONR(c, badstatus, err)
		return
	} else if user.ID == 0 {
		h.JSONR(c, badstatus, "not found this user")
		return
	}
	team := uic.Team{
		Name:    cteam.Name,
		Resume:  cteam.Resume,
		Creator: user.ID,
	}
	dt := db.Uic.Save(&team)
	if dt.Error != nil {
		h.JSONR(c, badstatus, dt.Error)
		return
	}
	if len(cteam.UserIDs) > 0 {
		rel_team_user := make([]uic.RelTeamUser, len(cteam.UserIDs))
		for indx, r := range rel_team_user {
			r.Tid = team.ID
			r.Uid = cteam.UserIDs[indx]
		}
		dt := db.Uic.Table("rel_team_user").Save(rel_team_user)
		if dt.Error != nil {
			h.JSONR(c, badstatus, dt.Error)
			return
		}
	}
	h.JSONR(c, "team created!")
	return
}

type APIUpdateTeamInput struct {
	ID      int64   `json:"team_id" binding:"required"`
	Resume  string  `json:"resume"`
	UserIDs []int64 `json:"users"`
}

func UpdateTeam(c *gin.Context) {
	var cteam APIUpdateTeamInput
	err := c.Bind(&cteam)
	if err != nil {
		h.JSONR(c, badstatus, err)
		return
	}
	user, err := h.GetUser(c)
	dt := db.Uic.Table("team")
	if err != nil {
		h.JSONR(c, badstatus, err)
		return
	} else if user.IsAdmin() {
		dt = dt.Where("id = ?", cteam.ID)
	} else {
		dt = dt.Where("creator = ? AND id = ?", user.ID, cteam.ID)
	}
	dt = dt.Update(&uic.Team{Resume: cteam.Resume})
	if dt.Error != nil {
		h.JSONR(c, badstatus, err)
		return
	} else {
		if len(cteam.UserIDs) > 0 {
			rel_team_user := make([]uic.RelTeamUser, len(cteam.UserIDs))
			for indx, r := range rel_team_user {
				r.Tid = cteam.ID
				r.Uid = cteam.UserIDs[indx]
			}
			dt = db.Uic.Table("rel_team_user").Save(rel_team_user)
			if dt.Error != nil {
				h.JSONR(c, badstatus, dt.Error)
				return
			}
		}
	}
	h.JSONR(c, fmt.Sprintf("team updated!, affect row: %v", dt.RowsAffected))
	return
}

type APIDeleteTeamInput struct {
	Name string `json:"team_name" binding:"required"`
}

func DeleteTeam(c *gin.Context) {
	var cteam APIDeleteTeamInput
	err := c.Bind(&cteam)
	if err != nil {
		h.JSONR(c, badstatus, err.Error())
		return
	}
	user, err := h.GetUser(c)
	if err != nil {
		h.JSONR(c, badstatus, err.Error())
		return
	}
	if user.IsAdmin() {
		dt := db.Uic.Table("team").Delete("name = ?", cteam.Name)
		err = dt.Error
	} else {
		team := uic.Team{
			Name:    cteam.Name,
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
	h.JSONR(c, fmt.Sprintf("team %s is deleted.", cteam.Name))
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
