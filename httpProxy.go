package main

import (
	"./paas/tools"
	"./paas/proxy"
	"net/http"
	"./paas/config"
	"./paas/redisClient"
	"runtime"
)

// 处理http proxy
func proxyHandler(w http.ResponseWriter, r *http.Request){
	host := r.Host
	host = tools.GetSplitFirstArr(host, ":")
	tools.Debug("请求的域名：" + host)

	// 获取一个后端地址
	client := redisClient.GetRedisClient()
	defer client.Close()

	// 根据域名获取应用
	appId, err := client.HGet(config.REDIS_KEY_PROXY_DOMAIN_APP_HASH, host).Result()
	if err != nil{
		w.WriteHeader(403)
		w.Write([]byte("该域名未绑定应用！"))
		return
	}

	// 根据应用获取一个后端服务
	proxyStr, err2 := client.SRandMember(config.REDIS_KEY_PROXY_APP_CONTAINER_IP_AND_PORT_SET + appId).Result()
	if err2 != nil{
		w.WriteHeader(403)
		w.Write([]byte("该应用没有绑定容器！"))
		return
	}

	proxy.ProxyHttp(proxyStr, w, r)
}

func main(){

	// 使用多核CPU
	runtime.GOMAXPROCS(runtime.NumCPU())

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