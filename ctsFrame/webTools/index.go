package webTools

import (
	"../stringTools"
)

/*
封装http服务相关操作
@author: yubang
 */

/*
获取文件名相对应content-type
@param name: 文件名
创建于2017年04月22日
 */
func GetContentTypeFromName(name string)string{
	suffixName := stringTools.GetSplitLastArr(name, ".")
	contentType := mimeTypesMap[suffixName]
	if contentType == ""{
		return "application/octet-stream"
	}
	return contentType
}
