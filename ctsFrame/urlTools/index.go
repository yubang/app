package urlTools

/*
url服务相关操作
@author: yubang
创建于2017年4月25日
 */

import "net/url"

/*
url encode
@param urlStr: 网址
@return string
 */
func UrlEncode(urlStr string)string{
	return url.QueryEscape(urlStr)
}