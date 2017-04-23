package webUi

/*
paas UI程序
@author: yubang
创建于2017年4月23日
 */

import "../../ctsFrame/webTools"


func StartHttpServer(){
	webTools.StartHttpServer(routes, "127.0.0.1:8000", rewriteList)
}