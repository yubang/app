package timeTools

import "time"

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
