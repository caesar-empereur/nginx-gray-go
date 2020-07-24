package rediss

import (
	"context"
	"github.com/astaxie/beego"
	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client
var LocalCtx = context.Background()

func GetRedisInstance() *redis.Client {

	if redisClient != nil {
		return redisClient
	}

	redisUrl := beego.AppConfig.String("redis_addr")

	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: "",
		DB:       0,
	})

	_, err := redisClient.Ping(LocalCtx).Result()
	if err != nil {
		panic("redis 连接出错")
	}

	return redisClient
}
