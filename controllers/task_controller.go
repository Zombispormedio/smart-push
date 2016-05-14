package controllers

import (
	"errors"

	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/Zombispormedio/smart-push/config"
	"github.com/Zombispormedio/smart-push/lib/request"
	"github.com/Zombispormedio/smart-push/lib/response"
	"github.com/Zombispormedio/smart-push/lib/store"
	"github.com/boltdb/bolt"
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

type PushSensorData struct{
	NodeID string `json:"node_id"`
	Value  string `json:"value"`
	Date 	string `json:"date"`
}

type SensorGrid struct {
	ClientID string       `json:"client_id"`
	Data     []SensorData `json:"data"`
}

func PushOver() error {
	var Error error
	freq := config.PacketFrequency()

	db, OpenDBError := store.OpenDB()

	if OpenDBError != nil {
		return OpenDBError
	}

	Send := func(packet []SensorGrid) error {
		return SendSensorGridPacket(db, packet)
	}
	grids := []SensorGrid{}

	Error = store.Iterate(db, "Grids", func(c *bolt.Cursor) error {
		var err error

		for k, v := c.First(); k != nil; k, v = c.Next() {

			if len(grids) == freq {
				err = Send(grids)
				if err != nil {
					break
				} else {
					grids = nil
				}
			}

			grid := SensorGrid{}
			grid.ClientID = string(k)

			sensorIDs := strings.Split(string(v), ",")

			for _, nodeID := range sensorIDs {
				sensorData := SensorData{}

				sensorData.NodeID = nodeID
				var DBGettingError error
				sensorData.Value, DBGettingError = store.GetWithDB(db, "Sensors", nodeID)

				if DBGettingError != nil {
					err = DBGettingError
					break
				}

				grid.Data = append(grid.Data, sensorData)

			}

			if err != nil {
				break
			}

			grids = append(grids, grid)

		}

		if err == nil && len(grids) > 0 {
			err = Send(grids)
		}

		return err
	})

	db.Close()

	return Error
}

func SendSensorGridPacket(db *bolt.DB, packet []SensorGrid) error {
	var Error error

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

	return Error
}

func Clean() error {
	var Error error

	db, OpenDBError := store.OpenDB()

	if OpenDBError != nil {
		return OpenDBError
	}
	

	SensorsDeleteError := store.Iterate(db, "Sensors", func(c *bolt.Cursor) error {
		var err error
		for k, _ := c.First(); k != nil; k, _ = c.Next() {

			SensorError := store.DeleteWithDB(db, k, "Sensors")
			if SensorError != nil {
				err = SensorError

				log.WithFields(log.Fields{
					"message":SensorError,
				}).Error("DeleteSensorError")
				break
			}

		}
		return err
	})

	if SensorsDeleteError != nil {
		return SensorsDeleteError
	}

	GridsDeleteError := store.Iterate(db, "Grids", func(c *bolt.Cursor) error {
		var err error
		for k, _ := c.First(); k != nil; k, _ = c.Next() {

			GridError :=store.DeleteWithDB(db, k, "Grids")
			if GridError != nil {
				err = GridError
				log.WithFields(log.Fields{
					"message":GridError,
				}).Error("DeleteGridError")
				break
			}

		}
		return err
	})

	if GridsDeleteError != nil {
		Error = SensorsDeleteError
	}
	
	 db.Close();

	return Error
}
