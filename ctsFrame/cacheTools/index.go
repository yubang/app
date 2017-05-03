package cacheTools

/*
简单缓存模块
@author: yubang
创建于2017年4月23日
 */

import (
	"../jsonTools"
	"../typeConversionTools"
	"github.com/go-redis/redis"
	"time"
)


type RedisClientObject struct {
	Host string
	Port int
	Db int
	Password string
	redisClient *redis.Client
}

func (config *RedisClientObject)GetRedisClient()*redis.Client{
	if config.redisClient == nil || config.redisClient.Ping().Err() != nil{
		config.getNewRedisClient()
	}
	return config.redisClient
}

func (config *RedisClientObject)getNewRedisClient()*redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + typeConversionTools.IntToString(config.Port),
		Password: config.Password, // no password set
		DB:       config.Db,  // use default DB
	})
	if config.redisClient != nil{
		config.redisClient.Close()
	}
	config.redisClient = client
	return client
}


func GetRedisClientObject(info map[string]interface{})RedisClientObject{
	host, port, db, password := "127.0.0.1", 6379, 0, ""
	if info["host"] != nil{
		host = info["host"].(string)
	}
	if info["port"] != nil{
		port = info["port"].(int)
	}
	if info["db"] != nil{
		db = info["db"].(int)
	}
	if info["password"] != nil{
		password = info["password"].(string)
	}
	return RedisClientObject{host, port, db, password, nil,}
}


type CacheClient struct {
	CachePrefix string
	RedisClient RedisClientObject // redis client 对象
}

func (clientObj *CacheClient)GetRedisClient()*redis.Client{
	return clientObj.RedisClient.GetRedisClient()
}

func (clientObj *CacheClient)Set(key string, value interface{}){
	data := map[string]interface{}{
		"data": value,
		"expirTime": int64(0),
	}
	c := clientObj.GetRedisClient()
	c.Set(clientObj.CachePrefix + key, jsonTools.InterfaceToJson(data), 0)
}

func (clientObj *CacheClient)Get(key string)interface{}{
	c := clientObj.GetRedisClient()
	s, err := c.Get(clientObj.CachePrefix + key).Result()
	if err != nil{
		return nil
	}
	obj := jsonTools.JsonToInterface([]byte(s))
	if obj == nil{
		return nil
	}
	return obj["data"]
}

func (clientObj *CacheClient)Remove(key string){
	c := clientObj.GetRedisClient()
	c.Del(clientObj.CachePrefix + key)
}

func (clientObj *CacheClient)Exist(key string) bool{
	c := clientObj.GetRedisClient()
	return  c.Get(clientObj.CachePrefix + key).Err() != nil
}

func (clientObj *CacheClient)Expir(key string, timeout int){
	c := clientObj.GetRedisClient()
	c.Expire(clientObj.CachePrefix + key,  time.Duration(timeout) * time.Second)
}
