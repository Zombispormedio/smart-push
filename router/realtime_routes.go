package router

import (
	"github.com/Zombispormedio/smart-push/controllers"
	"github.com/Zombispormedio/smart-push/lib/response"
	"github.com/Zombispormedio/smart-push/middleware"
	"github.com/labstack/echo"
)



func RealtimeRoutes(router *echo.Group) {

	router.Use(middleware.Realtime)

	router.GET("/:sensor", func(c echo.Context) error {
		var Error error

		sensorID := c.Param("sensor")

		sensor := response.RealTimeData{}

		sensor.ID = sensorID

		ControllerError := controllers.GetRealtimeData(&sensor)

		if ControllerError == nil {
			Error = response.Data(c, sensor)
		} else {
			Error = response.ExpectFail(c, ControllerError.Error())
		}

		return Error

	})

}
