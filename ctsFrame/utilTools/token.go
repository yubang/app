package utilTools

/*
唯一码相关操作
@author: yubang
创建于2017年4月23日
 */

import (
	"../hashTools"
	"time"
	"../typeConversionTools"
)

/*
获取一个随机32位字符串
@return string
 */
func GetToken32()string{
	t := time.Now()
	d := typeConversionTools.Int64ToString(t.UnixNano())
	return hashTools.Md5([]byte(d))
}
