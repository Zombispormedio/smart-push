package config

import (
	"os"
    
	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func Middleware(e *echo.Echo) {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

}

func Listen(e *echo.Echo) {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5065"
	}

	log.WithFields(log.Fields{
		"port": port,
	}).Info("Connected")
    
    
	e.Run(standard.New(":" + port))

}
