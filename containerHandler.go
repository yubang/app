package main

import (
	"net/url"
	"fmt"
	"./paas/config"
	"./paas/tools"
)

/*
容器处理模块
@author: yubang
 */

func main(){
	configObj := config.GetPaasConfig()
	apiUrl := "http://" + configObj.ApiServerConfigData.Ip + ":" + tools.Float64ToString(configObj.ApiServerConfigData.Port) + config.GetOptionContainerTaskAPI
	obj := tools.Post(apiUrl, url.Values{"token": {configObj.Token}})
	if obj == nil{
		tools.Error("拉取操作容器任务失败！")
		return
	}
	// 操作容器逻辑

	// 回调结果
	apiUrl = "http://" + configObj.ApiServerConfigData.Ip + ":" + tools.Float64ToString(configObj.ApiServerConfigData.Port) + config.OptionContainerCallbackAPI
	taskId := (obj["data"].(map[string]interface{}))["taskId"].(string)
	obj = tools.Post(apiUrl, url.Values{"token": {configObj.Token}, "taskId": {taskId}, "ip":{"127.0.0.1"}, "result": {"OK"}, "port": {"9000"}})

	fmt.Print(obj)
}
