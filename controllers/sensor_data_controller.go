package controllers

import (
	
	"reflect"
	"strings"

	"github.com/Zombispormedio/smart-push/lib/store"
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
	gridNodes := make([]string, numNodes)
	db, _:=store.OpenDB();
	
	
	for i := 0; i < numNodes; i++ {
		sensor := grid.Data[i]

		gridNodes[i] = sensor.NodeID

		NodeStoringError := store.PutWithDB(db, sensor.NodeID, sensor.Value, "Sensors")

		if NodeStoringError != nil {
			return NodeStoringError
		}
	}

	gridNodesStr := strings.Join(gridNodes, ",")

	GridStoringError := store.PutWithDB(db, sensorGridID, gridNodesStr, "Grids")
   
	db.Close();
	return GridStoringError
}
