package tools

import (
	"net/http"
	"io/ioutil"
)

func Post(url string, postData map[string][]string)map[string]interface{}{
	r, err := http.PostForm(url, postData)
	if err != nil || r.StatusCode != 200{
		return nil
	}
	result, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil{
		return nil
	}
	obj := JsonToInterface(result)
	return obj
}
