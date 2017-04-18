package webApi

import "net/http"
import "../httpCode"

/*
容器相关API模块
@author: yubang
 */

// 获取一个容器操作任务
func getContainerTask(w http.ResponseWriter, r *http.Request){
	d := make(map[string]interface{})
	d["taskId"] = "123"
	d["imageName"] = "abcdef"
	output(w, httpCode.OkCode, d)
}

// 处理容器操作回调处理
func optionContainerCallback(w http.ResponseWriter, r *http.Request){

	taskId := r.FormValue("taskId")
	result := r.FormValue("result")
	host := r.FormValue("ip")
	port := r.FormValue("port")

	if taskId == "" || result == "" || host == "" || port == ""{
		output(w, httpCode.ParameterMissingCode, nil)
		return
	}

	d := make(map[string]interface{})
	output(w, httpCode.OkCode, d)
}