package uic

import "time"

type User struct {
	ID      int64     `json:"id"`
	Name    string    `json:"name"`
	Cnname  string    `json:"cnname"`
	Passwd  string    `json:"-"`
	Email   string    `json:"email"`
	Phone   string    `json:"phone"`
	IM      string    `json:"im" gorm:"column:im"`
	QQ      string    `json:"qq" gorm:"column:qq"`
	Role    int       `json:"role"`
	Created time.Time `json:"-"`
}

type Team struct {
	ID      int64     `json:"id"`
	Name    string    `json:"name"`
	Resume  string    `json:"resume"`
	Creator int64     `json:"creator"`
	Created time.Time `json:"-"`
}

type RelTeamUser struct {
	ID  int64
	Tid int64
	Uid int64
}

type Session struct {
	ID      int64
	Uid     int64
	Sig     string
	Expired int
}

func (this Session) TableName() string {
	return "session"
}

//db.SingularTable(true)
