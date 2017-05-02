package typeConversionTools

import (
	"strconv"
	"errors"
	"../reTools"
)

/*
封装类型转换相关操作
@author: yubang
创建于2017年04月22日
 */

/*
[]byte -> string
@param b: 待转换参数
@return string
 */
func ByteListToString(b []byte)string{
	return string(b)
}

/*
float64-> string
@param n: 待转换参数
@return string
 */
func Float64ToString(n float64) string{
	return strconv.FormatFloat(n,'f', -1, 64)
}

/*
int -> string
@param n: 待转换参数
@return string
 */
func IntToString(n int) string{
	return strconv.Itoa(n)
}

/*
int64 -> string
@param n: 待转换参数
@return string
 */
func Int64ToString(n int64) string{
	return strconv.FormatInt(n,10)
}

/*
string -> int
@param s: 待转换参数
@return int, error
 */
func StringToInt(s string) (int, error){
	n, err := strconv.Atoi(s)
	if err != nil{
		return 0, errors.New("参数类型不正确！")
	}
	return n, nil
}

/*
string -> int64
@param s: 待转换参数
@return int64, error
 */
func StringToInt64(s string) (int64, error){
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil{
		return 0, errors.New("参数类型不正确！")
	}
	return n, nil
}

/*
string -> float64
@param s: 待转换参数
@return float64, error
 */
func StringToFloat64(s string) (float64, error){
	n, err := strconv.ParseFloat(s, 32)
	if err != nil{
		return 0.0, errors.New("参数类型不正确！")
	}
	return n, nil
}

/*
float64 -> int
@param n: 待转换参数
@return int
 */
func Float64ToInt(n float64) int{
	return int(n)
}

/*
float64 -> int64
@param n: 待转换参数
@return int
 */
func Float64ToInt64(n float64) int64{
	return int64(n)
}

/*
int64 -> int
@param n: 待转换参数
@return int
 */
func Int64ToInt(n int64)int{
	return int(n)
}


func IsNumber(s string)bool{
	return reTools.Match("\\d+[.]?\\d*", s)
}