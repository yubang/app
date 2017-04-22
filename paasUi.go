package main

/*
paas平台WEB UI模块
@author: yubang
 */

import (
	"./ctsFrame/webTools"
)

func main(){
	routes := webTools.GetBlockUrlRouteMap()
	webTools.StartHttpServer(routes, "127.0.0.1:9000")
}