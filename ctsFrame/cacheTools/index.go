package cacheTools

/*
简单缓存模块
@author: yubang
创建于2017年4月23日
 */

import "../timeTools"

var cacheMap = make(map[string]interface{})

func Set(key string, value interface{}){
	data := map[string]interface{}{
		"data": value,
		"expirTime": 0,
	}
	cacheMap[key] = data
}

func Get(key string)interface{}{
	if Exist(key){
		return cacheMap[key].(map[string]interface{})["data"]
	}
	return nil
}

func Remove(key string){
	delete(cacheMap, key)
}

func Exist(key string) bool{
	if cacheMap[key] != nil{
		return false
	}
	if cacheMap[key].(map[string]interface{})["expirTime"].(int64) <= timeTools.GetNowTimeSecond(){
		Remove(key)
		return false
	}
	return true
}

func Expir(key string, timeout int){
	if Exist(key){
		cacheMap[key].(map[string]interface{})["expirTime"] = timeTools.GetNowTimeSecond() + int64(timeout)
	}
}
