package main

import (
	"github.com/Zombispormedio/smart-push/config"
		"github.com/Zombispormedio/smart-push/lib/mosquito"
	"github.com/Zombispormedio/smart-push/router"
	"github.com/labstack/echo"
	"os"
	  "os/signal"

)

func main() {

	server := echo.New()

	config.Middleware(server)

	router.Use(server)
	
	subscriber:=mosquito.New(router.Mosquito)
	
	
	sigc := make(chan os.Signal, 1)

    signal.Notify(sigc, os.Interrupt, os.Kill)
	
	go config.Listen(server)
	
	
	go subscriber.Run()
	

	  <-sigc
	   if err := subscriber.Client.Disconnect(); err != nil {
        panic(err)
    }
}
