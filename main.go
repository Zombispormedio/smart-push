package main

import (
	"github.com/Zombispormedio/smart-push/config"
	"github.com/Zombispormedio/smart-push/router"
	"github.com/labstack/echo"
)

func main() {

	server := echo.New()

	config.Middleware(server)

	router.Use(server)

	config.Listen(server)

}
