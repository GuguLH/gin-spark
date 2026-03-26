/*
@File    : web.go
@Author  : GuguLH
@Date    : 2026/3/26 10:11
@Desc    : web服务
*/

package ioc

import (
	"strings"
	"time"

	"github.com/GuguLH/gin-spark/internal/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitWebServer(
	mdls []gin.HandlerFunc,
	studentHdl *web.StudentHandler,
) *gin.Engine {
	server := gin.Default()
	server.Use(mdls...)
	// Handler注册路由
	studentHdl.RegisterRoutes(server)
	return server
}

func InitGinMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		cors.New(cors.Config{
			AllowCredentials: true,
			AllowHeaders:     []string{"Content-Type"},
			//这个是允许前端访问你的后端响应中带的头部
			ExposeHeaders: []string{"x-jwt-token", "x-refresh-token"},
			AllowOriginFunc: func(origin string) bool {
				if strings.HasPrefix(origin, "http://localhost") {
					return true
				}
				return strings.Contains(origin, "your_company.com")
			},
			MaxAge: 12 * time.Hour,
		}),
	}
}
