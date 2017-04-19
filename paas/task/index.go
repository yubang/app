package task

/*
任务相关模块
@author: yubang
 */

import "../tools"


// 生成一个随机任务id
func BuildTaskId() string{
	return tools.GetNowTime("%Y%m%d_") + tools.GetToken()
}
