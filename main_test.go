package main

import (
	"testing"

	"github.com/Zombispormedio/smart-push/lib/rabbit"
)

func TestMain(m *testing.T) {

	/*client:=redis.Client();

	  cursor, result, error:= client.Scan(0, "*user*", 2).Result()*/

	r, Error := rabbit.New("logs_direct", "direct", false)
m.Log(Error)
	body := struct {
		Hello string `json:"hello"`
	}{
		"hello",
	}

	err := r.PublishJSON("info", &body)

	m.Log(err)

	r.Close()

}
