package controllers

import (
	"errors"
	"os"

	"github.com/Zombispormedio/smart-push/lib/request"
	"github.com/Zombispormedio/smart-push/lib/response"
	"github.com/Zombispormedio/smart-push/lib/store"
)

func RefreshCredentials() error {
	var Error error
	hostname := os.Getenv("SENSOR_STORE_HOSTNAME")
	url := hostname + "push/config/credentials"
	

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
