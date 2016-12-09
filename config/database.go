package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type DBPool struct {
	Falcon *gorm.DB
	Graph  *gorm.DB
	Uic    *gorm.DB
}

var (
	dbp DBPool
)

func Con() DBPool {
	return dbp
}

func SetLogLevel(loggerlevel bool) {
	dbp.Uic.LogMode(loggerlevel)
	dbp.Graph.LogMode(loggerlevel)
	dbp.Graph.LogMode(loggerlevel)
}
func InitDB(loggerlevel bool) (err error) {
	var p *sql.DB
	portal, err := gorm.Open("mysql", viper.GetString("db.faclon_portal"))
	portal.Dialect().SetDB(p)
	if err != nil {
		return
	}
	dbp.Falcon = portal

	var g *sql.DB
	graphd, err := gorm.Open("mysql", viper.GetString("db.graph"))
	graphd.Dialect().SetDB(g)
	if err != nil {
		return
	}
	dbp.Graph = graphd

	var u *sql.DB
	uicd, err := gorm.Open("mysql", viper.GetString("db.uic"))
	uicd.Dialect().SetDB(u)
	if err != nil {
		return
	}
	dbp.Uic = uicd

	SetLogLevel(loggerlevel)
	return
}

func CloseDB() (err error) {
	err = dbp.Falcon.Close()
	if err != nil {
		return
	}
	err = dbp.Graph.Close()
	if err != nil {
		return
	}
	err = dbp.Uic.Close()
	if err != nil {
		return
	}
	return
}
