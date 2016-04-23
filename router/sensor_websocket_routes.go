package router

import (
	"encoding/json"
	"errors"

	log "github.com/Sirupsen/logrus"
	"github.com/Zombispormedio/smart-push/controllers"
	"github.com/Zombispormedio/smart-push/lib/request"
	"github.com/Zombispormedio/smart-push/lib/response"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"golang.org/x/net/websocket"
)

func getJSONMessage(ws *websocket.Conn, obj interface{}) error {
	var msg string

	MessageError := websocket.Message.Receive(ws, &msg)

	if MessageError != nil {
		return MessageError
	}

	return json.Unmarshal([]byte(msg), obj)

}

func Login(Receive func(interface{}) error) (error, string) {
	var Error error
	var ID string
	var LoginData response.ReqSensorT

	LoginMessageError := Receive(&LoginData)

	if LoginMessageError != nil {
		return LoginMessageError, ID
	}

	if LoginData.ClientID == "" || LoginData.ClientSecret == "" {
		return errors.New("Empty Login Form"), ID
	}

	Accepted, LoginError := request.CheckSensorGrid(LoginData)

	if LoginError != nil || !Accepted {
		return errors.New("Login Error"), ID
	}

	ID = LoginData.ClientID

	return Error, ID

}

func SendSuccessMessage(ws *websocket.Conn, message string ) error{
	
	msg:=response.MessageT{
		Status:0,
		Message:message,
	}
	
	msgjson, _:=json.Marshal(msg)
	
	return websocket.Message.Send(ws, string(msgjson))
	
}


func SensorGridWebSocketRoutes(router *echo.Group) {

	router.GET("", standard.WrapHandler(websocket.Handler(func(ws *websocket.Conn) {

		Receive := func(obj interface{}) error {
			return getJSONMessage(ws, obj)
		}

		LoginError, SensorGridID := Login(Receive)

		if LoginError == nil {
			SendSuccessMessage(ws, "Login Successfully")
			for {

				var data map[string]interface{}

				DataError := Receive(&data)

				if DataError != nil {
					log.WithFields(log.Fields{
						"error": DataError,
					}).Error("SensorGridSocketDataError")
					break
				}

				ControllerError := controllers.ManageSensorData(SensorGridID, data)

				if ControllerError != nil {
					log.WithFields(log.Fields{
						"error": ControllerError,
					}).Error("SensorGridSocketControllerError")
					break
				}
				
				SendSuccessMessage(ws, "Data Successfully")
			}

		}

		ws.Close()

	})))

}
