package webTools

import "net/http"

/*
http服务相关操作
@author: yubang
创建于2017年4月23日
 */

/*
获取一个http服务处理函数
@param routeMap: 路由map
@return http服务处理函数
 */
func getHttpHandler(routeMap map[string]func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request){

		// 标识头
		w.Header().Set("Server", "ctsFrame1.0")
		w.Header().Set("golang-server", "true")

		f := routeMap[r.RequestURI]
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
@return
 */
func StartHttpServer(routeMap map[string]func(w http.ResponseWriter, r *http.Request), httpAdder string){
	http.HandleFunc("/", getHttpHandler(routeMap))
	http.ListenAndServe(httpAdder, nil)
}


/*
获取一个空路由map
return map[string]func(w http.ResponseWriter, r *http.Request)
 */
func GetBlockUrlRouteMap()map[string]func(w http.ResponseWriter, r *http.Request){
	return make(map[string]func(w http.ResponseWriter, r *http.Request))
}
