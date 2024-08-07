package persistance

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func Set(key string, value interface{}) error {
	return rdb.Set(ctx, key, value, 0).Err()
}
func Get(key string) (string, error) {
	return rdb.Get(ctx, key).Result()
}
