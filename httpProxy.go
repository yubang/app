package main

import (
	"./paas/tools"
	"net/http"
)

// 处理http proxy
func proxyHandler(w http.ResponseWriter, r *http.Request){
	host := r.Host
	tools.Debug("请求的域名：" + host)
	w.Write([]byte("测试"))
}

func main(){
	// 读取配置文件
	jsonString := tools.ReadFromFile("./config.json")
	configObj := tools.JsonToInterface(jsonString)

	if configObj == nil{
		tools.Error("配置文件格式有误！")
		return
	}

	httpProxy := configObj["httpProxy"].(map[string]interface{})
	addr := httpProxy["ip"].(string) + ":" + tools.Float64ToString(httpProxy["port"].(float64))

	tools.Info("监听地址：" + addr)

	http.HandleFunc("/", proxyHandler)
	http.ListenAndServe(addr, nil)
}