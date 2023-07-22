package redisService

import (
	"context"
	"strings"
	"time"

	"github.com/aokoli/goutils"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var redisClient *redis.Client

func init() {
	redisClient = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "mymaster",
		SentinelAddrs: strings.Split("10.0.20.67:26379", ","),
		Password:      "donghaotian",
		DB:            0,
	})
	// redisClient = redis.NewClient(&redis.Options{
	// 	Addr:     "10.0.20.67:6379",
	// 	DB:       0,
	// 	Password: "donghaotian",
	// })
}

func Get(key string) string {
	val, _ := redisClient.Get(context.Background(), key).Result()
	return val
}

func Put(key, val string, tm, unit time.Duration) {
	redisClient.Set(context.Background(), key, val, tm*unit)
}

func Del(key string) {
	redisClient.Del(context.Background(), key)
}

func Lock(key string, tm, unit time.Duration) bool {
	val, _ := goutils.CryptoRandomAscii(32)
	success, err := redisClient.SetNX(context.Background(), key, val, tm*unit).Result()
	if err != nil {
		return false
	}
	if success {
		// 锁续期，直到当前节点释放锁或节点死亡
		go func() {
			expire := tm * unit
			if expire < 10*time.Second {
				return
			}
			minSleepTime := expire / 3 * 2
			for {
				time.Sleep(minSleepTime)
				vv, err := redisClient.Get(context.Background(), key).Result()
				if err == nil && vv == val {
					redisClient.Expire(context.Background(), key, expire)
				} else {
					break
				}
			}
		}()
	}
	return success
}

func UnLock(key string) {
	redisClient.Del(context.Background(), key).Result()
}
