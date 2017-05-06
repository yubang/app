package main

import "./web"
import "./ctsFrame/cacheTools"
import "./ctsFrame/fileTools"
import "./ctsFrame/jsonTools"
import "./ctsFrame/typeConversionTools"

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

	redisObject := cacheTools.GetRedisClientObject(map[string]interface{}{
		"host": obj["Redis"].(map[string]interface{})["Host"].(string),
		"port": int(obj["Redis"].(map[string]interface{})["Port"].(float64)),
		"db": int(obj["Redis"].(map[string]interface{})["Db"].(float64)),
		"password": obj["Redis"].(map[string]interface{})["Password"].(string),
	})
	config := web.OwnConfigInfo{
		obj["TestImage"].(string),
		obj["Http"].(map[string]interface{})["Ip"].(string) + ":" + typeConversionTools.Float64ToString(obj["Http"].(map[string]interface{})["Port"].(float64)),
		redisObject,
		image,
		adminUser,
	}
	web.Init(config)

}
