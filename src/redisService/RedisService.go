package redisService

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-redis/redis"
)

var ctx = context.Background()
var redisClient *redis.Client

func init() {
	redisClient = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "main",
		SentinelAddrs: strings.Split("10.0.19.102:26379", ","),
		Password:      "6iDrKRF1OW5sKIvj",
	})
}

func Get(key string) {
	val, _ := redisClient.Get(key).Result()
	fmt.Println(val)
}

func Del(key string) {
	redisClient.Del(key)
}
