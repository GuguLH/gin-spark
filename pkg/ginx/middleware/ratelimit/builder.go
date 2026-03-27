/*
@File    : builder.go
@Author  : GuguLH
@Date    : 2026/3/27 10:37
@Desc    : IP限流器构建器
*/

package ratelimit

import (
	"context"
	"fmt"
	"net/http"

	"github.com/GuguLH/gin-spark/pkg/limiter"
	"github.com/GuguLH/gin-spark/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Builder struct {
	prefix  string
	limiter limiter.Limiter
	l       logger.Logger
}

func NewBuilder(limiter limiter.Limiter, l logger.Logger) *Builder {
	return &Builder{
		prefix:  "ip-limiter",
		limiter: limiter,
		l:       l,
	}
}

func (b *Builder) Prefix(prefix string) *Builder {
	b.prefix = prefix
	return b
}

func (b *Builder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if ctx.GetHeader("x-stress") == "true" {
			// 用 context.Context 来带这个标记位
			newCtx := context.WithValue(ctx, "x-stress", true)
			ctx.Request = ctx.Request.Clone(newCtx)
			ctx.Next()
			return
		}

		limited, err := b.limiter.Limit(ctx, fmt.Sprintf("%s:%s", b.prefix, ctx.ClientIP()))
		if err != nil {
			b.l.Error("限流器错误:", logger.Error(err))
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if limited {
			b.l.Error("触发了限流:", logger.String("IP", ctx.ClientIP()), logger.Error(err))
			ctx.AbortWithStatus(http.StatusTooManyRequests)
			return
		}
		ctx.Next()
	}
}
