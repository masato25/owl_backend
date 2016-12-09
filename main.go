package main

import (
	log "github.com/Sirupsen/logrus"
	yaag_gin "github.com/betacraft/yaag/gin"
	"github.com/betacraft/yaag/yaag"
	"github.com/gin-gonic/gin"
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
	routes := gin.Default()
	if viper.GetBool("gen_doc") {
		yaag.Init(&yaag.Config{On: true, DocTitle: "Gin", DocPath: "apidoc.html", BaseUrls: map[string]string{"Production": "/", "Staging": "/"}})
		routes.Use(yaag_gin.Document())
	}
	controller.StartGin(viper.GetString("web_port"), routes)
}
