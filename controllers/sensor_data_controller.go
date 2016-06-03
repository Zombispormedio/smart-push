package controllers

import (
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/Zombispormedio/smart-push/lib/redis"
	"github.com/Zombispormedio/smartdb/lib/struts"
)

type SensorData struct {
	NodeID string `json:"node_id"`
	Value  string `json:"value"`
}

func (sensorData *SensorData) FillByMap(Map map[string]interface{}, LiteralTag string) {
	struts.FillByMap(*sensorData, reflect.ValueOf(sensorData).Elem(), Map, LiteralTag)
}

type SensorGridData struct {
	Data []SensorData `json:"data"`
}

func (sensorGridData *SensorGridData) FillByMap(Map map[string]interface{}, LiteralTag string) {
	struts.FillByMap(*sensorGridData, reflect.ValueOf(sensorGridData).Elem(), Map, LiteralTag)
}

func ManageSensorData(sensorGridID string, data interface{}) error {
	grid := SensorGridData{}

	grid.FillByMap(data.(map[string]interface{}), "json")

	numNodes := len(grid.Data)

	client := redis.Client()

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	nodeIDs := make([]string, numNodes)

	for i := 0; i < numNodes; i++ {
		sensor := grid.Data[i]

		nodeIDs[i] = sensor.NodeID

		nodeKey := os.Getenv("SENSOR_KEY") + ":" + sensor.NodeID + ":" + timestamp

		SetError := client.SetWithExpiration(nodeKey, sensor.Value, time.Hour*4)

		if SetError != nil {
			return SetError
		}
	}

	gridGroup := os.Getenv("GRID_KEY") + ":" + sensorGridID
	gridKey := gridGroup + ":" + timestamp

	timestampKeys, KeyGroupError := client.KeysGroup(gridGroup)

	if KeyGroupError != nil {
		return KeyGroupError
	}

	var insert bool

	if len(timestampKeys) > 0 {
		oldGridKey := timestampKeys[0]
		gridOldValue, GETError := client.Get(oldGridKey)

		if GETError != nil {
			return GETError
		}

		if len(strings.Split(gridOldValue, ",")) != len(nodeIDs) {
			DelError := client.Del(oldGridKey)
			if DelError != nil {
				return DelError
			}
			insert = true

		}

	} else {
		insert = true
	}

	if insert {
		SetGridError := client.SetWithExpiration(gridKey, strings.Join(nodeIDs, ","), time.Hour*5)
		if SetGridError != nil {
			return SetGridError
		}
	}

	return client.Close()
}
