package webApi


/*
API http服务模块
@author: yubang
 */

import (
	"../../paas/config"
	"../../paas/tools"
	"net/http"
	"../../paas/httpCode"
)

// 输出结果
func output(w http.ResponseWriter, code httpCode.HttpCode, data interface{}){
	response := make(map[string]interface{})
	response["code"] = code.Code
	response["msg"] = code.Msg
	response["data"] = data
	r := tools.InterfaceToJson(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(r)
}

// 没有匹配到路有时候处理方法
func handleNotFound(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(404)
	w.Write([]byte("not found!"))
}

// api处理入口函数
func handler(w http.ResponseWriter, r *http.Request) {
	config.OutputPaasName(w)

	// 请求API必须是post方法
	if r.Method != "POST"{
		w.WriteHeader(405)
		w.Write([]byte("method not allow!"))
		return
	}

	// 判断token是否正确
	token := r.FormValue("token")
	configObj := config.GetPaasConfig()
	if token != configObj.Token{
		output(w, httpCode.TokenErrorCode, nil)
		return
	}

	// 获取路由配置
	routes := getRoutes()
	f := routes[r.URL.Path]
	if f == nil{
		f = handleNotFound
	}
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