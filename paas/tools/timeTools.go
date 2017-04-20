package tools

import (
	"time"
	"strings"
)

/*
时间工具类
@author: yubang
 */

// 获取系统当前时间
func GetNowTime(timeFormat string) string{
	// 2006-01-02 15:04:05 ======> %Y-%m-%d %H:%M:%S
	var arrs = [6][2]string{{"%Y", "2006"},{"%m", "01"},{"%d", "02"},{"%H", "15"},{"%M", "04"},{"%S", "05"}}
	for _, obj := range arrs{
		timeFormat = strings.Replace(timeFormat, obj[0], obj[1], -1)
	}

	return time.Now().Format(timeFormat)
}

// 获取系统当前时间
func GetNowTimeSecond()int{
	return Int64ToInt(time.Now().UnixNano() / 1000 /1000)
}