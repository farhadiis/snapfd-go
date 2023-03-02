package redis

import (
	"context"
	"farhadiis/order-process/utils"
	"github.com/redis/go-redis/v9"
	"log"
)

func Connect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: utils.GetEnv("REDIS_URI"),
		DB:   utils.GetEnvInt("REDIS_DB"),
	})
	err := client.Ping(context.Background()).Err()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Redis initialize successful")
	return client
}

func Disconnect(client *redis.Client) {
	if err := client.Close(); err != nil {
		panic(err)
	}
}
