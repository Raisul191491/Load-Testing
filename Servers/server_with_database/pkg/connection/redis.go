package connection

import (
	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func RedisConnect() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}

func GetRedis() *redis.Client {
	if redisClient == nil {
		RedisConnect()
		// if err := redisClient.FlushDB(context.Background()).Err(); err != nil {
		// 	fmt.Println("Failed to flush all databases: ", err)
		// }
	}
	return redisClient
}
