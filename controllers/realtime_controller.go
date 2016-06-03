package controllers

import (
	"os"

	"time"

	"github.com/Zombispormedio/smart-push/lib/redis"
	"github.com/Zombispormedio/smart-push/lib/response"
	"github.com/Zombispormedio/smart-push/lib/utils"
)

func GetRealtimeData(sensor *response.RealTimeData) error {
	var Error error

	client := redis.Client()

	defer client.Close()

	sensorKeyGroup := os.Getenv("SENSOR_KEY") + ":" + sensor.ID

	timestampKeys, SensorDataError := client.KeysGroup(sensorKeyGroup)

	if SensorDataError != nil {
		Error = SensorDataError

	} 
	
	max:=utils.GetMaxTimestampKey(timestampKeys)
	
	nodeValue, GetError:= client.Get(max.Key)
	if GetError != nil{
		Error=GetError
	}else{	
	
		sensor.Value = nodeValue
		sensor.TimeStamp = time.Unix(max.Timestamp, 0).Format(time.RFC3339)
	}
	
	

	return Error
}
