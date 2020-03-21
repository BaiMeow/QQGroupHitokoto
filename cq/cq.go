package cq

import (
	"fmt"

	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
)

//Info 添加酷QInfo
func Info(a ...interface{}) {
	cqp.AddLog(cqp.Info, "消息", fmt.Sprintf("%v", a))
}

//Error 添加酷qError
func Error(a ...interface{}) {
	cqp.AddLog(cqp.Error, "错误", fmt.Sprintf("%v", a))
}
