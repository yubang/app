package webApi

import "net/http"

/*
PAAS API路由模块
@author: yubang
 */

// 返回路由
func getRoutes() map[string]func(w http.ResponseWriter, r *http.Request){
	routes := make(map[string]func(w http.ResponseWriter, r *http.Request))
	routes["/a"] = a
	return routes
}

func a(w http.ResponseWriter, r *http.Request){

}