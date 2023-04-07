package redisService

import (
	"context"
	"time"

	"github.com/aokoli/goutils"
	"github.com/go-redis/redis"
)

var ctx = context.Background()
var redisClient *redis.Client

func init() {
	// redisClient = redis.NewFailoverClient(&redis.FailoverOptions{
	// 	MasterName:    "main",
	// 	SentinelAddrs: strings.Split("10.0.19.102:26379", ","),
	// 	Password:      "6iDrKRF1OW5sKIvj",
	// })
	redisClient = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
}

func Get(key string) string {
	val, _ := redisClient.Get(key).Result()
	return val
}

func Del(key string) {
	redisClient.Del(key)
}

func Lock(key string, tm, unit time.Duration) bool {
	val, _ := goutils.CryptoRandomAscii(32)
	success, err := redisClient.SetNX(key, val, tm*unit).Result()
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
				vv, err := redisClient.Get(key).Result()
				if err == nil && vv == val {
					redisClient.Expire(key, expire)
				} else {
					break
				}
			}
		}()
	}
	return success
}

func UnLock(key string) {
	redisClient.Del(key).Result()
}