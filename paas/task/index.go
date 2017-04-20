package task

/*
任务相关模块
@author: yubang
 */

import "../tools"
import "../config"
import (
	"../redisClient"
	"github.com/go-redis/redis"
)


// 生成一个随机任务id
func BuildTaskId() string{
	return tools.GetNowTime("%Y%m%d_") + tools.GetToken()
}

// 获取任务具体信息
func GetTaskObjFromTaskId(taskId string) map[string]interface{}{
	client := redisClient.GetRedisClient()
	taskJson, err := client.HGet(config.REDIS_KEY_TASK_HASH + tools.GetSplitFirstArr(taskId, "_"), taskId).Result()

	if err == redis.Nil{
		return nil
	}else if err != nil{
		return nil
	}

	taskObj := tools.JsonToInterface([]byte(taskJson))
	return taskObj
}