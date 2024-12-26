package datasources

import (
	"os"

	"github.com/go-redis/redis"
)

var Cache *redis.Client

func init() {
	Cache = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("LMS_REDIS_ADDRESS"),
		Password: "",
		DB:       0,
	})
	err := Cache.Ping().Err()
	if err != nil {
		panic(err)
	}
}
