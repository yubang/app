package web

import (
	"../ctsFrame/reTools"
	"../ctsFrame/stringTools"
	"../ctsFrame/webTools"
)


var beforeRequest = []webTools.BeforeRequestHandler{
	handleStatic,
}

func handleStatic(r *webTools.HttpObject)bool{

	url := stringTools.GetSplitFirstArr(r.Request.RequestURI, "?")

	if reTools.Match("^/static/.*", url){
		r.SendFile("." + url, 0)
		return false
	}

	if reTools.Match("^/admin/web/.*", url){
		r.SendFile("./index.html", 0)
		return false
	}

	return true
}