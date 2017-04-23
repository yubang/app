package webUi

/*
paas UI程序
@author: yubang
创建于2017年4月23日
 */

import "../../ctsFrame/webTools"


func StartHttpServer(){
	errFuncMap := make(map[string]func(request webTools.HttpRequest)webTools.HttpRequest)
	beforeRequest := []func(request webTools.HttpRequest)(webTools.HttpRequest, bool){}
	afterRequest := []func(request webTools.HttpRequest)webTools.HttpRequest{}
	httpServerInfo := webTools.HttpServerInfo{routes, rewriteList, errFuncMap, beforeRequest, afterRequest}
	webTools.StartHttpServer("127.0.0.1:8000", httpServerInfo)
}