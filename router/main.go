package router

import (
	"github.com/labstack/echo"
)

func Use(echo *echo.Echo) {

	ConfigRouter := echo.Group("/config")

	SetConfigRoutes(ConfigRouter)

}
