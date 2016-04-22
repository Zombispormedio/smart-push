package controllers

type SensorData struct{
    NodeID     string `json:"node_id"`
	Value string `json:"value"`
}

type SensorGridData struct{
    
}


func ManageSensorData(data interface{}) error{
    return  nil
}