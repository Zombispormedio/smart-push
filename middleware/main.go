package middleware

import ( 
    "os"
    "github.com/labstack/echo"
"github.com/Zombispormedio/smart-push/response"
)

func Task(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var Error error

		taskAuth := c.Request().Header().Get("Authorization")

		if taskAuth != "" && taskAuth==os.Getenv("SMART_TASK_SECRET") {
            
            Error=next(c)
            
		} else {
            Error=response.Forbidden(c, "No Authorization")
		}

		return Error
	}
}
