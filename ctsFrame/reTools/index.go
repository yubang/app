package reTools

import "regexp"

/*
正则表达式相关封装
@author: yubang
创建于2017年4月23日
 */

/*
字符串替换
@param regStr: 正则表达式字符串
@param afterReplaceStr: 替换后字符串
@param sourceStr: 源字符串
@return string
 */
func Replace(regStr string, afterReplaceStr string, sourceStr string)string{
	reg := regexp.MustCompile(regStr)
	return reg.ReplaceAllString(sourceStr, afterReplaceStr)
}

/*
字符串匹配
@param regStr: 正则表达式字符串
@param sourceStr: 源字符串
@return bool
 */
func Match(regStr string, sourceStr string)bool{
	reg := regexp.MustCompile(regStr)
	return reg.Match([]byte(sourceStr))
}