package main

import (
	"fmt"
	"os"
	"os/signal"

	server "github.com/DarkSoul94/golang-template/cmd/httpserver"
	"github.com/DarkSoul94/golang-template/config"
	"github.com/DarkSoul94/golang-template/pkg/logger"
)

func main() {
	conf := config.InitConfig()
	logger.InitLogger(conf)

	apphttp := server.NewApp(conf)
	go apphttp.Run(conf)

	fmt.Println(
		fmt.Sprintf(
			"Service %s is running",
			conf.AppName,
		),
	)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	apphttp.Stop()

	fmt.Println(
		fmt.Sprintf(
			"Service %s is stopped",
			conf.AppName,
		),
	)
}
