package main

import (
	"github.com/Zombispormedio/smart-push/config"
		"github.com/Zombispormedio/smart-push/lib/mosquito"
	"github.com/Zombispormedio/smart-push/router"
	"github.com/labstack/echo"

)

func main() {

	server := echo.New()

	config.Middleware(server)

	router.Use(server)
	
	subscriber:=mosquito.New(router.Mosquito)
	
	
	forever:=make(chan bool)
	
	go config.Listen(server)
	
	
	go subscriber.Run(forever)
	

	<-forever
}
