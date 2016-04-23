package router

import (
	"github.com/labstack/echo"
	"github.com/Zombispormedio/smart-push/lib/response"
)

func Use(e *echo.Echo) {
	
	e.GET("/", func(c echo.Context) error {
		return response.Success(c, "Works Perfectly")
	})
	
	
	ConfigRouter := e.Group("/config")
	
	ConfigRoutes(ConfigRouter)
	
	SensorGridRouter:=e.Group("/sensor_grid")
	
	SensorGridHTTPRoutes(SensorGridRouter)
	
	SensorGridWebSocketRoutes(SensorGridRouter)

}
