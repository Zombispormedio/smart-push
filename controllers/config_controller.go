package controllers

import (
	"errors"

	"fmt"
	"os"

	"github.com/Zombispormedio/smart-push/request"
	"github.com/Zombispormedio/smart-push/response"
)

func RefreshCredentials() error {
	var Error error
	hostname := os.Getenv("SENSOR_STORE_HOSTNAME")
	url := hostname + "push/config/credentials"

	body := response.MessageT{}

	request.GET(url, &body)

	fmt.Println(body.Message)

	Error = errors.New("hello")

	return Error
}
