/*
@File    : student.go
@Author  : GuguLH
@Date    : 2026/3/26 11:28
@Desc    : Student 处理器
*/

package web

import (
	"errors"
	"net/http"

	"github.com/GuguLH/gin-spark/internal/domain"
	"github.com/GuguLH/gin-spark/internal/errs"
	"github.com/GuguLH/gin-spark/internal/service"
	"github.com/GuguLH/gin-spark/pkg/logger"
	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	svc service.IStudentService
	l   logger.Logger
}

func NewStudentHandler(
	svc service.IStudentService,
	l logger.Logger,
) *StudentHandler {
	return &StudentHandler{
		svc: svc,
		l:   l,
	}
}

func (h *StudentHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/stu")
	ug.POST("/create", h.CreateStudent)
	ug.GET("/:id", h.GetStudent)
	ug.POST("/edit", h.Edit)
	ug.POST("/remove", h.Remove)
}

func (h *StudentHandler) CreateStudent(ctx *gin.Context) {
	type CreateStudentReq struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}
	var req CreateStudentReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	err := h.svc.Create(ctx, domain.Student{
		Name:  req.Name,
		Phone: req.Phone,
	})
	if err != nil {
		if errors.Is(err, service.ErrDuplicatePhone) {
			ctx.JSON(http.StatusOK, R{
				Msg:  "手机号冲突",
				Code: errs.UserDuplicatePhone,
			})
			h.l.Error(
				"手机号冲突",
				logger.String("phone", req.Phone),
				logger.Error(err),
			)
			return
		}
		ctx.JSON(http.StatusOK, R{
			Msg:  "系统错误",
			Code: errs.UserInternalServerError,
		})
		return
	}
	ctx.JSON(http.StatusOK, R{Msg: "ok"})
}

func (h *StudentHandler) GetStudent(ctx *gin.Context) {

}

func (h *StudentHandler) Edit(ctx *gin.Context) {

}

func (h *StudentHandler) Remove(ctx *gin.Context) {

}
