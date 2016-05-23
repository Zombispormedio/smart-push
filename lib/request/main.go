package request

import (
	"encoding/json"
	"os"

	"github.com/Zombispormedio/smart-push/lib/response"
	"github.com/Zombispormedio/smart-push/lib/store"
	"github.com/parnurzeal/gorequest"
)

func GET(url string, body interface{}) error {

	request := gorequest.New()
	_, resBody, _ := request.Get(url).End()

	return json.Unmarshal([]byte(resBody), body)
}

func GETWithHeader(url string, headers map[string]string, body interface{}) error {

	request := gorequest.New()
	get := request.Get(url)

	for k, v := range headers {
		get = get.Set(k, v)
	}

	_, resBody, _ := get.End()

	return json.Unmarshal([]byte(resBody), body)
}

func GetWithAuthorization(url string, body interface{}) error {

	var identifier string
	GetKeyError := store.Get("identifier", "Config", func(value string) {
		identifier = value
	})

	if GetKeyError != nil {
		return GetKeyError
	}

	headers := map[string]string{
		"Authorization": identifier,
	}
	
	return GETWithHeader(url, headers, body)

}

func PostWithAuthorization(url string, reqBody interface{}, result interface{}) error {

	var identifier string
	GetKeyError := store.Get("identifier", "Config", func(value string) {
		identifier = value
	})

	if GetKeyError != nil {
		return GetKeyError
	}

	request := gorequest.New()
	_, resBody, _ := request.Post(url).
		Set("Authorization", identifier).
		Send(reqBody).
		End()
		
		

	return json.Unmarshal([]byte(resBody), result)
}


func PostWithHeaders(url string, reqBody interface{}, headers map[string]string, result interface{}) error {

	request := gorequest.New()
	post := request.Post(url)

	for k, v := range headers {
		post = post.Set(k, v)
	}
	post=post.Send(reqBody)

	_, resBody, _ := post.End()

	return json.Unmarshal([]byte(resBody), result)
}

func CheckSensorGrid(req response.ReqSensorT) (bool, error) {
	accepted := false
	hostname := os.Getenv("SENSOR_STORE_HOSTNAME")
	url := hostname + "push/sensor_grid/check"
	resBody := &response.MixedMessageT{}
	RequestError := PostWithAuthorization(url, req, resBody)

	if RequestError != nil {
		return accepted, RequestError
	}

	if resBody.Status == 0 {
		accepted = true
	}

	return accepted, RequestError

}


func DBStatus() (bool, error){
	accepted := false
	hostname := os.Getenv("SENSOR_STORE_HOSTNAME")
	url := hostname + "status"
	resBody := &response.DataT{}
	RequestError := GetWithAuthorization(url,resBody)

	
	if RequestError != nil {
		return accepted, RequestError
	}

	if resBody.Status == 0 {
		accepted = true
	}

	return accepted, RequestError
}
