package tools

import "encoding/json"

/*
json处理相关工具类
@author: yubang
 */

// json字符串转map对象
func JsonToInterface(jsonString []byte) map[string]interface{}{
	r := make(map[string]interface{})
	err := json.Unmarshal(jsonString, &r)
	if err != nil{
		return nil
	}
	return r
}

//map对象转json字符串
func InterfaceToJson(obj map[string]interface{})[]byte{
	r, err := json.Marshal(obj)
	if err != nil{
		return nil
	}
	return r
}