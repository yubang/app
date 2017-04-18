package tools

import "fmt"

/*
日志模块
@author: yubang
*/

// 输出日志
func loginLog(level string, msg string){
	fmt.Print(GetNowTime("%Y-%m-%d %H:%M:%S") + " " + level + " " + msg + "\n")
}

// 输出错误日志
func Error(msg string){
	loginLog("ERROR", msg)
}

// 输出普通日志
func Info(msg string){
	loginLog("INFO", msg)
}

// 输出debug日志
func Debug(msg string){
	loginLog("DEBUG", msg)
}