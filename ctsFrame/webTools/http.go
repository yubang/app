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

/*
获取一个http服务处理函数
@param routeMap: 路由map
@param rewriteListList: rewrite规则
@return http服务处理函数
 */
func getHttpHandler(routeMap map[string]func(w http.ResponseWriter, r *http.Request), rewriteListList [][]string) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request){

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
		f(w, r)
	}
}

/*
启动一个http服务
@param routeMap: 路由map
@param httpAdder: 服务器监听地址
@param rewriteListList: rewrite规则
@return
 */
func StartHttpServer(routeMap map[string]func(w http.ResponseWriter, r *http.Request), httpAdder string, rewriteListList [][]string){
	http.HandleFunc("/", getHttpHandler(routeMap, rewriteListList))
	http.ListenAndServe(httpAdder, nil)
}


/*
获取一个空路由map
return map[string]func(w http.ResponseWriter, r *http.Request)
 */
func GetBlockUrlRouteMap()map[string]func(w http.ResponseWriter, r *http.Request){
	return make(map[string]func(w http.ResponseWriter, r *http.Request))
}
