package hashTools

import (
	"crypto/md5"
	"encoding/hex"
)

/*
md5模块
@author: yubang
创建于2017年4月23日
 */

/*
md5计算
@param s: 需要计算的[]byte
@return md5字符串
 */
func Md5(s []byte)string{
	md5Ctx := md5.New()
	md5Ctx.Write(s)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}