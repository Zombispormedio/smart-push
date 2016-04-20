package controllers

import (

	
	"os"

	"github.com/Zombispormedio/smart-push/lib/request"
	"github.com/Zombispormedio/smart-push/lib/response"
    "github.com/Zombispormedio/smart-push/lib/store"
)





func RefreshCredentials() error {
	var Error error
	hostname := os.Getenv("SENSOR_STORE_HOSTNAME")
	url := hostname + "push/config/credentials"

	makeReqWithAuthorization:=func(auth string){
        headers:=map[string]string{
            "Authorization":auth,
        }
        
        msg:=response.MessageT{}
        
        request.GETWithHeader(url, headers, msg )
        
       
    }
    
    Error=store.Get("identifier", "Config", makeReqWithAuthorization)
    
	return Error
}
