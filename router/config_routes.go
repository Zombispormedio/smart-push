package router

import(
    "github.com/labstack/echo"
    "github.com/Zombispormedio/smart-push/controllers"
    	"net/http"
)



func SetConfigRoutes(router *echo.Group){
    
    router.GET("/credentials", func (c echo.Context) error{
        
        controllers.RefreshCredentials()
        
        return c.String(http.StatusOK, "hello")
    })
    
    
    
    
}