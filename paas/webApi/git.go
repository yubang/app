package webApi

import "net/http"
import "../httpCode"
import "../redisClient"
import "../config"
import (
	"../tools"
	"github.com/go-redis/redis"
)

/*
与git，docker镜像打包相关模块
@author: yubang
 */

// 拉取一条打包镜像任务
func getAGitPullTask(w http.ResponseWriter, r *http.Request){

	// 从redis获取任务
	client := redisClient.GetRedisClient()
	defer client.Close()

	redisKey, err := client.LPop(config.REDIS_KEY_TASK_IMAGE_LIST).Result()
	if err == redis.Nil{
		output(w, httpCode.NotTaskCode, nil)
		return
	} else if err != nil{
		output(w, httpCode.ServerErrorCode, nil)
		return
	}

	taskJson, err2 := client.HGet(config.REDIS_KEY_TASK_HASH + tools.GetSplitFirstArr(redisKey, "_"), redisKey).Result()

	if err2 == redis.Nil{
		output(w, httpCode.NotTaskCode, nil)
		return
	}else if err != nil{
		output(w, httpCode.ServerErrorCode, nil)
		return
	}

	taskObj := tools.JsonToInterface([]byte(taskJson))

	d := make(map[string]interface{})
	d["taskId"] = taskObj["taskId"]
	d["gitUrl"] = taskObj["gitUrl"]
	d["dockerImage"] = taskObj["image"]
	output(w, httpCode.OkCode, d)
}

// 打包镜像后通知
func buildImageCallback(w http.ResponseWriter, r *http.Request){

	taskId := r.FormValue("taskId")
	imageName := r.FormValue("imageName")
	result := r.FormValue("result")

	if taskId == "" || imageName == "" || result == ""{
		output(w, httpCode.ParameterMissingCode, nil)
		return
	}

	d := make(map[string]interface{})
	output(w, httpCode.OkCode, d)
}