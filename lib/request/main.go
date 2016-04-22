package request

import (
	"encoding/json"

	"github.com/Zombispormedio/smartdb/lib/store"
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
