package controllers

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/Zombispormedio/smart-push/config"
	"github.com/Zombispormedio/smart-push/lib/rabbit"
	"github.com/Zombispormedio/smart-push/lib/redis"

	"github.com/Zombispormedio/smart-push/lib/request"
	"github.com/Zombispormedio/smart-push/lib/response"
	"github.com/Zombispormedio/smart-push/lib/store"
	"github.com/Zombispormedio/smart-push/lib/utils"
)

func RefreshCredentials() error {
	var Error error
	hostname := os.Getenv("SENSOR_STORE_HOSTNAME")
	url := hostname + "push/credentials"

	msg := response.DataT{}

	RequestError := request.GetWithAuthorization(url, &msg)

	if RequestError != nil {
		return RequestError
	}

	if msg.Data == nil {
		return errors.New("No Authorized")
	}

	data := msg.Data.(map[string]interface{})

	StoringError := store.Put("identifier", data["key"].(string), "Config")

	if StoringError != nil {
		Error = StoringError
	}

	return Error
}

type PushSensorData struct {
	NodeID string `json:"node_id"`
	Value  string `json:"value"`
	Date   string `json:"date"`
}

type PushSensorGrid struct {
	ClientID string           `json:"client_id"`
	Data     []PushSensorData `json:"data"`
}

func GetSensorData(client *redis.RedisWrapper, sensorKeys []string, grid *PushSensorGrid) error {
	var Error error
	for _, nodeID := range sensorKeys {
		sensorData := PushSensorData{}

		sensorData.NodeID = nodeID

		sensorGroup := os.Getenv("SENSOR_KEY") + ":" + nodeID

		keys, GroupError := client.KeysGroup(sensorGroup)

		if GroupError != nil {
			Error = GroupError
			log.WithFields(log.Fields{
				"group": sensorGroup,
				"error": GroupError.Error(),
			}).Error("SensorGroupError")

			break
		}

		max := utils.GetMaxTimestampKey(keys)
		min := utils.GetMinTimestampKey(keys)

		value, GetError := client.Average(keys...)

		if GetError != nil {
			Error = GetError
			log.WithFields(log.Fields{
				"group": max.Key,
				"error": GetError.Error(),
			}).Error("SensorGetError")
			break
		}
		

		sensorData.Value = strconv.FormatFloat(value,'f', 6,64)
		sensorData.Date = strconv.FormatInt(max.Timestamp, 10)

		grid.Data = append(grid.Data, sensorData)
		
		SetExpirationError := client.Expire(time.Hour*4, max.Key)
		
		if SetExpirationError !=nil{
			return SetExpirationError
		}
		
		
		
		if len(keys) > 1 {
			
			Error = client.Expire(time.Second, min.Key)
			
			if Error != nil {
				log.WithFields(log.Fields{
					"error": Error.Error(),
				}).Error("SensorDelError")
				break
			}

		}

	}

	return Error
}

func PushOver() error {
	var Error error
	freq := config.PacketFrequency()

	Send := func(packet []PushSensorGrid) error {
		return SendSensorGridPacket(packet)
	}
	grids := []PushSensorGrid{}

	client := redis.Client()

	defer client.Close()

	gridKeys, GroupError := client.KeysGroup(os.Getenv("GRID_KEY"))

	if GroupError != nil {
		log.WithFields(log.Fields{
			"error": GroupError.Error(),
		}).Error("GridGroupError")
		return GroupError
	}

	for _, gridkey := range gridKeys {

		if len(grids) == freq {
			SendError := Send(grids)
			if SendError != nil {
				Error = SendError
				log.WithFields(log.Fields{
					"error": Error.Error(),
				}).Error("SendError")
				break
			} else {
				grids = nil
			}
		}

		sensorStr, SensorKeysError := client.Get(gridkey)

		if SensorKeysError != nil {
			Error = SensorKeysError
			log.WithFields(log.Fields{
				"error": Error.Error(),
			}).Error("SensorKeysError")
			break
		}

		sensorKeys := strings.Split(sensorStr, ",")

		elems := strings.Split(gridkey, ":")
		clientID := elems[1]

		grid := PushSensorGrid{}
		grid.ClientID = clientID

		Error = GetSensorData(client, sensorKeys, &grid)

		if Error != nil {
			break
		}

		grids = append(grids, grid)

	}

	if Error == nil && len(grids) > 0 {
		Error = Send(grids)
	}
	timeRegistryExpire, _ := strconv.Atoi(os.Getenv("PUSH_TIME"))
	dateRegistryKey := os.Getenv("TIME_KEY") + ":" + strconv.FormatInt(time.Now().Unix(), 10)
	log.Info(time.Minute*time.Duration(timeRegistryExpire));
	client.SetWithExpiration(dateRegistryKey, "0", time.Minute*time.Duration(timeRegistryExpire))

	return Error
}

func SendSensorGridPacket(packet []PushSensorGrid) error {
	var Error error

	db, OpenDBError := store.OpenDB()

	if OpenDBError != nil {
		return OpenDBError
	}

	identifier, GetKeyError := store.GetWithDB(db, "Config", "identifier")

	if GetKeyError != nil {
		return GetKeyError
	}

	hostname := os.Getenv("SENSOR_STORE_HOSTNAME")
	url := hostname + "push/sensor_grid"
	headers := map[string]string{"Authorization": identifier}

	resBody := &response.MixedMessageT{}

	RequestError := request.PostWithHeaders(url, packet, headers, resBody)

	if RequestError != nil {
		return RequestError
	}

	if resBody.Status != 0 {
		Error = errors.New(resBody.Error)
	}

	db.Close()

	return Error
}

func Clean() error {
	var Error error

	client := redis.Client()

	defer client.Close()

	sensorKeys, SensorKeysError := client.KeysGroup(os.Getenv("Sensor_KEY"))

	if SensorKeysError != nil {
		return SensorKeysError
	}

	SensorCleanError := client.Del(sensorKeys...)

	if SensorCleanError != nil {
		return SensorCleanError

	}

	gridKeys, GridKeysError := client.KeysGroup(os.Getenv("GRID_KEY"))

	if GridKeysError != nil {
		return GridKeysError
	}

	gridCleanError := client.Del(gridKeys...)

	if gridCleanError != nil {
		Error = gridCleanError

	}

	return Error
}

func PushRabbit() error {
	var Error error

	awake, DBStatusError := request.DBStatus()

	if DBStatusError != nil {
		return DBStatusError
	}

	if !awake {
		return errors.New("No awake DB")
	}

	client := redis.Client()
	rClient, RError := rabbit.New(os.Getenv("EX_RABBIT"), "topic", true)

	if RError != nil {
		return RError
	}

	defer client.Close()
	defer rClient.Close()

	var rKey string

	GetKeyError := store.Get("identifier", "Config", func(value string) {
		rKey = value + ".push"
	})

	if GetKeyError != nil {
		return GetKeyError
	}

	gridKeys, Error := client.KeysGroup(os.Getenv("GRID_KEY"))

	if Error != nil {
		return Error
	}

	for _, gridkey := range gridKeys {

		sensorStr, SensorKeysError := client.Get(gridkey)

		if SensorKeysError != nil {
			Error = SensorKeysError
			break
		}

		sensorKeys := strings.Split(sensorStr, ",")

		elems := strings.Split(gridkey, ":")
		clientID := elems[1]

		grid := PushSensorGrid{}
		grid.ClientID = clientID

		SensorDataError := GetSensorData(client, sensorKeys, &grid)

		if SensorDataError != nil {
			Error = SensorDataError
			break
		}

		Error = rClient.PublishJSON(rKey, &grid)

		if Error != nil {
			break
		}
	}

	if Error == nil {
		timeRegistryExpire, _ := strconv.ParseInt(os.Getenv("PUSH_TIME"), 10, 64)
		dateRegistryKey := os.Getenv("TIME_KEY") + ":" + strconv.FormatInt(time.Now().Unix(), 10)

		client.SetWithExpiration(dateRegistryKey, "0", time.Minute*time.Duration(timeRegistryExpire))
	}

	return Error
}
