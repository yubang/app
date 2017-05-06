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
	"./web/docker"
	"time"
)

func buildDockerfile(dirPath, baseImage string)bool{
	t := "FROM "+baseImage+"\n"
	t += "MAINTAINER paas（yubang93@gmail.com）\n"
	t += "ADD web /var/web\n"
	t += "RUN /bin/bash /var/install.sh\n"
	return fileTools.WriteNewFile(dirPath+"/Dockerfile", []byte(t))
}

func buildImage(appId string, imageName string, cacheObj cacheTools.RedisClientObject, imageMap map[string]interface{}, imageUrl string)bool{

	redisClient := cacheObj.GetRedisClient()

	// 获取必须信息
	image, err2 := redisClient.HGet(web.REDIS_KEY_APP_INFO_HSET+appId, "image").Result()
	if err2 != nil{
		return false
	}

	gitUrl, err2 := redisClient.HGet(web.REDIS_KEY_APP_INFO_HSET+appId, "git").Result()
	if err2 != nil || gitUrl == ""{
		return false
	}

	// 获取基础镜像name
	baseImage := imageMap[image]
	if baseImage == nil{
		return false
	}

	// 创建临时文件夹和Dockerfile文件
	dirPath, err := fileTools.MakeTempDir()
	if err != nil{
		return false
	}
	defer fileTools.RemoveDir(dirPath)

	shellClient := docker.ShellStruct{cacheObj}
	// clone代码
	if shellClient.ExecShell("git clone --depth=1 " + gitUrl + " " + dirPath + "/web") == nil{
		return false
	}

	// 生成Docker文件
	if !buildDockerfile(dirPath, baseImage.(string)){
		return false
	}

	// 生成docker镜像
	if shellClient.ExecShell("docker build -t  " + imageName + " " + dirPath) == nil{
		return false
	}

	// 提交镜像到仓库
	if shellClient.ExecShell("docker tag "+imageName+" "+imageUrl+"/"+imageName) == nil{
		return false
	}
	if shellClient.ExecShell("docker push "+imageUrl+"/"+imageName) == nil{
		return false
	}
	return true
}

func updateMessage(appId string, imageName string, imageAbout string, cacheObj cacheTools.RedisClientObject, imageUrl string)bool{
	redisClient := cacheObj.GetRedisClient()

	imageCreateTime := timeTools.GetNowTime("%Y-%m-%d %H:%M:%S")
	image := jsonTools.InterfaceToJson(map[string]interface{}{
		"imageName": imageUrl + "/" +imageName,
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

func handleTask(cacheObj cacheTools.RedisClientObject, imageMap map[string]interface{}, imageUrl string)bool{
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
	if buildImage(appId, imageName, cacheObj, imageMap, imageUrl){
		sign = updateMessage(appId, imageName, imageAbout, cacheObj, imageUrl)
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

	// 读取ssh公钥
	shellClient := docker.ShellStruct{cache}
	sshContent := shellClient.ExecShell("cat ~/.ssh/id_rsa.pub")
	cache.GetRedisClient().Set(web.REDIS_KEY_SSH_STR, sshContent, 0)

	for ;true;{
		if !handleTask(cache, obj["Image"].(map[string]interface{}), obj["ImageUrl"].(string)) {
			// 出错时候休眠3秒
			time.Sleep(time.Second * 3)
		}

	}

}