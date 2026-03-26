/*
@File    : student.go
@Author  : GuguLH
@Date    : 2026/3/26 9:52
@Desc    : Logger 抽象
*/

package logger

type Field struct {
	Key string
	Val any
}

type Logger interface {
	Debug(msg string, args ...Field)
	Info(msg string, args ...Field)
	Warn(msg string, args ...Field)
	Error(msg string, args ...Field)
}
