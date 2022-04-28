package main

import (
	"log"

	"github.com/DarkSoul94/golang-template/pkg/config"
	"github.com/DarkSoul94/golang-template/pkg/logger"
	"github.com/DarkSoul94/golang-template/server"
	"github.com/spf13/viper"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatal(err)
	}
	logger.InitLogger()
	app := server.NewApp()
	app.Run(viper.GetString("app.http_port"))
}
