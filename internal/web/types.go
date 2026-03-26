/*
@File    : types.go
@Author  : GuguLH
@Date    : 2026/3/26 11:33
@Desc    :
*/

package web

import "github.com/gin-gonic/gin"

type Handler interface {
	RegisterRoutes(server *gin.Engine)
}

type Page struct {
	Limit  int
	Offset int
}

type R struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
