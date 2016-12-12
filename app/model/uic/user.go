package uic

type User struct {
	ID     int64  `json:"id" `
	Name   string `json:"name" binding:"required"`
	Cnname string `json:"cnname" binding:"required"`
	Passwd string `json:"-" binding:"required"`
	Email  string `json:"email" binding:"required"`
	Phone  string `json:"phone"`
	IM     string `json:"im" gorm:"column:im"`
	QQ     string `json:"qq" gorm:"column:qq"`
	Role   int    `json:"role"`
}

func (this User) IsAdmin() bool {
	if this.Role == 2 || this.Role == 1 {
		return true
	}
	return false
}

func (this User) IsSuperAdmin() bool {
	if this.Role == 2 {
		return true
	}
	return false
}

type Team struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Resume  string `json:"resume"`
	Creator int64  `json:"creator"`
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

func (this User) TableName() string {
	return "user"
}

//db.SingularTable(true)
