package main

/*
paas平台WEB UI模块
@author: yubang
 */

import (
	"./ctsFrame/webTools"
	"fmt"
)

func main(){
	fmt.Print(webTools.GetContentTypeFromName(".html"))
}