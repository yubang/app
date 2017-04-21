package task

/*
容器分配策略模块
@author: yubang
 */

import (
	"../config"
	"../redisClient"
	"../tools"
)


// analyApp函数的一个中间算法
func getContainerPlanInFuncAnalyApp(index int64, client *redisClient.TypeRedisClient) (map[string]interface{}, string){
	// 获取容器服务器ip
	containerIp, err2 := client.LIndex(config.REDIS_KEY_CONTAINER_SERVER_IP_LIST, index).Result()
	if err2 != nil{
		return nil, ""
	}

	// 获取应用相关信息
	planJson, err3 := client.Get(config.REDIS_KEY_PLAN_CONTAIN_USE_STR + containerIp).Result()
	if err3 != nil || planJson == ""{
		return nil, ""
	}
	return tools.JsonToInterface([]byte(planJson)), containerIp
}

// 计算出一个应用的容器分配策略
func analyApp(appId string, client *redisClient.TypeRedisClient){
	// 先判断应用是否需要移除容器
	length, err := client.LLen(config.REDIS_KEY_CONTAINER_SERVER_IP_LIST).Result()
	if err != nil{
		length = 0
	}

	// 提取镜像相关信息
	image, errImage := client.HGet(config.REDIS_KEY_APP_MESSAGE_HASH + appId, "image").Result()
	if errImage != nil || image == "" {
		length = 0
	}

	// 提取内存相关信息
	memoryStr, errMemory := client.HGet(config.REDIS_KEY_APP_MESSAGE_HASH + appId, "memory").Result()
	var memory int
	if errMemory != nil || memoryStr == ""{
		length = 0
		memory = 0
	}else{
		memory = tools.StringToInt(memoryStr)
	}

	// 提取应用容器数
	planContainerNumStr, errContainerNumber := client.HGet(config.REDIS_KEY_APP_MESSAGE_HASH + appId, "minContainerNumber").Result()
	var planContainerNum int
	if errContainerNumber != nil|| planContainerNumStr == ""{
		length = 0
		planContainerNum = 0
	}else{
		planContainerNum = tools.StringToInt(planContainerNumStr)
	}

	// 移除多余的容器
	appContainerNumber := 0 // 容器数量计数器
	for index:=int64(0);index<length;index++{

		/*
		该对象结构，
		{
			appId:{image: 镜像名字, num: 容器数量, memory: 内存m},
		}
		 */
		planObj, containerIp := getContainerPlanInFuncAnalyApp(index, client)
		if planObj == nil{
			continue
		}
		if planObj != nil && planObj[appId] != nil{
			planNum := tools.Float64ToInt(planObj[appId].(map[string]interface{})["num"].(float64))
			if memory !=  tools.Float64ToInt(planObj[appId].(map[string]interface{})["memory"].(float64)) || image != planObj[appId].(map[string]interface{})["image"].(string){
				// 判断镜像名字和内存有没有变化
				delete(planObj[appId].(map[string]interface{}), appId)
			}else if planNum + appContainerNumber > planContainerNum{
				// 判断数量是否超过上限
				if appContainerNumber >= planContainerNum{
					// 释放所有容器
					delete(planObj[appId].(map[string]interface{}), appId)
				}else{
					// 释放超过部分容器
					planObj[appId].(map[string]interface{})["num"] = float64(planContainerNum - appContainerNumber)
				}
			}else{
				// 符合规定，计算当前容器数
				appContainerNumber += tools.Float64ToInt(planObj[appId].(map[string]interface{})["num"].(float64))
			}
			planJsonByte := tools.InterfaceToJson(planObj)
			// 写回redis
			client.Set(config.REDIS_KEY_PLAN_CONTAIN_USE_STR + containerIp, planJsonByte, 0)
		}

	}
	// 判断是否需要添加容器
	for index:=int64(0);index<length;index++{
		if appContainerNumber >= planContainerNum{
			// 容器数量足够
			break
		}
		// 分析该服务器有没有空余资源分配
		planObj, containerIp := getContainerPlanInFuncAnalyApp(index, client) // 计划分配资源情况
		if planObj == nil{
			continue
		}
		ableUseMemoryStr, errAbleUseMemory := client.HGet(config.REDIS_KEY_CONTAINER_SERVER_INFO_HASH + containerIp, "memory").Result() // 可使用资源
		if errAbleUseMemory != nil || ableUseMemoryStr == ""{
			continue
		}
		ableUseMemory := tools.StringToInt(ableUseMemoryStr)
		totalUseMemory := 0 // 总使用内存
		for _, v := range planObj{
			totalUseMemory += tools.Float64ToInt(v.(map[string]interface{})["memory"].(float64) * v.(map[string]interface{})["num"].(float64))
		}
		if planObj[appId] == nil{ // 初始化数据，防止空数据
			planObj[appId] = make(map[string]interface{})
			planObj[appId].(map[string]interface{})["num"] = 0
			planObj[appId].(map[string]interface{})["memory"] = memory
			planObj[appId].(map[string]interface{})["image"] = image
		}

		// 尝试添加分配资源计划
		for;;{
			if ableUseMemory + memory > totalUseMemory || appContainerNumber >= planContainerNum{
				break
			}
			planObj[appId].(map[string]interface{})["num"] = planObj[appId].(map[string]interface{})["num"].(float64) + float64(1)
			ableUseMemory += memory
			appContainerNumber += 1
		}

		if planObj[appId].(map[string]interface{})["num"] == 0{ //删除空数据
			delete(planObj, appId)
		}

		// 写回redis
		planJsonByte := tools.InterfaceToJson(planObj)
		client.Set(config.REDIS_KEY_PLAN_CONTAIN_USE_STR + containerIp, planJsonByte, 0)


	}
	// 判断容器分配是否成功完成
	if appContainerNumber < planContainerNum{
		// TODO：预警服务器没有足够资源分配容器
	}

}

// 计算一次容器分配策略
func AnalyContainer(){

	client := redisClient.GetRedisClient()
	defer client.Close()

	// 遍历所有应用
	length, err := client.LLen(config.REDIS_KEY_APP_LIST).Result()
	if err != nil{
		length = 0
	}
	for ;length > 0;length--{
		appId, err2 := client.LIndex(config.REDIS_KEY_APP_LIST, length-1).Result()
		if err2 == nil && appId != ""{
			analyApp(appId, (*redisClient.TypeRedisClient)(client)) // 处理应用分配
		}
	}
}