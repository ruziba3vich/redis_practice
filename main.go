package main

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
	redisDb := redis.NewClient(
		&redis.Options{
			Addr: "localhost:6379",
		},
	)

	err := redisDb.Set(ctx, "name", "Dostonbek", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

}
