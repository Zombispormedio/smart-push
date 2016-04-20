package router

import (
	"github.com/Zombispormedio/smart-push/controllers"
	"github.com/Zombispormedio/smart-push/response"
    "github.com/Zombispormedio/smart-push/middleware"
	"github.com/labstack/echo"
)

func SetConfigRoutes(router *echo.Group) {
    
    router.Use(middleware.Task)
    
	router.GET("/credentials", func(c echo.Context) error {
        var Error error    
        
		ControllerError := controllers.RefreshCredentials()

		if 	ControllerError == nil {
			Error= response.Success(c, "Refreshed Perfectly")
		} else {
			Error= response.ExpectFail(c,  ControllerError.Error())
		}
        
        return Error

	})

}
