package router

import (
	"github.com/Zombispormedio/smart-push/controllers"
	"github.com/Zombispormedio/smart-push/lib/response"
    "github.com/Zombispormedio/smart-push/middleware"
	"github.com/labstack/echo"
)

func ConfigRoutes(router *echo.Group) {
    
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

	});
	
	
	router.GET("/pushover", func(c echo.Context) error {
        var Error error    
        
		ControllerError := controllers.PushOver()

		if 	ControllerError == nil {
			Error= response.Success(c, "Push Over Completed")
		} else {
			Error= response.ExpectFail(c,  ControllerError.Error())
		}
        
        return Error

	})
	
	router.GET("/clean", func(c echo.Context) error {
        var Error error    
        
		ControllerError := controllers.Clean()

		if 	ControllerError == nil {
			Error= response.Success(c, "Clean Completed")
		} else {
			Error= response.ExpectFail(c,  ControllerError.Error())
		}
        
        return Error

	})

}
