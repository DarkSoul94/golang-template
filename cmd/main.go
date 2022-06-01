package main

import (
	"log"
	"os"
	"os/signal"

	server "github.com/DarkSoul94/golang-template/cmd/httpserver"
	"github.com/DarkSoul94/golang-template/pkg/config"
	"github.com/DarkSoul94/golang-template/pkg/logger"
	"github.com/spf13/viper"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatal(err)
	}
	logger.InitLogger()
	
	apphttp := server.NewApp()
	apphttp.Run(viper.GetString("app.http_port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	apphttp.Stop()
}
