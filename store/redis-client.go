package store

import (
	"fmt"

	"github.com/go-redis/redis"
)

// RedisClient returns a client.
func RedisClient(cfg *Conf) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisHost + ":" + cfg.RedisPort,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(pong)
	return client
}
