package web

import "../ctsFrame/webTools"
import "../ctsFrame/httpCode"
import "../ctsFrame/stringTools"
import "../ctsFrame/typeConversionTools"
import (
	"../ctsFrame/timeTools"
	"github.com/go-redis/redis"
	"../ctsFrame/jsonTools"
	"../ctsFrame/utilTools"
	"math/rand"
)

var routes = map[string]webTools.HttpHandler{
	"/admin/api/createApp": createApp,
	"/admin/api/appList": appList,
	"/admin/api/deleteApp": deleteApp,
	"/admin/api/appInfo": appInfo,
	"/admin/api/updateAppContainerInfo": updateAppContainerInfo,
	"/admin/api/buildImage": buildImage,
}

func createApp(r *webTools.HttpObject){

	name := r.Request.FormValue("name")
	desc := r.Request.FormValue("desc")
	domain := r.Request.FormValue("domain")
	git := r.Request.FormValue("git")
	cpu := r.Request.FormValue("cpu")
	memory := r.Request.FormValue("memory")
	nums := r.Request.FormValue("nums")
	image := r.Request.FormValue("image")

	testImageName := "paas_test"
	testImageCreateTime := timeTools.GetNowTime("%Y-%m-%d %H:%M:%S")
	testImageAbout := "paas平台初始化测试镜像"

	// 检查参数
	if name == "" || desc == "" || domain == "" || git == "" || cpu == "" || memory == "" || nums == ""{
		r.Output(httpCode.ParameterMissingCode, nil)
		return
	}

	// 检查参数类型是否存在问题
	cpu = stringTools.SubString(cpu, 0, (len(cpu)/2) - 1)
	memory = stringTools.SubString(memory, 0, len(memory) - 1)
	numsInt, err := typeConversionTools.StringToInt(nums)
	if err != nil{
		r.Output(httpCode.ParameterMissingCode, nil)
		return
	}

	// 生成唯一appid
	appId := utilTools.GetToken32()

	// 获取当前时间戳用于zset排序
	power := timeTools.GetNowTimeSecond()

	// 获取redis实例
	redisClient := r.OwnObj.(*OwnConfigInfo).RedisObject.GetRedisClient()

	// 获取一个可用端口
	port := rand.Intn(10000) + 30000
	portSign := false
	for index:=0;index<10000;index++{
		i, _ := redisClient.SAdd(REDIS_KEY_POST_USE, port).Result()
		if i != 0{
			portSign = true
			break
		}
		port++
	}

	if !portSign{
		r.Output(httpCode.MessageErrorCode, "没有可用端口！")
		return
	}

	// 关联应用和域名
	if redisClient.HGet(REDIS_KEY_DOMAIN_APP_HSET, domain).Err() != redis.Nil{
		redisClient.SRem(REDIS_KEY_POST_USE, port) // 删除端口占用
		r.Output(httpCode.MessageErrorCode, "域名已经被绑定！")
		return
	}
	if redisClient.HSet(REDIS_KEY_APP_DOMAIN_HSET, appId, domain).Err() != nil{
		redisClient.SRem(REDIS_KEY_POST_USE, port) // 删除端口占用
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}
	if redisClient.HSet(REDIS_KEY_DOMAIN_APP_HSET, domain, appId).Err() != nil{
		redisClient.SRem(REDIS_KEY_POST_USE, port) // 删除端口占用
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}

	// 写入hset，保存应用相关信息
	hsetKey := REDIS_KEY_APP_INFO_HSET + appId
	if redisClient.HSet(hsetKey, "name", name).Err() != nil{
		redisClient.SRem(REDIS_KEY_POST_USE, port) // 删除端口占用
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}
	if redisClient.HSet(hsetKey, "desc", desc).Err() != nil{
		redisClient.SRem(REDIS_KEY_POST_USE, port) // 删除端口占用
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}
	if redisClient.HSet(hsetKey, "domain", domain).Err() != nil{
		redisClient.SRem(REDIS_KEY_POST_USE, port) // 删除端口占用
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}
	if redisClient.HSet(hsetKey, "git", git).Err() != nil{
		redisClient.SRem(REDIS_KEY_POST_USE, port) // 删除端口占用
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}
	if redisClient.HSet(hsetKey, "cpu", cpu).Err() != nil{
		redisClient.SRem(REDIS_KEY_POST_USE, port) // 删除端口占用
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}
	if redisClient.HSet(hsetKey, "memory", memory).Err() != nil{
		redisClient.SRem(REDIS_KEY_POST_USE, port) // 删除端口占用
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}
	if redisClient.HSet(hsetKey, "nums", numsInt).Err() != nil{
		redisClient.SRem(REDIS_KEY_POST_USE, port) // 删除端口占用
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}
	if redisClient.HSet(hsetKey, "image", image).Err() != nil{
		redisClient.SRem(REDIS_KEY_POST_USE, port) // 删除端口占用
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}
	if redisClient.HSet(hsetKey, "port", port).Err() != nil{
		redisClient.SRem(REDIS_KEY_POST_USE, port) // 删除端口占用
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}
	// 写入一些固定默认值
	if redisClient.HSet(hsetKey, "nowImageName", testImageName).Err() != nil{
		redisClient.SRem(REDIS_KEY_POST_USE, port) // 删除端口占用
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}
	if redisClient.HSet(hsetKey, "nowImageCreateTime", testImageCreateTime).Err() != nil{
		redisClient.SRem(REDIS_KEY_POST_USE, port) // 删除端口占用
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}
	if redisClient.HSet(hsetKey, "nowImageAbount", testImageAbout).Err() != nil{
		redisClient.SRem(REDIS_KEY_POST_USE, port) // 删除端口占用
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}
	if redisClient.HSet(hsetKey, "nowImageStatus", "1").Err() != nil{
		// 0为没有镜像，1为镜像打包成功，2为镜像打包中，3为镜像打包失败
		redisClient.SRem(REDIS_KEY_POST_USE, port) // 删除端口占用
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}

	// 写入应用日志
	if redisClient.LPush(REDIS_KEY_APP_LOG_LIST+appId, jsonTools.InterfaceToJson(map[string]interface{}{
		"type": "success",
		"content": "应用创建成功！",
		"time": timeTools.GetNowTime("%Y-%m-%d %H:%M:%S"),
	})).Err() != nil{
		redisClient.SRem(REDIS_KEY_POST_USE, port) // 删除端口占用
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}

	// 默认使用测试镜像
	newImage := jsonTools.InterfaceToJson(map[string]interface{}{
		"imageName": testImageName,
		"imageCreateTime": testImageCreateTime,
		"imageAbout": testImageAbout,
	})
	if redisClient.LPush(REDIS_KEY_APP_IMAGE_LIST+appId, newImage).Err() != nil{
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}


	// 写入zset，只写入appid
	if redisClient.ZAdd(REDIS_KEY_APP_ZSET, redis.Z{float64(power), appId}).Err() != nil{
		redisClient.SRem(REDIS_KEY_POST_USE, port) // 删除端口占用
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}

	r.Output(httpCode.OkCode, nil)

}

func appList(obj *webTools.HttpObject){

	page := obj.Request.FormValue("page")
	if page == ""{
		page = "1"
	}

	pageNumber, err := typeConversionTools.StringToInt(page)
	if err != nil{
		obj.Output(httpCode.ParameterMissingCode, nil)
		return
	}

	redisClient := obj.OwnObj.(*OwnConfigInfo).RedisObject.GetRedisClient()
	// 获取应用数量
	appNumber, err2 := redisClient.ZCard(REDIS_KEY_APP_ZSET).Result()
	if err2 != nil{
		obj.Output(httpCode.ServerErrorCode, nil)
		return
	}
	// 获取一页应用信息
	offset := 5
	index := (pageNumber - 1) * 5
	arrs, err3 := redisClient.ZRevRange(REDIS_KEY_APP_ZSET, int64(index), int64(index + offset)).Result()
	if err3 != nil{
		obj.Output(httpCode.ServerErrorCode, nil)
		return
	}

	apps := make([]map[string]string, len(arrs))

	for index, v := range arrs{
		appInfo, err4 := redisClient.HGetAll(REDIS_KEY_APP_INFO_HSET + v).Result()
		if err4 != nil{
			obj.Output(httpCode.ServerErrorCode, nil)
			return
		}
		apps[index] = map[string]string{
			"appId": v,
			"name": appInfo["name"],
			"desc": appInfo["desc"],
			"image": appInfo["image"],
			"nums": appInfo["nums"],
			"cpu": appInfo["cpu"],
			"memory": appInfo["memory"],
			"git": appInfo["git"],
			"domain": appInfo["domain"],
			"nowImageName": appInfo["nowImageName"] + "" + appInfo["nowImageCreateTime"],
		}
	}

	result := map[string]interface{}{
		"nums": appNumber,
		"apps": apps,
	}
	obj.Output(httpCode.OkCode, result)
}

func deleteApp(obj *webTools.HttpObject){

	appId := obj.Request.FormValue("appId")

	if appId == ""{
		obj.Output(httpCode.ParameterMissingCode, -1)
		return
	}

	// 删除应用信息
	redisClient := obj.OwnObj.(*OwnConfigInfo).RedisObject.GetRedisClient()
	if redisClient.Del(REDIS_KEY_APP_INFO_HSET + appId).Err() != nil{
		obj.Output(httpCode.ServerErrorCode, -2)
		return
	}

	// 删除应用日志
	if redisClient.Del(REDIS_KEY_APP_LOG_LIST + appId).Err() != nil{
		obj.Output(httpCode.ServerErrorCode, -3)
		return
	}

	// 删除域名
	domain, _ := redisClient.HGet(REDIS_KEY_APP_DOMAIN_HSET, appId).Result()
	if redisClient.HDel(REDIS_KEY_DOMAIN_APP_HSET, domain).Err() != nil{
		obj.Output(httpCode.ServerErrorCode, -4)
		return
	}
	if redisClient.HDel(REDIS_KEY_APP_DOMAIN_HSET, appId).Err() != nil{
		obj.Output(httpCode.ServerErrorCode, -5)
		return
	}

	// 删除端口占用
	port, err := redisClient.HGet(REDIS_KEY_APP_INFO_HSET + appId, "port").Result()
	if err == nil{
		if redisClient.SRem(REDIS_KEY_POST_USE, port).Err() != nil{
			obj.Output(httpCode.ServerErrorCode, -7)
			return
		}
	}


	// 删除应用镜像
	// todo: 删除仓库镜像
	if redisClient.Del(REDIS_KEY_APP_IMAGE_LIST+appId).Err() != nil{
		obj.Output(httpCode.ServerErrorCode, -8)
		return
	}

	// 从应用列表移除
	if redisClient.ZRem(REDIS_KEY_APP_ZSET, appId).Err() != nil{
		obj.Output(httpCode.ServerErrorCode, -9)
		return
	}

	obj.Output(httpCode.OkCode, "移除应用成功！")
}

func appInfo(obj *webTools.HttpObject){

	appId := obj.Request.FormValue("appId")

	redisClient := obj.OwnObj.(*OwnConfigInfo).RedisObject.GetRedisClient()

	// 获取应用相关信息
	appInfo, err := redisClient.HGetAll(REDIS_KEY_APP_INFO_HSET + appId).Result()
	if err != nil{
		obj.Output(httpCode.ServerErrorCode, nil)
		return
	}

	// 拉取所有应用日志
	logs, err2 := redisClient.LRange(REDIS_KEY_APP_LOG_LIST+appId, 0, -1).Result()
	if err2 != nil{
		obj.Output(httpCode.ServerErrorCode, nil)
		return
	}

	logList := make([]map[string]interface{}, len(logs))
	for index, v := range logs{
		logList[index] = jsonTools.JsonToInterface([]byte(v))
	}

	// 拉取应用所有打包镜像
	images, err3 := redisClient.LRange(REDIS_KEY_APP_IMAGE_LIST+appId, 0, -1).Result()
	if err3 != nil{
		obj.Output(httpCode.ServerErrorCode, nil)
		return
	}

	result := map[string]interface{}{
		"appInfo": appInfo,
		"appId": appId,
		"logs": logList,
		"images": images,
	}

	obj.Output(httpCode.OkCode, result)
}

func updateAppContainerInfo(obj *webTools.HttpObject){

	appId := obj.Request.FormValue("appId")
	nums := obj.Request.FormValue("nums")
	cpu := obj.Request.FormValue("cpu")
	memory := obj.Request.FormValue("memory")

	redisClient := obj.OwnObj.(*OwnConfigInfo).RedisObject.GetRedisClient()

	// 检查参数是否完整
	if appId == "" || nums == "" || cpu == "" || memory == ""{
		obj.Output(httpCode.ParameterMissingCode, nil)
		return
	}

	// 判断参数类型是否正确
	if !typeConversionTools.IsNumber(nums) || !typeConversionTools.IsNumber(cpu) || !typeConversionTools.IsNumber(memory){
		obj.Output(httpCode.ParameterMissingCode, nil)
		return
	}

	// 更新信息
	if redisClient.HSet(REDIS_KEY_APP_INFO_HSET+appId, "cpu", cpu).Err() != nil{
		obj.Output(httpCode.ServerErrorCode, nil)
		return
	}
	if redisClient.HSet(REDIS_KEY_APP_INFO_HSET+appId, "memory", memory).Err() != nil{
		obj.Output(httpCode.ServerErrorCode, nil)
		return
	}
	if redisClient.HSet(REDIS_KEY_APP_INFO_HSET+appId, "nums", nums).Err() != nil{
		obj.Output(httpCode.ServerErrorCode, nil)
		return
	}

	// todo: 调用docker service update

	// 记录日志
	log := jsonTools.InterfaceToJson(map[string]interface{}{
		"type": "info",
		"content": "应用配置变更为，容器：" + nums + "个，cpu："+ cpu + "核，内存："+memory+"M",
		"time": timeTools.GetNowTime("%Y-%m-%d %H:%M:%S"),
	})
	redisClient.LPush(REDIS_KEY_APP_LOG_LIST+appId, log)

	obj.Output(httpCode.OkCode, "更新应用配置成功！")
}

// 打包镜像
func buildImage(obj *webTools.HttpObject){

	appId := obj.Request.FormValue("appId")

	// todo: 拉取代码，生成docker image
	imageName := utilTools.GetToken32()
	imageCreateTime := timeTools.GetNowTime("%Y-%m-%d %H:%M:%S")
	imageAbout := obj.Request.FormValue("imageAbout")
	image := jsonTools.InterfaceToJson(map[string]interface{}{
		"imageName": imageName,
		"imageCreateTime": imageCreateTime,
		"imageAbout": imageAbout,
	})

	// 更新信息
	redisClient := obj.OwnObj.(*OwnConfigInfo).RedisObject.GetRedisClient()
	if redisClient.LPush(REDIS_KEY_APP_IMAGE_LIST+appId, image).Err() != nil{
		obj.Output(httpCode.ServerErrorCode, nil)
		return
	}

	if redisClient.HSet(REDIS_KEY_APP_INFO_HSET+appId, "nowImageStatus", 2).Err() != nil{
		obj.Output(httpCode.ServerErrorCode, nil)
		return
	}

	// 记录日志
	log := jsonTools.InterfaceToJson(map[string]interface{}{
		"type": "info",
		"content": "申请构建镜像！",
		"time": timeTools.GetNowTime("%Y-%m-%d %H:%M:%S"),
	})
	redisClient.LPush(REDIS_KEY_APP_LOG_LIST+appId, log)

	obj.Output(httpCode.OkCode, "打包镜像操作已经提交！")

}