package controllers

import (
	"os"

	"github.com/Zombispormedio/smart-push/lib/redis"
	"github.com/Zombispormedio/smart-push/lib/response"
)

func GetRealtimeData(sensor *response.RealTimeData) error {
	var Error error

	client := redis.Client()

	defer client.Close()

	sensorKey := os.Getenv("SENSOR_KEY") + ":" + sensor.ID

	dataMap, SensorDataError := client.HGetAllMap(sensorKey)

	if SensorDataError != nil {
		Error = SensorDataError

	} else {
		sensor.Value = dataMap["value"]
		sensor.TimeStamp = dataMap["date"]
	}

	return Error
}
