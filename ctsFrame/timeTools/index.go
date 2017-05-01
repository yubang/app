package timeTools

import (
	"time"
	"strings"
)

/*
时间相关模块
@author: yubang
创建于2017年4月23日
 */

/*
获取当前时间秒数
 */
func GetNowTimeSecond()int64{
	return time.Now().Unix()
}

// 获取系统当前时间
func GetNowTime(timeFormat string) string{
	// 2006-01-02 15:04:05 ======> %Y-%m-%d %H:%M:%S
	var arrs = [6][2]string{{"%Y", "2006"},{"%m", "01"},{"%d", "02"},{"%H", "15"},{"%M", "04"},{"%S", "05"}}
	for _, obj := range arrs{
		timeFormat = strings.Replace(timeFormat, obj[0], obj[1], -1)
	}

	return time.Now().Format(timeFormat)
}