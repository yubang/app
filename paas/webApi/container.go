package webApi

import (
	"../tools"
	"../config"
	"../redisClient"
	"../httpCode"
	"net/http"
	"github.com/go-redis/redis"
)

/*
容器相关API模块
@author: yubang
 */

// 获取一个容器操作任务
func getContainerTask(w http.ResponseWriter, r *http.Request){

	client := redisClient.GetRedisClient()
	defer client.Close()

	// 读取整个容器服务器计划分配资源方案
	ip, _ := getRequestIp(r)
	s, err := client.Get(config.REDIS_KEY_PLAN_CONTAIN_USE_STR + ip).Result()
	if err != nil{
		output(w, httpCode.ServerErrorCode, nil)
		return
	}
	sObj := tools.JsonToInterface([]byte(s))
	if sObj == nil{
		output(w, httpCode.ServerErrorCode, nil)
		return
	}

	output(w, httpCode.OkCode, sObj)
}

// 处理容器操作回调处理
func optionContainerCallback(w http.ResponseWriter, r *http.Request){

	ip, _ := getRequestIp(r)
	containerInfo := r.FormValue("containerInfo")

	if containerInfo == ""{
		output(w, httpCode.ParameterMissingCode, nil)
		return
	}

	containerList := tools.JsonToInterfaceList([]byte(containerInfo))
	if containerList == nil{
		output(w, httpCode.ParameterMissingCode, nil)
		return
	}

	// redis对象
	client := redisClient.GetRedisClient()
	defer client.Close()

	// 判断需要解除绑定的端口
	oldContainerInfoStr, err := client.HGet(config.REDIS_KEY_PROXY_APP_CONTAINER_HASH, ip).Result()
	if err == redis.Nil{
		oldContainerInfoStr = "[]"
	}else if err != nil{
		output(w, httpCode.ServerErrorCode, nil)
		return
	}
	oldContainerInfoList := tools.JsonToInterfaceList([]byte(oldContainerInfoStr))
	for _, v := range oldContainerInfoList{
		deleteSign := true
		containerMap := v.(map[string]interface{})

		// 遍历判断
		for _, v2 := range containerList{
			newContainerMap := v2.(map[string]interface{})
			if containerMap["name"].(string) == newContainerMap["name"].(string) && containerMap["port"].(float64) == newContainerMap["port"].(float64){
				deleteSign = false
				break
			}
		}

		if deleteSign{
			client.SRem(config.REDIS_KEY_PROXY_APP_CONTAINER_HASH + containerMap["app"].(string), ip + ":" + tools.Float64ToString(containerMap["port"].(float64)))
		}
	}

	// 绑定应用与后端的映射端口
	for _, v := range containerList{
		containerMap := v.(map[string]interface{})
		redisKey := config.REDIS_KEY_PROXY_APP_CONTAINER_IP_AND_PORT_SET + containerMap["app"].(string)
		client.SAdd(redisKey, ip + ":" + tools.Float64ToString(containerMap["port"].(float64)))
	}

	// 更新容器服务器数据
	client.HSet(config.REDIS_KEY_PROXY_APP_CONTAINER_HASH, ip, containerInfo)

	d := make(map[string]interface{})
	d["ip"] = ip
	output(w, httpCode.OkCode, d)
}

// 登记容器服务器
func loginContainerServer(w http.ResponseWriter, r *http.Request){
	ip, _ := getRequestIp(r)
	memory := r.FormValue("memory")
	disk := r.FormValue("disk")

	if memory == "" || disk == ""{
		output(w, httpCode.ParameterMissingCode, nil)
		return
	}

	client := redisClient.GetRedisClient()
	defer client.Close()

	// 设置容器服务器最后登录时间
	client.HSet(config.REDIS_KEY_CONTAINER_SERVER_LAST_TIME_HASH, ip, tools.GetNowTimeSecond())

	// 把容器扔到set
	length, err := client.SAdd(config.REDIS_KEY_CONTAINER_SERVER_IP_SET, ip).Result()

	if err == nil && length != 0{
		// client.ZAdd(config.REDIS_KEY_CONTAINER_SERVER_IP_ZSET, redis.Z{float64(length), ip})
		client.RPush(config.REDIS_KEY_CONTAINER_SERVER_IP_LIST, ip)
		// 同时初始化容器服务器分配计划
		client.Set(config.REDIS_KEY_PLAN_CONTAIN_USE_STR + ip, "{}", 0)
	}

	// 记录容器服务器信息
	redisKey := config.REDIS_KEY_CONTAINER_SERVER_INFO_HASH + ip
	client.HSet(redisKey, "memory", memory)
	client.HSet(redisKey, "disk", disk)

	d := make(map[string]interface{})
	d["ip"] = ip
	output(w, httpCode.OkCode, d)
}