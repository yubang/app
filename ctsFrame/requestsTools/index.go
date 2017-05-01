package requestsTools

/*
get,post模块
@author: yubang
创建于2017年4月25日
 */

import (
	"../jsonTools"
	"io/ioutil"
	"net/url"
	"net/http"
)

type HttpResponse struct {
	StatusCode int
	Content []byte
}

func (obj *HttpResponse)Json()map[string]interface{}{
	return jsonTools.JsonToInterface(obj.Content)
}

func (obj *HttpResponse)String()string{
	return string(obj.Content)
}

func Get(urlString string, getMap map[string]string)HttpResponse{
	httpResponse := HttpResponse{-1, []byte("")}
	u, _ := url.Parse(urlString)
	q := u.Query()
	for _, v := range getMap{
		q.Set(v, getMap[v])
	}
	u.RawQuery = q.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		return httpResponse
	}
	result, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil{
		return httpResponse
	}
	httpResponse.StatusCode = res.StatusCode
	httpResponse.Content = result
	return httpResponse
}

func Post(urlString string, postMap map[string]interface{})HttpResponse{
	return HttpResponse{}
}
