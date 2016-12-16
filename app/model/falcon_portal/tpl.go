package falcon_portal

import (
	"github.com/masato25/owl_backend/app/model/uic"
	con "github.com/masato25/owl_backend/config"
)

type Template struct {
	ID         uint   `json:"id" gorm:"column:id"`
	Name       string `json:"tpl_name" gorm:"column:tpl_name"`
	ParentID   int64  `json:"parent_id" orm:"column:parent_id"`
	ActionID   int64  `json:"action_id" orm:"column:action_id"`
	CreateUser string `json:"create_user" orm:"column:create_user"`
}

func (this Template) TableName() string {
	return "tpl"
}

func (this Template) FindUserName() (name string, err error) {
	var user uic.User
	user.Name = this.CreateUser
	db := con.Con()
	dt := db.Uic.Find(&user)
	if dt.Error != nil {
		err = dt.Error
		return
	}
	name = user.Name
	return
}
