package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/masato25/owl_backend/app/controller"
	"github.com/masato25/owl_backend/config"
	yaag_gin "github.com/masato25/yaag/gin"
	"github.com/masato25/yaag/yaag"
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
		yaag.Init(&yaag.Config{On: true, DocTitle: "Gin", DocPath: viper.GetString("gen_doc_path"), BaseUrls: map[string]string{"Production": "/api/v1", "Staging": "/api/v1"}})
		routes.Use(yaag_gin.Document())
	}
	controller.StartGin(viper.GetString("web_port"), routes)
}
