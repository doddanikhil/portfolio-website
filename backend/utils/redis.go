package utils

import (
	"log"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var RedisClient *redis.Client

func InitializeRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // Redis service name in Docker
		Password: "",           // Set the password here if Redis requires one
		DB:       0,            // Default DB
	})
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}
