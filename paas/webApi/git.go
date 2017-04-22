package webApi

import "net/http"
import "../httpCode"
import "../redisClient"
import "../config"
import (
	"../task"
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

	taskId, err := client.LPop(config.REDIS_KEY_TASK_IMAGE_LIST).Result()
	if err == redis.Nil{
		output(w, httpCode.NotTaskCode, nil)
		return
	} else if err != nil{
		output(w, httpCode.ServerErrorCode, nil)
		return
	}

	taskObj := task.GetTaskObjFromTaskId(taskId)

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

	if taskId == ""{
		output(w, httpCode.ParameterMissingCode, nil)
		return
	}

	// 获取应用id
	obj := task.GetTaskObjFromTaskId(taskId)
	if obj == nil{
		output(w, httpCode.ServerErrorCode, nil)
		return
	}
	appId := obj["appId"].(string)


	client := redisClient.GetRedisClient()
	defer client.Close()

	redisKey := config.REDIS_KEY_APP_MESSAGE_HASH + appId
	if imageName != ""{
		// 标志镜像打包完成
		client.HSet(redisKey, "image" , imageName)
		client.HSet(redisKey, "buildingImage" , 1)
	}else{
		// pass
		client.HSet(redisKey, "buildingImage" , 2)
	}

	d := make(map[string]interface{})
	output(w, httpCode.OkCode, d)
}