package utils

import (
	"log"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var RedisClient *redis.Client

func InitializeRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}
