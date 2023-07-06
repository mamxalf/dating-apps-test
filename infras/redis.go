package infras

import (
	"dating-apps/configs"
	"fmt"

	"github.com/go-redis/redis"
)

const (
	ErrRedisNil = redis.Nil
)

type RedisConn struct {
	Client *redis.Client
}

// RedisNewClient create new instance of redis
func RedisNewClient(config *configs.Config) *RedisConn {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Cache.Redis.Primary.Host, config.Cache.Redis.Primary.Port),
		Password: config.Cache.Redis.Primary.Password,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	return &RedisConn{client}
}
