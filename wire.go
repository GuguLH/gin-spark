//go:build wireinject

/*
@File    : wire.go
@Author  : GuguLH
@Date    : 2026/3/26 10:10
@Desc    : wire依赖注入
*/
package main

import (
	"github.com/GuguLH/gin-spark/internal/repository"
	"github.com/GuguLH/gin-spark/internal/repository/dao"
	"github.com/GuguLH/gin-spark/internal/service"
	"github.com/GuguLH/gin-spark/internal/web"
	"github.com/GuguLH/gin-spark/ioc"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// 第三方依赖
var thirdPartySet = wire.NewSet(
	ioc.InitDB,
	ioc.InitLogger,
)

// Web服务
var webSet = wire.NewSet(
	ioc.InitGinMiddlewares,
	ioc.InitWebServer,
)

// Student服务
var studentSet = wire.NewSet(
	web.NewStudentHandler,
	service.NewStudentService,
	repository.NewStudentRepository,
	dao.NewGormStudentDAO,
)

func InitWebServer() *gin.Engine {
	wire.Build(
		// 第三方依赖
		thirdPartySet,

		// Web服务
		webSet,

		// Student服务
		studentSet,
	)
	return gin.Default()
}
