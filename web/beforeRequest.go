package web

import (
	"../ctsFrame/reTools"
	"../ctsFrame/stringTools"
	"../ctsFrame/webTools"
	"../ctsFrame/httpCode"
)


var beforeRequest = []webTools.BeforeRequestHandler{
	checkAdmin, handleStatic,
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

	if reTools.Match("^/web/html/.*", url){
		r.SendFile("./index.html", 0)
		return false
	}

	return true
}

func checkAdmin(obj *webTools.HttpObject)bool{
	url := stringTools.GetSplitFirstArr(obj.Request.RequestURI, "?")

	if url == "/admin/web/login" || url == "/admin/api/login"{
		return true
	}

	if reTools.Match("^/admin/web/.*", url) && obj.Session["admin"] == nil{
		obj.Redirect("/admin/web/login")
		return false
	}

	if reTools.Match("^/admin/api/.*", url) && obj.Session["admin"] == nil{
		obj.Output(httpCode.NeedLoginCode, "请登录！")
		return false
	}

	return true
}