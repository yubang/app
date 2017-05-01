package webTools

import "net/http"
import (
	"time"
	"../utilTools"
	"../cacheTools"
)

/*
session模块
@author: yubang
创建于2017年4月23日
 */

func getSession(r *http.Request, cacheClient cacheTools.CacheClient)map[string]interface{}{
	cookie, err := r.Cookie("sessionId")
	if err != nil{
		return make(map[string]interface{})
	}
	d := cacheClient.Get(cookie.Value)
	if d == nil{
		return make(map[string]interface{})
	}
	return d.(map[string]interface{})
}

func setSession(request *HttpObject, sessionMap map[string]interface{}){
	w := request.Response
	r := request.Request
	cookie, err := r.Cookie("sessionId")

	COOKIE_MAX_MAX_AGE := time.Hour * 24 / time.Second   // 单位：秒。
	maxAge := int(COOKIE_MAX_MAX_AGE)

	if err != nil{
		cookie = &http.Cookie{
			Name:   "sessionId",
			Value:    utilTools.GetToken32(),
			Path:     "/",
			HttpOnly: false,
			MaxAge:   maxAge,
		}

	}

	// 记录数据
	request.CacheClient.Set(cookie.Value, sessionMap)
	request.CacheClient.Expir(cookie.Value, 30 * 60)

	http.SetCookie(w, cookie)
}