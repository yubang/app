package webTools

/*
HttpRequest对象方法绑定
@author: yubang
创建于2017年4月25日
 */

import (
	"../jsonTools"
	"../reTools"
	"../stringTools"
	"../httpCode"
	"../typeConversionTools"
	"fmt"
	"net/http"
	"strings"
)

/*
跳转到一个页面
@param url: 跳转地址
@return HttpRequest
 */
func (request *HttpObject)Redirect(url string){
	request.StatusCode = 302
	request.Response.Header().Set("Location", url)
	request.ResponseData = []byte("")
}

/*
输出json
@param code: HttpCode对象
@param data: 自定义数据
 */
func (request *HttpObject)Output(code httpCode.HttpCode, data interface{}){
	request.Response.Header().Set("Content-Type", "application/json")
	d := map[string]interface{}{
		"code": code.Code,
		"msg": code.Msg,
		"data": data,
	}
	request.StatusCode = 200
	request.ResponseData = jsonTools.InterfaceToJson(d)
}

/*
获取get参数
 */
func (request *HttpObject)GetArgs(key string, defaultValue string)string{
	q := request.Request.URL.Query()
	if len(q[key]) != 1{
		return defaultValue
	}
	return q[key][0]
}

/*
获取请求完整地址
@param r: HttpRequest
@return string
 */
func (requestObject *HttpObject)GetRequestUrl()string{
	scheme := "http://"
	if requestObject.Request.TLS != nil {
		scheme = "https://"
	}
	return strings.Join([]string{scheme, requestObject.Request.Host, requestObject.Request.RequestURI}, "")
}

/*
获取加上域名地址
@param r: HttpRequest
@return string
 */
func (requestObject *HttpObject)GetNewUrl(path string)string{
	scheme := "http://"
	if requestObject.Request.TLS != nil {
		scheme = "https://"
	}
	return strings.Join([]string{scheme, requestObject.Request.Host, path}, "")
}

/*
启动一个http服务
@param routeMap: 路由map
@param httpAdder: 服务器监听地址
@param rewriteListList: rewrite规则
@return
 */
func (httpServer *HttpServerInfo)StartHttpServer(httpAdder string){
	http.HandleFunc("/", httpServer.getHttpHandler())
	fmt.Print("http服务器启动，监听地址：" + httpAdder + "\n")
	http.ListenAndServe(httpAdder, nil)
}


/*
获取一个http服务处理函数
@param routeMap: 路由map
@param rewriteListList: rewrite规则
@return http服务处理函数
 */
func (httpServerInfo *HttpServerInfo)getHttpHandler() func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request){
		// 变量
		rewriteListList := httpServerInfo.RewriteListList
		routeMap := httpServerInfo.RouteMap

		// 标识头
		w.Header().Set("Server", "ctsFrame1.0")
		w.Header().Set("golang-server", "true")

		url := stringTools.GetSplitFirstArr(r.RequestURI, "?")

		// 处理rewrite
		for index := 0; index < len(rewriteListList);index++{
			tmpList := rewriteListList[index]
			url = reTools.Replace(tmpList[0], tmpList[1], url)
		}
		f := routeMap[url]
		if f == nil{
			w.WriteHeader(404)
			w.Write([]byte("你访问的页面已经被吃掉了！"))
			fmt.Print(r.Method + " " + r.RequestURI + "【404】\n")
			return
		}
		// 构建request对象
		request := HttpObject{r, w, getSession(r, httpServerInfo.CacheClient), 200, nil, httpServerInfo.CacheClient}

		// 前置处理
		var continueSign bool
		for _, v := range httpServerInfo.BeforeRequest{
			continueSign = v(&request)
			if !continueSign {
				setSession(&request, request.Session)
				w.WriteHeader(request.StatusCode)
				w.Write(request.ResponseData)
				fmt.Print(r.Method + " " + r.RequestURI + "【" + typeConversionTools.IntToString(request.StatusCode) + "】\n")
				return
			}
		}

		// 调用路由函数
		f(&request)

		// 后置处理
		for _, v := range httpServerInfo.AfterRequest {
			v(&request)
		}

		// 输出结果
		setSession(&request, request.Session)
		w.WriteHeader(request.StatusCode)
		w.Write(request.ResponseData)

		fmt.Print(r.Method + " " + r.RequestURI + "【" + typeConversionTools.IntToString(request.StatusCode) + "】\n")

	}
}