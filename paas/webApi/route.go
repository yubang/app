package webApi

import (
	"../httpCode"
	"strings"
	"../config"
	"net/http"
	"../tools"
)

/*
PAAS API路由模块
@author: yubang
 */


func getRequestIp(r *http.Request) (string, int){
	arrs := strings.Split(r.RemoteAddr, ":")
	return arrs[0], tools.StringToInt(arrs[1])
}


// 返回客户端真实ip和端口
func getIp(w http.ResponseWriter, r *http.Request){
	d := make(map[string]interface{})
	d["ip"], d["port"] = getRequestIp(r)
	output(w, httpCode.OkCode, d)
}

// 返回路由
func getRoutes() map[string]func(w http.ResponseWriter, r *http.Request){
	routes := make(map[string]func(w http.ResponseWriter, r *http.Request))

	// 镜像相关操作API
	routes[config.BuildImageTaskAPI] = getAGitPullTask
	routes[config.BuildImageCallbackAPI] = buildImageCallback

	// ip相关API
	routes[config.GetIpApi] = getIp

	// 容器相关操作API
	routes[config.GetOptionContainerTaskAPI] = getContainerTask
	routes[config.OptionContainerCallbackAPI] = optionContainerCallback
	routes[config.LoginContainerAPI] = loginContainerServer

	// UI控制台相关API
	routes[config.WEBAPI_BuildImage] = buildImage

	// APP操作相关API
	routes[config.WEBAPI_AddApp] = addApp
	routes[config.WEBAPI_RemoveApp] = removeApp
	routes[config.WEBAPI_UpdateApp] = updateApp

	return routes
}