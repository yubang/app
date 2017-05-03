package main

/*
paas平台代理模块
@author: yubang
创建于2017年5月3日
 */

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"./ctsFrame/jsonTools"
	"./ctsFrame/fileTools"
	"./ctsFrame/cacheTools"
	"./ctsFrame/stringTools"
	"./web"
)

func getHandler()func(http.ResponseWriter, *http.Request){

	obj := jsonTools.JsonToInterface(fileTools.ReadFromFile("./config.json"))
	cache := cacheTools.GetRedisClientObject(map[string]interface{}{
		"host": obj["Redis"].(map[string]interface{})["Host"].(string),
		"port": int(obj["Redis"].(map[string]interface{})["Port"].(float64)),
		"db": int(obj["Redis"].(map[string]interface{})["Db"].(float64)),
		"password": obj["Redis"].(map[string]interface{})["Password"].(string),
	})
	return func(w http.ResponseWriter, r *http.Request){
		host := stringTools.GetSplitFirstArr(r.Host, ":")
		client := cache.GetRedisClient()
		appId, _ := client.HGet(web.REDIS_KEY_DOMAIN_APP_HSET, host).Result()
		port, err := client.HGet(web.REDIS_KEY_APP_INFO_HSET+appId, "port").Result()

		if err != nil{
			w.WriteHeader(502)
			w.Write([]byte("请求的域名没有绑定应用！"))
			return
		}

		sourceUrl := "http://127.0.0.1:" + port
		proxyUrl, _ := url.Parse(sourceUrl)
		httpProxy := httputil.NewSingleHostReverseProxy(proxyUrl)
		httpProxy.ServeHTTP(w, r)
	}
}

func main(){
	http.HandleFunc("/", getHandler())
	http.ListenAndServe("0.0.0.0:80", nil)
}