/*
@File    : types.go
@Author  : GuguLH
@Date    : 2026/3/27 10:34
@Desc    : 限流器抽象
*/

package limiter

import "context"

type Limiter interface {
	// Limit 是否触发限流
	// 返回 true，就是触发限流
	Limit(ctx context.Context, key string) (bool, error)
}
