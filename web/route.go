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
)

var routes = map[string]webTools.HttpHandler{
	"/admin/api/createApp": createApp,
	"/admin/api/appList": appList,
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

	// 检查参数
	if name == "" || desc == "" || domain == "" || git == "" || cpu == "" || memory == "" || nums == ""{
		r.Output(httpCode.ParameterMissingCode, nil)
		return
	}

	// 检查参数类型是否存在问题
	cpu = stringTools.SubString(cpu, 0, len(cpu) - 1)
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

	// 写入hset，保存应用相关信息
	hsetKey := REDIS_KEY_APP_INFO_HSET + appId
	if redisClient.HSet(hsetKey, "name", name).Err() != nil{
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}
	if redisClient.HSet(hsetKey, "desc", desc).Err() != nil{
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}
	if redisClient.HSet(hsetKey, "domain", domain).Err() != nil{
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}
	if redisClient.HSet(hsetKey, "git", git).Err() != nil{
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}
	if redisClient.HSet(hsetKey, "cpu", cpu).Err() != nil{
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}
	if redisClient.HSet(hsetKey, "memory", memory).Err() != nil{
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}
	if redisClient.HSet(hsetKey, "nums", numsInt).Err() != nil{
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}
	if redisClient.HSet(hsetKey, "image", image).Err() != nil{
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}
	// 写入一些固定默认值
	if redisClient.HSet(hsetKey, "nowImageName", "").Err() != nil{
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}
	if redisClient.HSet(hsetKey, "nowImageCreateTime", "").Err() != nil{
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}
	if redisClient.HSet(hsetKey, "nowImageStatus", "0").Err() != nil{
		// 0为没有镜像，1为镜像打包成功，2为镜像打包中，3为镜像打包失败
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}

	// 写入应用日志
	if redisClient.LPush(REDIS_KEY_APP_LOG_LIST+appId, jsonTools.InterfaceToJson(map[string]interface{}{
		"type": "success",
		"content": "应用创建成功！",
		"time": timeTools.GetNowTime("%Y-%m-%d %H:%M:%S"),
	})).Err() != nil{
		// 0为没有镜像，1为镜像打包成功，2为镜像打包中，3为镜像打包失败
		r.Output(httpCode.ServerErrorCode, nil)
		return
	}


	// 写入zset，只写入appid
	if redisClient.ZAdd(REDIS_KEY_APP_ZSET, redis.Z{float64(power), appId}).Err() != nil{
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

	result := map[string]interface{}{
		"nums": appNumber,
		"apps": arrs,
	}
	obj.Output(httpCode.OkCode, result)
}