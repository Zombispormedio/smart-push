package middleware

import (
	"os"

	"github.com/Zombispormedio/smart-push/lib/request"
	"github.com/Zombispormedio/smart-push/lib/response"
	"github.com/labstack/echo"
)

func Task(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var Error error

		headerAuth := c.Request().Header().Get("Authorization")
		queryAuth := c.QueryParam("authorization")

		if (headerAuth != "" && headerAuth == os.Getenv("SMART_TASK_SECRET")) || (queryAuth != "" && queryAuth == os.Getenv("SMART_TASK_SECRET")) {

			Error = next(c)

		} else {

			Error = response.Forbidden(c, "No Authorization")

		}

		return Error
	}
}

func SensorGrid(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var Error error

		ClientID := c.Request().Header().Get("ClientID")

		ClientSecret := c.Request().Header().Get("ClientSecret")

		if ClientID != "" && ClientSecret != "" {

			reqBody := response.ReqSensorT{}

			reqBody.ClientID = ClientID
			reqBody.ClientSecret = ClientSecret

			RequestAccepted, RequestError := request.CheckSensorGrid(reqBody)

			if RequestError != nil {
				return response.Forbidden(c, "No Authorization")
			}

			if RequestAccepted {
				c.Set("ClientID", ClientID)

				Error = next(c)
			} else {
				return response.Forbidden(c, "No Authorization")
			}

		} else {
			Error = response.Forbidden(c, "No Authorization: Empty Headers")
		}

		return Error
	}
}

func Body(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var Error error

		var u interface{}

		if err := c.Bind(&u); err != nil {
			return response.Forbidden(c, err.Error())
		}

		if u != nil {

			c.Set("body", u)

			Error = next(c)

		} else {
			Error = response.Forbidden(c, "No body")
		}

		return Error
	}
}


func Realtime(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var Error error

		headerAuth := c.Request().Header().Get("Authorization")
		

		if (headerAuth != "" && headerAuth == os.Getenv("OPEN_API_SECRET")) {

			Error = next(c)

		} else {

			Error = response.Forbidden(c, "No Authorization")

		}

		return Error
	}
}