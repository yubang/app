package httpCode

type HttpCode struct {
	Code int
	Msg string
}

var OkCode = HttpCode{0, "success"}
var TokenErrorCode = HttpCode{10001, "Token错误"}
var NotTaskCode = HttpCode{10002, "任务队列为空"}
var ParameterMissingCode = HttpCode{10003, "参数缺失"}
var ServerErrorCode = HttpCode{10004, "服务器内部错误"}