package cache

import (
	"api/config"
	"log"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func ConnectRedis() {

	redis_addr := config.Get("REDIS_HOST") + ":" + config.Get("REDIS_PORT")
	redis_password := config.Get("REDIS_PASSWORD")

	redis_options := &redis.Options{
		Addr: redis_addr,
	}

	if len(redis_password) != 0 {
		redis_options.Password = redis_password
	}

	client := redis.NewClient(redis_options)

	_, err := client.Ping().Result()

	if err != nil {
		log.Fatal(err)
	}

	RedisClient = client
}
