package tools

import (
	"strconv"
	"strings"
)

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

func Int64ToString(n int64) string{
	return strconv.FormatInt(n,10)
}

func StringToInt(s string) int{
	n, err := strconv.Atoi(s)
	if err != nil{
		panic("参数类型不正确！")
		return 0
	}
	return n
}

func Float64ToInt(n float64) int{
	return int(n)
}

// 字符串截取
func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

// 获取分割后最后的字符串
func GetSplitLastArr(s string, splitString string) string{
	arrs := strings.Split(s, splitString)
	return arrs[len(arrs) - 1]
}

// 获取分割后第一位的字符串
func GetSplitFirstArr(s string, splitString string) string{
	arrs := strings.Split(s, splitString)
	return arrs[0]
}