package redis

import (
	"os"
	"time"

	"gopkg.in/redis.v3"
)

type RedisWrapper struct {
	Client *redis.Client
}

func Client() *RedisWrapper {

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_AUTH"),
	})

	r := &RedisWrapper{}
	r.Client = client
	return r
}

func Status() error {
	r := Client()
	Error := r.Client.Ping().Err()
	r.Close()

	return Error
}

func (r *RedisWrapper) Get(key string) (string, error) {
	return r.Client.Get(key).Result()
}

func (r *RedisWrapper) Incr(key string) (int64, error) {
	return r.Client.Incr(key).Result()
}

func (r *RedisWrapper) KeysGroup(group string) ([]string, error) {
	return r.Client.Keys("*" + group + "*").Result()
}

func (r *RedisWrapper) Set(key string, value string) error {
	return r.Client.Set(key, value, -1).Err()
}

func (r *RedisWrapper) SetWithExpiration(key string, value string, expire time.Duration) error {
	return r.Client.Set(key, value, expire).Err()
}

func (r *RedisWrapper) Del(keys ...string) error {
	var Error error
	for _, k := range keys {
		Error = r.Client.Del(k).Err()
		if Error != nil {
			break
		}
	}

	return Error
}

func (r *RedisWrapper) SAdd(key string, members string) error {
	return r.Client.SAdd(key, members).Err()
}

func (r *RedisWrapper) SMembers(key string) ([]string, error) {
	return r.Client.SMembers(key).Result()
}

func (r *RedisWrapper) HMSetMap(key string, fields map[string]string) error {
	return r.Client.HMSetMap(key, fields).Err()
}

func (r *RedisWrapper) HGetAllMap(key string) (map[string]string, error) {
	return r.Client.HGetAllMap(key).Result()
}

func (r *RedisWrapper) Close() error {
	return r.Client.Close()
}

func (r *RedisWrapper) Expire(exp time.Duration, keys ...string) error {
    var Error error
    for _, k:=range keys{
        Error=r.Client.Expire(k, exp).Err()
        if Error != nil{
            break
        }
    }
    
    
	return Error
}
