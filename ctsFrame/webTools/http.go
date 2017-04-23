package webTools

import "net/http"
import (
	"../../ctsFrame/reTools"
)

/*
http服务相关操作
@author: yubang
创建于2017年4月23日
 */

type HttpRequest struct {
	Request *http.Request
	Response http.ResponseWriter
	Session map[string]interface{}
	StatusCode int
	ResponseData []byte
}

type HttpServerInfo struct {
	RouteMap map[string]func(request HttpRequest)HttpRequest
	RewriteListList [][]string
	ErrorFuncMap map[string]func(request HttpRequest)HttpRequest
	BeforeRequest []func(request HttpRequest)(HttpRequest, bool)
	AfterRequest []func(request HttpRequest)HttpRequest
}

/*
获取一个http服务处理函数
@param routeMap: 路由map
@param rewriteListList: rewrite规则
@return http服务处理函数
 */
func getHttpHandler(httpServerInfo HttpServerInfo) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request){

		// 变量
		rewriteListList := httpServerInfo.RewriteListList
		routeMap := httpServerInfo.RouteMap

		// 标识头
		w.Header().Set("Server", "ctsFrame1.0")
		w.Header().Set("golang-server", "true")

		url := r.RequestURI

		// 处理rewrite
		for index := 0; index < len(rewriteListList);index++{
			tmpList := rewriteListList[index]
			url = reTools.Replace(tmpList[0], tmpList[1], url)
		}
		f := routeMap[url]
		if f == nil{
			w.WriteHeader(404)
			w.Write([]byte("你访问的页面已经被吃掉了！"))
			return
		}
		// 构建request对象
		request := HttpRequest{r, w, getSession(r), 200, nil}

		// 前置处理
		var continueSign bool
		for _, v := range httpServerInfo.BeforeRequest{
			request, continueSign = v(request)
			if !continueSign {
				setSession(request.Response, request.Request, request.Session)
				w.WriteHeader(request.StatusCode)
				w.Write(request.ResponseData)
				return
			}
		}

		// 调用路由函数
		request = f(request)

		// 后置处理
		for _, v := range httpServerInfo.AfterRequest {
			request = v(request)
		}

		// 输出结果
		setSession(request.Response, request.Request, request.Session)
		w.WriteHeader(request.StatusCode)
		w.Write(request.ResponseData)
	}
}

/*
启动一个http服务
@param routeMap: 路由map
@param httpAdder: 服务器监听地址
@param rewriteListList: rewrite规则
@return
 */
func StartHttpServer(httpAdder string, httpServerInfo HttpServerInfo){
	http.HandleFunc("/", getHttpHandler(httpServerInfo))
	http.ListenAndServe(httpAdder, nil)
}


/*
获取一个空路由map
return map[string]func(w http.ResponseWriter, r *http.Request)
 */
func GetBlockUrlRouteMap()map[string]func(w http.ResponseWriter, r *http.Request){
	return make(map[string]func(w http.ResponseWriter, r *http.Request))
}
