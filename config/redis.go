package config

import (
	"github.com/go-redis/redis"
	"go-project/global"
	"log"
)

func InitRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "",
	})

	_, err := RedisClient.Ping().Result()

	if err != nil {
		log.Fatal("连接redis数据库失败：" + err.Error())
	}

	global.RedisDb = RedisClient
}
