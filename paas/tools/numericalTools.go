package tools

import "strconv"

/*
类型转换工具类
@author: yubang
 */

func Float64ToString(n float64) string{
	return strconv.FormatFloat(n,'f', -1, 64)
}

func IntToString(n int) string{
	return strconv.Itoa(n)
}
