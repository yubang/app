package webApi


/*
API http服务模块
@author: yubang
 */

import (
	"../../paas/config"
	"../../paas/tools"
	"net/http"
)

func handleNotFound(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(404)
	w.Write([]byte("not found!"))
}

func handler(w http.ResponseWriter, r *http.Request){
	routes := getRoutes()
	f := routes[r.URL.Path]
	if f == nil{
		f = handleNotFound
	}
	config.OutputPaasName(w)
	f(w, r)
}

// 启动api服务
func StartHttpServer(){
	configObj := config.GetPaasConfig()
	addr := configObj.ApiServerConfigData.Ip + ":" + tools.Float64ToString(configObj.ApiServerConfigData.Port)
	tools.Info("监听地址："+addr)
	http.HandleFunc("/", handler)
	http.ListenAndServe(addr, nil)
}