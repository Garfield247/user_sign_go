package db

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func GetRedisInstance() *redis.Client {
	if redisClient != nil {
		return redisClient
	}
	redisConfig := ""
	fmt.Printf("redisConfig: %v\n", redisConfig)
	// 使用 URI 连接 Redis
	opt, err := redis.ParseURL(redisConfig)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return nil
	}

	client := redis.NewClient(opt)
	redisClient = client
	return client
}
