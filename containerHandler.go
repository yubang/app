package main

import (
	"net/url"
	"./paas/config"
	"./paas/tools"
	"./paas/osutil"
	"./paas/docker"
)

/*
容器处理模块
@author: yubang
 */

// 登记容器服务器
func loginContainer(){
	memory, disk := osutil.GetTotalMemory(), osutil.GetTotalDisk()
	configObj := config.GetPaasConfig()

	// 调用接口请求数据库
	apiUrl := "http://" + configObj.ApiServerConfigData.Ip + ":" + tools.Float64ToString(configObj.ApiServerConfigData.Port ) + config.LoginContainerAPI
	tools.Post(apiUrl, url.Values{"token":{configObj.Token}, "memory": {tools.IntToString(memory)}, "disk":{tools.IntToString(disk)}})
}

// 同步计划分配资源与实际分配资源
func synchronizationAllocation(){
	configObj := config.GetPaasConfig()
	apiUrl := "http://" + configObj.ApiServerConfigData.Ip + ":" + tools.Float64ToString(configObj.ApiServerConfigData.Port) + config.GetOptionContainerTaskAPI
	obj := tools.Post(apiUrl, url.Values{"token": {configObj.Token}})
	if obj == nil || obj["code"].(float64) != 0 {
		tools.Error("拉取操作同步任务失败！")
		return
	}

	// 拉取服务器上所有容器
	containerList := docker.GetAllContainerList()

	// 删除不存在imageName的容器
	for _, v1 := range containerList{
		sign := true
		for _, v2 := range obj["data"].([]map[string]interface{}){
			if v1["imageName"].(string) == v2["image"].(string){
				sign = true
				break
			}
		}
		if sign{
			// 移除容器
			docker.RemoveAContainer(v1["containerId"].(string))
		}
	}

	addPopMap := make(map[string]int)

	for _, v := range obj["data"].(map[string]interface{}){
		// 计算容器需要增加或者减少的次数
		n := docker.GetATypeOfImageContainerNumber(v.(map[string]interface{})["image"].(string))
		addPopMap[v.(map[string]interface{})["image"].(string)] = tools.Float64ToInt(v.(map[string]interface{})["num"].(float64)) - n
	}

	// 处理加减容器
	for k, v := range addPopMap{
		if v > 0{
			for index:= v;index >0; index--{
				// 添加一个容器
				port := docker.GetAbleUserPort()
				docker.StartAContainer(k, port, 80)
			}
		}else if v < 0{
			for index:= v;index <0; index++{
				// 移除一个容器
				containerId := docker.GetAContainer(k)
				docker.RemoveAContainer(containerId)
			}
		}
	}

}

// 记录容器情况
func callbackAllocation(){

	containerList := docker.GetAllContainerList()
	

}

func main(){
	loginContainer() // 登记服务器
	synchronizationAllocation() // 同步资源
	callbackAllocation() // 记录容器资源

	//// 操作容器逻辑
	//
	//// 回调结果
	//apiUrl = "http://" + configObj.ApiServerConfigData.Ip + ":" + tools.Float64ToString(configObj.ApiServerConfigData.Port) + config.OptionContainerCallbackAPI
	//taskId := (obj["data"].(map[string]interface{}))["taskId"].(string)
	//obj = tools.Post(apiUrl, url.Values{"token": {configObj.Token}, "taskId": {taskId}, "ip":{"127.0.0.1"}, "result": {"OK"}, "port": {"9000"}})
	//
	//fmt.Print(obj)
}
