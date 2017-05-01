package web

import "../ctsFrame/webTools"
import "../ctsFrame/stringTools"
import "../ctsFrame/reTools"


var beforeRequest = []webTools.BeforeRequestHandler{
	handleStatic,
}

func handleStatic(r *webTools.HttpObject)bool{

	url := stringTools.GetSplitFirstArr(r.Request.RequestURI, "?")

	if reTools.Match("^/static/.*", url){
		r.SendFile("." + url, 0)
		return false
	}

	return true
}