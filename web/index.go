package web

import "../ctsFrame/webTools"
import "../ctsFrame/cacheTools"

func Init(ownObj OwnConfigInfo){

	cache := cacheTools.CacheClient{"session_", ownObj.RedisObject}

	obj := webTools.HttpServerInfo{
		routes,
		[][]string{},
		map[int]webTools.ErrorHandler{},
		beforeRequest,
		afterRequest,
		cache,
		&ownObj,
	}
	obj.StartHttpServer(ownObj.HttpAddr)
}