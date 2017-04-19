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


// 返回客户端真实ip和端口
func getIp(w http.ResponseWriter, r *http.Request){
	d := make(map[string]interface{})
	arrs := strings.Split(r.RemoteAddr, ":")
	d["ip"] = arrs[0]
	d["port"] = tools.StringToInt(arrs[1])
	output(w, httpCode.OkCode, d)
}

// 返回路由
func getRoutes() map[string]func(w http.ResponseWriter, r *http.Request){
	routes := make(map[string]func(w http.ResponseWriter, r *http.Request))
	routes[config.BuildImageTaskAPI] = getAGitPullTask
	routes[config.BuildImageCallbackAPI] = buildImageCallback
	routes[config.GetIpApi] = getIp
	routes[config.GetOptionContainerTaskAPI] = getContainerTask
	routes[config.OptionContainerCallbackAPI] = optionContainerCallback

	routes[config.WEBAPI_BuildImage] = buildImage
	return routes
}