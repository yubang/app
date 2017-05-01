package web

import "../ctsFrame/webTools"
import "../ctsFrame/cacheTools"

func Init(){

	redisObject := cacheTools.GetRedisClientObject(map[string]interface{}{})
	cache := cacheTools.CacheClient{"session_", redisObject}

	obj := webTools.HttpServerInfo{
		routes,
		[][]string{},
		map[int]webTools.ErrorHandler{},
		beforeRequest,
		afterRequest,
		cache,
		nil,
	}
	obj.StartHttpServer(":8000")
}