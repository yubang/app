package stringTools

import "strings"

/*
字符串相关操作封装
@author: yubang
创建于2017年04月22日
 */

/*
获取分割后最后的字符串
@param s: 源字符串
@param splitString: 切割分隔符
@return 字符串
 */
func GetSplitLastArr(s string, splitString string) string{
	arrs := strings.Split(s, splitString)
	return arrs[len(arrs) - 1]
}

/*
获取分割后第一位的字符串
@param s: 源字符串
@param splitString: 切割分隔符
@return 字符串
 */
func GetSplitFirstArr(s string, splitString string) string{
	arrs := strings.Split(s, splitString)
	return arrs[0]
}