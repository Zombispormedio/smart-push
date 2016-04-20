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
