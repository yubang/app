package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"../../paas/tools"
)

/*
http proxy相关业务封装
@author: yubang
 */

// 处理http proxy
func ProxyHttp(proxyHost string, proxyPort int, w http.ResponseWriter, r *http.Request){
	sourceUrl := "http://" + proxyHost + ":" + tools.IntToString(proxyPort) + r.RequestURI
	tools.Debug("反向代理到URL：" + sourceUrl)
	proxyUrl, _ := url.Parse(sourceUrl)
	httpProxy := httputil.NewSingleHostReverseProxy(proxyUrl)
	httpProxy.ServeHTTP(w, r)
}
