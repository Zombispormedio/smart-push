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

		taskAuth := c.Request().Header().Get("Authorization")

		if taskAuth != "" && taskAuth == os.Getenv("SMART_TASK_SECRET") {

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
			hostname := os.Getenv("SENSOR_STORE_HOSTNAME")
			url := hostname + "push/config/sensor_grid"

			reqBody := response.ReqSensorT{}

			reqBody.ClientID = ClientID
			reqBody.ClientSecret = ClientSecret

			resBody := &response.MixedMessageT{}

			RequestError := request.PostWithAuthorization(url, reqBody, resBody)

			if RequestError != nil {
				return response.Forbidden(c, "No Authorization")
			}

			if resBody.Status == 0 {
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
