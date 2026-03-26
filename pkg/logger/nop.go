/*
@File    : nop.go
@Author  : GuguLH
@Date    : 2026/3/26 9:54
@Desc    : 无日志实现
*/

package logger

type NopLogger struct {
}

func NewNopLogger() *NopLogger {
	return &NopLogger{}
}

func (n *NopLogger) Debug(msg string, args ...Field) {
}

func (n *NopLogger) Info(msg string, args ...Field) {
}

func (n *NopLogger) Warn(msg string, args ...Field) {
}

func (n *NopLogger) Error(msg string, args ...Field) {
}
