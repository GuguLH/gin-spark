/*
@File    : logger.go
@Author  : GuguLH
@Date    : 2026/3/26 10:18
@Desc    : 日志服务
*/

package ioc

import (
	"github.com/GuguLH/gin-spark/pkg/logger"
	"go.uber.org/zap"
)

func InitLogger() logger.Logger {
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return logger.NewZapLogger(l)
}
