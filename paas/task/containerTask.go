package task

/*
容器分配策略模块
@author: yubang
 */

import (
	"../config"
	"../redisClient"
)

// 计算出一个应用的容器分配策略
func analyApp(appId string, client *redisClient.TypeRedisClient){
	// 先判断应用是否需要移除容器
	length, err := client.LLen(config.REDIS_KEY_CONTAINER_SERVER_IP_LIST).Result()
	if err != nil{
		length = 0
	}
	for index:=int64(0);index<length;index++{
		// 获取容器服务器ip
		containerIp, err2 := client.LIndex(config.REDIS_KEY_CONTAINER_SERVER_IP_LIST, index).Result()
		if err2 != nil{
			continue
		}
		// 获取容器服务器内存

		// 获取容器服务器当前分配信息
		
	}
}

// 计算一次容器分配策略
func AnalyContainer(){

	client := redisClient.GetRedisClient()
	defer client.Close()

	// 遍历所有应用
	length, err := client.LLen(config.REDIS_KEY_APP_LIST).Result()
	if err != nil{
		length = 0
	}
	for ;length > 0;length--{
		appId, err2 := client.LIndex(config.REDIS_KEY_APP_LIST, length-1).Result()
		if err2 == nill && appId != ""{
			analyApp(appId, client) // 处理应用分配
		}
	}
}