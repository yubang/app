package main

/*
打包镜像模块
@author: yubang
创建于2017年5月4日
 */

import "./ctsFrame/timeTools"
import "./ctsFrame/jsonTools"
import "./ctsFrame/cacheTools"
import "./ctsFrame/fileTools"
import "./ctsFrame/utilTools"
import (
	"./web"
	"time"
)

func buildImage()bool{

	// todo: 拉取代码，生成docker image
	return true
}

func updateMessage(appId string, imageName string, imageAbout string, cacheObj cacheTools.RedisClientObject)bool{
	redisClient := cacheObj.GetRedisClient()

	imageCreateTime := timeTools.GetNowTime("%Y-%m-%d %H:%M:%S")
	image := jsonTools.InterfaceToJson(map[string]interface{}{
		"imageName": imageName,
		"imageCreateTime": imageCreateTime,
		"imageAbout": imageAbout,
	})

	// 更新信息
	if redisClient.LPush(web.REDIS_KEY_APP_IMAGE_LIST+appId, image).Err() != nil{
		return false
	}

	if redisClient.HSet(web.REDIS_KEY_APP_INFO_HSET+appId, "nowImageStatus", 1).Err() != nil{
		return false
	}
	return true
}

func signErrorImage(appId string, cacheObj cacheTools.RedisClientObject)bool{
	redisClient := cacheObj.GetRedisClient()
	if redisClient.HSet(web.REDIS_KEY_APP_INFO_HSET+appId, "nowImageStatus", 3).Err() != nil{
		return false
	}
	return true
}

func handleTask(cacheObj cacheTools.RedisClientObject)bool{
	redisClient := cacheObj.GetRedisClient()
	arrs, err := redisClient.BLPop(0, web.REDIS_KEY_BUILD_IMAGE_TASK_LIST).Result()
	if err != nil{
		return false
	}

	obj := jsonTools.JsonToInterface([]byte(arrs[1]))
	if obj == nil{
		return false
	}

	imageName := utilTools.GetToken32()
	imageAbout := obj["imageAbout"].(string)
	appId := obj["appId"].(string)
	var sign bool
	var log []byte
	if buildImage(){
		sign = updateMessage(appId, imageName, imageAbout, cacheObj)
		log = jsonTools.InterfaceToJson(map[string]interface{}{
			"type": "success",
			"content": "构建镜像成功！",
			"time": timeTools.GetNowTime("%Y-%m-%d %H:%M:%S"),
		})
	}else{
		sign = signErrorImage(appId, cacheObj)
		log = jsonTools.InterfaceToJson(map[string]interface{}{
			"type": "error",
			"content": "构建镜像失败！",
			"time": timeTools.GetNowTime("%Y-%m-%d %H:%M:%S"),
		})
	}

	cacheObj.GetRedisClient().LPush(web.REDIS_KEY_APP_LOG_LIST+appId, log)
	return sign
}

func main(){
	obj := jsonTools.JsonToInterface(fileTools.ReadFromFile("./config.json"))
	cache := cacheTools.GetRedisClientObject(map[string]interface{}{
		"host": obj["Redis"].(map[string]interface{})["Host"].(string),
		"port": int(obj["Redis"].(map[string]interface{})["Port"].(float64)),
		"db": int(obj["Redis"].(map[string]interface{})["Db"].(float64)),
		"password": obj["Redis"].(map[string]interface{})["Password"].(string),
	})
	for ;true;{
		if !handleTask(cache) {
			// 出错时候休眠3秒
			time.Sleep(time.Second * 3)
		}

	}

}