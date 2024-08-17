package initializers

import (
	"github.com/redis/go-redis/v9"
	"os"
)

var RedisClient *redis.Client

func StartRedisClient() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
}
