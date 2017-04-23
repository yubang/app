package webTools

/*
静态资源处理
@author: yubang
创建于2017年4月23日
 */

import "net/http"
import "../fileTools"

/*
处理静态资源文件
@param filePath: 文件路径
@param cacheTimeoutSecond: 静态资源缓存时间
@param w: http.ResponseWriter
 */
func SendFile(filePath string, cacheTimeoutSecond int, w http.ResponseWriter){
	fileContent := fileTools.ReadFromFile(filePath)
	if fileContent == nil{
		// 资源找不到
		w.WriteHeader(404)
		return
	}
	w.Header().Set("Content-Type", GetContentTypeFromName(filePath))
	w.Write(fileContent)
}
