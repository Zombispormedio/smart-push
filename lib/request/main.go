package request

import (
	"encoding/json"

	"github.com/parnurzeal/gorequest"
)

func GET(url string, body interface{}) error {

	request := gorequest.New()
	_, resBody, _ := request.Get(url).End()

	return json.Unmarshal([]byte(resBody), body)
}

func GETWithHeader(url string, headers map[string]string,  body interface{}) error {

	request := gorequest.New()
	get:= request.Get(url)
	
	for k, v := range headers{
		get=get.Set(k, v)
	}
	
	_, resBody, _:=get.End()

	return json.Unmarshal([]byte(resBody), body)
}
