package falcon_portal

import (
	con "github.com/masato25/owl_backend/config"
)

// +----------------+------------------+------+-----+-------------------+-----------------------------+
// | Field          | Type             | Null | Key | Default           | Extra                       |
// +----------------+------------------+------+-----+-------------------+-----------------------------+
// | id             | int(11)          | NO   | PRI | NULL              | auto_increment              |
// | hostname       | varchar(255)     | NO   | UNI |                   |                             |
// | ip             | varchar(16)      | NO   |     |                   |                             |
// | agent_version  | varchar(16)      | NO   |     |                   |                             |
// | plugin_version | varchar(128)     | NO   |     |                   |                             |
// | maintain_begin | int(10) unsigned | NO   |     | 0                 |                             |
// | maintain_end   | int(10) unsigned | NO   |     | 0                 |                             |
// | update_at      | timestamp        | NO   |     | CURRENT_TIMESTAMP | on update CURRENT_TIMESTAMP |
// +----------------+------------------+------+-----+-------------------+-----------------------------+

type Host struct {
	ID            int64  `json:"id" gorm:"column:id"`
	Hostname      string `json:"hostname" gorm:"column:hostname"`
	Ip            string `json:"ip" gorm:"column:ip"`
	AgentVersion  string `json:"agent_version"  gorm:"column:agent_version"`
	PluginVersion string `json:"plugin_version"  gorm:"column:plugin_version"`
	MaintainBegin uint16 `json:"maintain_begin"  gorm:"column:maintain_begin"`
	MaintainEnd   uint16 `json:"maintain_end"  gorm:"column:maintain_end"`
}

func (this Host) TableName() string {
	return "host"
}

func (this Host) Existing() (int64, bool) {
	db := con.Con()
	db.Falcon.Table(this.TableName()).Where("hostname = ?", this.Hostname).Scan(&this)
	if this.ID != 0 {
		return this.ID, true
	} else {
		return 0, false
	}
}
