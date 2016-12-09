package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/masato25/owl_backend/app/controller"
	"github.com/masato25/owl_backend/config"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("cfg")
	viper.ReadInConfig()
	err := config.InitLog(viper.GetString("log_level"))
	if err != nil {
		log.Fatal(err)
	}
	config.InitDB(viper.GetBool("db.db_bug"))
	controller.StartGin()
}
