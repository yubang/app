package webApi

import "net/http"
import "../httpCode"

/*
与git，docker镜像打包相关模块
@author: yubang
 */

// 拉取一条打包镜像任务
func getAGitPullTask(w http.ResponseWriter, r *http.Request){
	d := make(map[string]interface{})
	d["taskId"] = "1"
	d["gitUrl"] = "git@github.com:yubang/navigation.git"
	d["dockerImage"] = "static"
	output(w, httpCode.OkCode, d)
}

// 打包镜像后通知
func buildImageCallback(w http.ResponseWriter, r *http.Request){

	taskId := r.FormValue("taskId")
	imageName := r.FormValue("imageName")
	result := r.FormValue("result")

	if taskId == "" || imageName == "" || result == ""{
		output(w, httpCode.ParameterMissingCode, nil)
		return
	}

	d := make(map[string]interface{})
	output(w, httpCode.OkCode, d)
}