package main

/*
打包应用成为镜像的程序
@author: yubang
 */

import (
	"./paas/config"
	"./paas/tools"
	"fmt"
	"net/url"
)

func main(){
	configObj := config.GetPaasConfig()
	apiUrl := "http://" + configObj.ApiServerConfigData.Ip + ":" + tools.Float64ToString(configObj.ApiServerConfigData.Port) + config.BuildImageTaskAPI
	obj := tools.Post(apiUrl, url.Values{"token": {configObj.Token}})
	if obj == nil{
		tools.Error("拉取构建镜像任务失败！")
		return
	}
	// 构建镜像逻辑

	// 回调结果
	apiUrl = "http://" + configObj.ApiServerConfigData.Ip + ":" + tools.Float64ToString(configObj.ApiServerConfigData.Port) + config.BuildImageCallbackAPI
	taskId := (obj["data"].(map[string]interface{}))["taskId"].(string)
	obj = tools.Post(apiUrl, url.Values{"token": {configObj.Token}, "taskId": {taskId}, "imageName":{"abc"}, "result": {"OK"}})

	fmt.Print(obj)
}