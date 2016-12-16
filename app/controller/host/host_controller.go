package host

import (
	"fmt"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	h "github.com/masato25/owl_backend/app/helper"
	f "github.com/masato25/owl_backend/app/model/falcon_portal"
)

func GetHostGroups(c *gin.Context) {
	var hostgroups []f.HostGroup
	if dt := db.Falcon.Table("grp").Limit(3).Scan(&hostgroups); dt.Error != nil {
		h.JSONR(c, expecstatus, dt.Error)
		return
	}
	h.JSONR(c, hostgroups)
	return
}

type APICrateHostGroup struct {
	Name string `json:"name" binding:"required"`
}

func CrateHostGroup(c *gin.Context) {
	var inputs APICrateHostGroup
	if err := c.Bind(&inputs); err != nil {
		h.JSONR(c, badstatus, err)
		return
	}
	user, _ := h.GetUser(c)
	hostgroup := f.HostGroup{Name: inputs.Name, CreateUser: user.Name, ComeFrom: 1}
	if dt := db.Falcon.Create(&hostgroup); dt.Error != nil {
		h.JSONR(c, expecstatus, dt.Error)
		return
	}
	h.JSONR(c, hostgroup)
	return
}

type APIBindHostToHostGroupInput struct {
	Hosts       []string `json:"hosts" binding:"required"`
	HostGroupID int64    `json:"hostgroup_id" binding:"required"`
}

func BindHostToHostGroup(c *gin.Context) {
	var inputs APIBindHostToHostGroupInput
	if err := c.Bind(&inputs); err != nil {
		h.JSONR(c, badstatus, err)
		return
	}
	user, _ := h.GetUser(c)
	hostgroup := f.HostGroup{ID: inputs.HostGroupID}
	if dt := db.Falcon.Find(&hostgroup); dt.Error != nil {
		h.JSONR(c, expecstatus, dt.Error)
		return
	}
	if !user.IsAdmin() && hostgroup.CreateUser != user.Name {
		h.JSONR(c, expecstatus, "You don't have permission.")
		return
	}
	tx := db.Falcon.Begin()
	if dt := tx.Where("grp_id = ?", hostgroup.ID).Delete(&f.GrpHost{}); dt.Error != nil {
		h.JSONR(c, expecstatus, fmt.Sprintf("delete grp_host got error: %v", dt.Error))
		dt.Rollback()
		return
	}
	var ids []int64
	for _, host := range inputs.Hosts {
		ahost := f.Host{Hostname: host}
		var id int64
		var ok bool
		if id, ok = ahost.Existing(); ok {
			ids = append(ids, id)
		} else {
			if dt := tx.Save(&ahost); dt.Error != nil {
				h.JSONR(c, expecstatus, dt.Error)
				tx.Rollback()
				return
			}
			id = ahost.ID
			ids = append(ids, id)
		}
		if dt := tx.Debug().Create(&f.GrpHost{GrpID: hostgroup.ID, HostID: id}); dt.Error != nil {
			h.JSONR(c, expecstatus, fmt.Sprintf("create grphost got error: %s , grp_id: %v, host_id: %v", dt.Error, hostgroup.ID, id))
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	h.JSONR(c, fmt.Sprintf("%v bind to hostgroup: %v", ids, hostgroup.ID))
	return
}

type APIUnBindAHostToHostGroup struct {
	HostID      int64 `json:"host_id" binding:"required"`
	HostGroupID int64 `json:"hostgroup_id" binding:"required"`
}

func UnBindAHostToHostGroup(c *gin.Context) {
	var inputs APIUnBindAHostToHostGroup
	if err := c.Bind(&inputs); err != nil {
		h.JSONR(c, badstatus, err)
		return
	}
	user, _ := h.GetUser(c)
	hostgroup := f.HostGroup{ID: inputs.HostGroupID}
	if !user.IsAdmin() {
		if dt := db.Falcon.Find(&hostgroup); dt.Error != nil {
			h.JSONR(c, badstatus, dt.Error)
			return
		}
		if hostgroup.CreateUser == user.Name {
			h.JSONR(c, badstatus, "You don't have permission!")
			return
		}
	}
	if dt := db.Falcon.Where("grp_id = ? AND host_id = ?", inputs.HostGroupID, inputs.HostID).Delete(&f.GrpHost{}); dt.Error != nil {
		h.JSONR(c, expecstatus, dt.Error)
		return
	}
	h.JSONR(c, fmt.Sprintf("unbind host:%v of hostgroup: %v", inputs.HostID, inputs.HostGroupID))
	return
}

func DeleteHostGroup(c *gin.Context) {
	grpIDtmp := c.Params.ByName("host_group")
	if grpIDtmp == "" {
		h.JSONR(c, badstatus, "grp id is missing")
		return
	}
	grpID, err := strconv.Atoi(grpIDtmp)
	if err != nil {
		h.JSONR(c, badstatus, err)
		return
	}
	tx := db.Falcon.Begin()
	//delete hostgroup referance of grp_host table
	if dt := tx.Where("grp_id = ?", grpID).Delete(&f.GrpHost{}); dt.Error != nil {
		h.JSONR(c, expecstatus, fmt.Sprintf("delete grp_host got error: %v", dt.Error))
		dt.Rollback()
		return
	}
	//delete plugins of hostgroup
	if dt := tx.Where("grp_id = ?", grpID).Delete(&f.Plugin{}); dt.Error != nil {
		h.JSONR(c, expecstatus, fmt.Sprintf("delete plugins got error: %v", dt.Error))
		dt.Rollback()
		return
	}
	//delete aggreators of hostgroup
	if dt := tx.Where("grp_id = ?", grpID).Delete(&f.Cluster{}); dt.Error != nil {
		h.JSONR(c, expecstatus, fmt.Sprintf("delete aggreators got error: %v", dt.Error))
		dt.Rollback()
		return
	}
	//finally delete hostgroup
	if dt := tx.Delete(&f.HostGroup{ID: int64(grpID)}); dt.Error != nil {
		h.JSONR(c, expecstatus, dt.Error)
		tx.Rollback()
		return
	}
	tx.Commit()
	h.JSONR(c, fmt.Sprintf("hostgroup:%v has been deleted", grpID))
	return
}

func GetHostGroup(c *gin.Context) {
	grpIDtmp := c.Params.ByName("host_group")
	if grpIDtmp == "" {
		h.JSONR(c, badstatus, "grp id is missing")
		return
	}
	grpID, err := strconv.Atoi(grpIDtmp)
	if err != nil {
		log.Debugf("grpIDtmp: %v", grpIDtmp)
		h.JSONR(c, badstatus, err)
		return
	}
	hostgroup := f.HostGroup{ID: int64(grpID)}
	if dt := db.Falcon.Find(&hostgroup); dt.Error != nil {
		h.JSONR(c, expecstatus, dt.Error)
		return
	}
	hosts := []f.Host{}
	grpHosts := []f.GrpHost{}
	if dt := db.Falcon.Where("grp_id = ?", grpID).Find(&grpHosts); dt.Error != nil {
		h.JSONR(c, expecstatus, dt.Error)
		return
	}
	for _, grph := range grpHosts {
		var host f.Host
		db.Falcon.Find(&host, grph.HostID)
		if host.ID != 0 {
			hosts = append(hosts, host)
		}
	}
	h.JSONR(c, map[string]interface{}{
		"hostgroup": hostgroup,
		"hosts":     hosts,
	})
	return
}
