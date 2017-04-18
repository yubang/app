package webApi

import "net/http"
import "../config"

/*
PAAS API路由模块
@author: yubang
 */

// 返回路由
func getRoutes() map[string]func(w http.ResponseWriter, r *http.Request){
	routes := make(map[string]func(w http.ResponseWriter, r *http.Request))
	routes[config.BuildImageTaskAPI] = getAGitPullTask
	routes[config.BuildImageCallbackAPI] = buildImageCallback
	return routes
}