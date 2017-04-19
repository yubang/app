package tools

import (
	"crypto/md5"
	"time"
	"encoding/hex"
)

/*
hash模块工具类
@author: yubang
 */


// md5计算
func Md5(s string)string{
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// 获取一个32位随机码
func GetToken() string{
	t := time.Now()
	return Md5(Int64ToString(t.UnixNano()))
}