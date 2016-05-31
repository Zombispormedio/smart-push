package router

import (
	"encoding/json"
	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/Zombispormedio/smart-push/controllers"
)

func Mosquito(data []byte) error {
	var Error error

	result := map[string]interface{}{}

	JSONError := json.Unmarshal(data, &result)

	if JSONError != nil {
		return JSONError
	}

	if result["sensor_grid"] == nil{
		return errors.New("No sensor_grid client_id")
	}
	sensorGrid := result["sensor_grid"].(string)

	log.WithFields(log.Fields{
		"client_id": sensorGrid,
        "by":"MQTT",
	}).Info("Sensor Data Published")
	Error = controllers.ManageSensorData(sensorGrid, result)

	return Error
}
