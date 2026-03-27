/*
@File    : result.go
@Author  : GuguLH
@Date    : 2026/3/27 10:23
@Desc    : 统一返回结构体
*/

package ginx

type R struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
