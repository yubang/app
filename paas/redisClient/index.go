package redisClient

import (
	"../config"
	"../tools"
	"github.com/go-redis/redis"
)

type TypeRedisClient  redis.Client

func GetRedisClient()*redis.Client{
	configObj := config.GetPaasConfig()
	client := redis.NewClient(&redis.Options{
		Addr:     configObj.RedisConfigData.Host + ":" + tools.IntToString(configObj.RedisConfigData.Port),
		Password: configObj.RedisConfigData.Password, // no password set
		DB:       configObj.RedisConfigData.Db, // use default DB
	})
	return client
}