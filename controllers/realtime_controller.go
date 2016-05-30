package controllers

import (
	"os"
"time"
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
		var date=dataMap["date"]
		unixDate, _ := strconv.ParseInt(date, 10, 64)
		sensor.Value = dataMap["value"]
		sensor.TimeStamp = time.Unix(unixDate, 0)
	}

	return Error
}
