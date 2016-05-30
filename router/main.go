package router

import (
	"github.com/Zombispormedio/smart-push/lib/redis"
	"github.com/Zombispormedio/smart-push/lib/response"
	"github.com/labstack/echo"
)

func Use(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return response.Success(c, "Works Perfectly")
	})

	e.GET("/status", func(c echo.Context) error {

		RedisError := redis.Status()

		status := struct {
			Redis bool `json:"redis_status"`
		}{}

		status.Redis = RedisError == nil

		return response.Data(c, status)
	})

	ConfigRouter := e.Group("/task")

	ConfigRoutes(ConfigRouter)

	SensorGridRouter := e.Group("/sensor_grid")

	SensorGridHTTPRoutes(SensorGridRouter)

	SensorGridWebSocketRoutes(SensorGridRouter)

	RealtimeRouter := e.Group("/realtime")

	RealtimeRoutes(RealtimeRouter)



}
