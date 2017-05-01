package webTools

import "net/http"
import (
	"../../ctsFrame/cacheTools"
)

/*
http服务相关操作
@author: yubang
创建于2017年4月23日
 */

type HttpObject struct {
	Request *http.Request
	Response http.ResponseWriter
	Session map[string]interface{}
	StatusCode int
	ResponseData []byte
	CacheClient cacheTools.CacheClient
	OwnObj interface{} // 自己的对象，生命周期是进程启动到结束
}

type HttpServerInfo struct {
	RouteMap map[string]HttpHandler  // 路由map对象
	RewriteListList [][]string  // rewrite规则列表
	ErrorFuncMap map[int]ErrorHandler  // 错误路由处理
	BeforeRequest []BeforeRequestHandler  // 请求前拦截，返回false终止请求
	AfterRequest []AfterRequestHandler  // 请求后处理
	CacheClient cacheTools.CacheClient  // cache对象
	OwnObj interface{} // 自己的对象，生命周期是进程启动到结束
}

type HttpHandler func(*HttpObject)
type ErrorHandler func(*HttpObject)
type BeforeRequestHandler func(request *HttpObject)bool
type AfterRequestHandler func(request *HttpObject)
