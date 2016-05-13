package redis
import (
    "os"
    "gopkg.in/redis.v3"
)


type RedisWrapper struct{
    Client *redis.Client
}


func Client() *RedisWrapper{
 
    client:=redis.NewClient(&redis.Options{
        Addr:os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"),
         Password:os.Getenv("REDIS_AUTH"),
    });
    
    r:=&RedisWrapper{}
    r.Client=client
    return r
}


func (r *RedisWrapper)Get(key string) (string, error){
    return r.Client.Get(key).Result()
}


