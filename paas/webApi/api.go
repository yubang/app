package webApi

import (
	"../tools"
	"../config"
	"../task"
	"../redisClient"
	"../httpCode"
	"net/http"
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

// 处理应用信息
func handleApp(w http.ResponseWriter, r *http.Request, isAdd bool){
	appId := r.FormValue("appId")
	appHost := r.FormValue("appHost")
	sourceImage := r.FormValue("sourceImage")
	minContainerNumber := r.FormValue("minContainerNumber")
	maxContainerNumber := r.FormValue("maxContainerNumber")
	memory := r.FormValue("memory")
	gitUrl := r.FormValue("gitUrl")

	// 检查参数
	if appId == "" || appHost == "" || sourceImage == "" || minContainerNumber == "" || maxContainerNumber == "" || gitUrl == "" || memory == ""{
		output(w, httpCode.ParameterMissingCode, nil)
		return
	}

	client := redisClient.GetRedisClient()
	defer client.Close()

	// 设置应用信息
	redisKey := config.REDIS_KEY_APP_MESSAGE_HASH + appId
	client.HSet(redisKey, "gitUrl", gitUrl)
	client.HSet(redisKey, "appHost", appHost)
	client.HSet(redisKey, "minContainerNumber", minContainerNumber)
	client.HSet(redisKey, "maxContainerNumber", maxContainerNumber)
	client.HSet(redisKey, "buildingImage", "0")
	client.HSet(redisKey, "sourceImage", sourceImage) // 原始镜像，用于打包
	client.HSet(redisKey, "image", "") // 还没打包成镜像
	client.HSet(redisKey, "memory", memory)

	// 绑定域名与应用关系
	client.HSet(config.REDIS_KEY_PROXY_DOMAIN_APP_HASH, appHost, appId)

	if !isAdd{
		// 非新增应用，到此结束
		output(w, httpCode.OkCode, nil)
		return
	}

	// 压入队列
	err2 := client.RPush(config.REDIS_KEY_APP_LIST, appId).Err()
	if err2 != nil{
		output(w, httpCode.ServerErrorCode, nil)
		return
	}
	output(w, httpCode.OkCode, nil)
}

// 添加一个应用接口
func addApp(w http.ResponseWriter, r *http.Request){
	handleApp(w, r, true)
}

// 更新应用
func updateApp(w http.ResponseWriter, r *http.Request){
	handleApp(w, r, false)
}

// 删除应用
func removeApp(w http.ResponseWriter, r *http.Request){
	appId := r.FormValue("appId")
	if appId == ""{
		output(w, httpCode.ParameterMissingCode, nil)
		return
	}

	client := redisClient.GetRedisClient()
	defer client.Close()

	length, err := client.LLen(config.REDIS_KEY_APP_LIST).Result()
	if err != nil{
		output(w, httpCode.ServerErrorCode, nil)
		return
	}

	for ;length > 0;length--{
		d, err2 := client.LPop(config.REDIS_KEY_APP_LIST).Result()
		if err2 != nil{
			break
		}
		if d == appId{
			// 删除app信息
			client.Del(config.REDIS_KEY_APP_MESSAGE_HASH + appId)
			break
		}
		client.RPush(config.REDIS_KEY_APP_LIST, d)
	}

	output(w, httpCode.OkCode, nil)
}