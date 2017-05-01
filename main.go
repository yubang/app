package main

import "./web"
import "./ctsFrame/cacheTools"

func main(){
	redisObject := cacheTools.GetRedisClientObject(map[string]interface{}{})
	config := web.OwnConfigInfo{
		":8000",
		redisObject,
	}
	web.Init(config)

}
