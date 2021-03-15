package clients

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)
var (
	rdClient *redis.Client
)
func init() {
	// Set client options
	rdClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
		DB:       0,
	})

	Ctx    := context.TODO()
	if err := rdClient.Ping(Ctx).Err(); err != nil {
		log.Fatal(err)
	}
}

func GetRedisClient() *redis.Client {
	return rdClient
}
