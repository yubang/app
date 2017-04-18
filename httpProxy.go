package main

import (
	"./paas/tools"
	"./paas/proxy"
	"net/http"
	"./paas/config"
)

// 处理http proxy
func proxyHandler(w http.ResponseWriter, r *http.Request){
	host := r.Host
	tools.Debug("请求的域名：" + host)
	proxy.ProxyHttp("127.0.0.5", 9000, w, r)
}

func main(){
	// 读取配置文件
	var configObj config.PassConfig
	configObj = config.GetPaasConfig()

	if configObj.Err{
		tools.Error("配置文件格式有误！")
		return
	}

	addr := configObj.HttpProxyConfigData.Ip + ":" + tools.Float64ToString(configObj.HttpProxyConfigData.Port)

	tools.Info("监听地址：" + addr)

	http.HandleFunc("/", proxyHandler)
	http.ListenAndServe(addr, nil)
}