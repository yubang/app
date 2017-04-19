package webApi

import "net/http"
import "../httpCode"
import "../redisClient"
import "../task"
import "../config"
import (
	"../tools"
)

/*
提供给UI的paas平台接口
@author: yubang
 */


// 构建一个镜像接口
func buildImage(w http.ResponseWriter, r *http.Request){
	appId := r.FormValue("appId")
	gitUrl := r.FormValue("gitUrl")
	image := r.FormValue("image")
	if appId == "" || gitUrl == "" || image == ""{
		output(w, httpCode.ParameterMissingCode, nil)
		return
	}

	// 构建一条任务信息
	taskId := task.BuildTaskId()
	client := redisClient.GetRedisClient()
	defer client.Close()

	taskObj := make(map[string]interface{})
	taskObj["appId"] = appId
	taskObj["taskId"] = taskId
	taskObj["gitUrl"] = gitUrl
	taskObj["image"] = image
	taskJson := tools.InterfaceToJson(taskObj)

	// 标识应用正在构建镜像中
	client.HSet(config.REDIS_KEY_APP_MESSAGE_HASH + appId, "buildingImage", "1")

	// 写入任务池
	taskHashKey := config.REDIS_KEY_TASK_HASH + tools.GetSplitFirstArr(taskId, "_")
	client.HSet(taskHashKey, taskId, taskJson)
	client.Expire(taskHashKey, 3600 * 24 * 1000 * 1000)

	// 写入构建镜像队列
	client.RPush(config.REDIS_KEY_TASK_IMAGE_LIST, taskId)

	d := make(map[string]interface{})
	d["taskId"] = taskId
	output(w, httpCode.OkCode, d)
}
