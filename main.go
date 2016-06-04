package main

import (
	"os"
	"os/signal"

	"github.com/Zombispormedio/smart-push/config"
	"github.com/Zombispormedio/smart-push/lib/mosquito"
	"github.com/Zombispormedio/smart-push/router"
	"github.com/labstack/echo"
)

func main() {

	server := echo.New()

	config.Middleware(server)

	router.Use(server)

	if os.Getenv("IS_MQTT") == "" {
		subscriber := mosquito.New(router.Mosquito)

		sigc := make(chan os.Signal, 1)

		signal.Notify(sigc, os.Interrupt, os.Kill)

		go config.Listen(server)

		go subscriber.Run()

		<-sigc
		if err := subscriber.Client.Disconnect(); err != nil {
			panic(err)
		}
	}else{
		config.Listen(server)
	}

}
