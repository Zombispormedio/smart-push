package router

import (
	"github.com/Zombispormedio/smart-push/lib/response"
	"github.com/Zombispormedio/smart-push/controllers"
	"github.com/Zombispormedio/smart-push/middleware"
	"github.com/labstack/echo"
)

func SensorGridHTTPRoutes(router *echo.Group) {
	
	router.POST("", func(c echo.Context) error {
        var Error error    
		
		ControllerError:=controllers.ManageSensorData(c.Get("ClientID").(string),c.Get("body"))
		
		if 	ControllerError == nil {
			Error= response.Success(c, "Sync Perfectly")
		} else {
			Error= response.ExpectFail(c,  ControllerError.Error())
		}
             
        return Error

	}, middleware.SensorGrid, middleware.Body )

}
