package webApi

import (
	"../tools"
	"../config"
	"../redisClient"
	"../httpCode"
	"net/http"
)

/*
容器相关API模块
@author: yubang
 */

// 获取一个容器操作任务
func getContainerTask(w http.ResponseWriter, r *http.Request){
	d := make(map[string]interface{})
	d["taskId"] = "123"
	d["imageName"] = "abcdef"
	output(w, httpCode.OkCode, d)
}

// 处理容器操作回调处理
func optionContainerCallback(w http.ResponseWriter, r *http.Request){

	taskId := r.FormValue("taskId")
	result := r.FormValue("result")
	host := r.FormValue("ip")
	port := r.FormValue("port")

	if taskId == "" || result == "" || host == "" || port == ""{
		output(w, httpCode.ParameterMissingCode, nil)
		return
	}

	d := make(map[string]interface{})
	output(w, httpCode.OkCode, d)
}

// 登记容器服务器
func loginContainerServer(w http.ResponseWriter, r *http.Request){
	ip, _ := getRequestIp(r)
	memory := r.FormValue("memory")
	disk := r.FormValue("disk")

	client := redisClient.GetRedisClient()
	defer client.Close()

	// 设置容器服务器最后登录时间
	client.HSet(config.REDIS_KEY_CONTAINER_SERVER_LAST_TIME_HASH, ip, tools.GetNowTimeSecond())

	// 把容器扔到set
	length, err := client.SAdd(config.REDIS_KEY_CONTAINER_SERVER_IP_SET, ip).Result()

	if err == nil && length != 0{
		// client.ZAdd(config.REDIS_KEY_CONTAINER_SERVER_IP_ZSET, redis.Z{float64(length), ip})
		client.RPush(config.REDIS_KEY_CONTAINER_SERVER_IP_LIST, ip)
	}

	// 记录容器服务器信息
	redisKey := config.REDIS_KEY_CONTAINER_SERVER_INFO_HASH + ip
	client.HSet(redisKey, "memory", memory)
	client.HSet(redisKey, "disk", disk)

	d := make(map[string]interface{})
	d["ip"] = ip
	output(w, httpCode.OkCode, d)
}