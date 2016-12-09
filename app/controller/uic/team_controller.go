package uic

import (
	"github.com/gin-gonic/gin"
	h "github.com/masato25/owl_backend/app/helper"
	"github.com/masato25/owl_backend/app/model/uic"
	"github.com/masato25/owl_backend/app/utils"
)

//this is root action pending
func Teams(c *gin.Context) {

}

func CreateTeam(c *gin.Context) {
	//team_name is uniq column on db, so need check existing
	team_name := c.DefaultQuery("team_name", "")
	if team_name == "" || utils.HasDangerousCharacters(team_name) {
		h.ErrorRepose(c, 400, "team_name is empty or format not vaild")
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
