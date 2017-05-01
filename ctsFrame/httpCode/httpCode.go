package httpCode

/*
内置的http状态码
@author: yubang
创建于2017年4月25日
 */

type HttpCode struct {
	Code int
	Msg string
}

var OkCode = HttpCode{0, "success"}
var TokenErrorCode = HttpCode{10001, "Token错误"}
var NotTaskCode = HttpCode{10002, "任务队列为空"}
var ParameterMissingCode = HttpCode{10003, "参数缺失"}
var ServerErrorCode = HttpCode{10004, "服务器内部错误"}
var NeedLoginCode = HttpCode{10005, "需要登录"}
var StopCode = HttpCode{10006, "越权操作"}
var PowerErrorCode = HttpCode{10007, "权限错误"}
