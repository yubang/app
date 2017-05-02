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

// 获取子串在字符串的字节位置
func UnicodeIndex(str,substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str,substr)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	}

	return result
}

// 字符串截取
func SubString(str string, start, length int) (substr string) {
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