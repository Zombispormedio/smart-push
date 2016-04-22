package controllers

import (
    "reflect"
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

func ManageSensorData(data interface{}) error {

    

	return nil
}
