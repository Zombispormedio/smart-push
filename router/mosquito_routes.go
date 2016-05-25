package router

import (
    "fmt"
    "encoding/json"
)

func Mosquito(data []byte) error{
    var Error error
    
    result:=map[string]interface{}{}
    
     JSONError:=json.Unmarshal(data, &result)
    
    if JSONError !=nil{
        return JSONError
    }
    
    
    fmt.Println(result)
    
    
    return Error
}