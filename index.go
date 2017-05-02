package main

import "./web"
import "./ctsFrame/cacheTools"
import "./ctsFrame/fileTools"
import "./ctsFrame/jsonTools"

func main(){

	// 读取配置文件
	t := fileTools.ReadFromFile("./config.json")
	obj := jsonTools.JsonToInterface(t)

	image := make(map[string]string)
	for k, v := range obj["Image"].(map[string]interface{}){
		image[k] = v.(string)
	}

	adminUser := web.AdminAccountStruct{
		obj["Admin"].(map[string]interface{})["Username"].(string),
		obj["Admin"].(map[string]interface{})["Password"].(string),
	}

	redisObject := cacheTools.GetRedisClientObject(map[string]interface{}{})
	config := web.OwnConfigInfo{
		":8000",
		redisObject,
		image,
		adminUser,
	}
	web.Init(config)

}
