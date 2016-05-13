package main
import (
    "testing"
    "github.com/Zombispormedio/smart-push/lib/redis"
)



func TestMain(m *testing.T){
    
     
    client:=redis.Client();
    
   cursor, result, error:= client.Scan(0, "*user*", 2).Result()
    
    
    
  m.Log(cursor, result, error)
    
}