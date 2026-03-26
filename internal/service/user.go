/*
@File    : user.go
@Author  : GuguLH
@Date    : 2026/3/26 11:16
@Desc    :
*/

package service

import (
	"context"

	"github.com/GuguLH/gin-spark/internal/domain"
	"github.com/GuguLH/gin-spark/internal/repository"
)

var (
	ErrDuplicatePhone = repository.ErrDuplicatePhone
	ErrRecordNotFound = repository.ErrRecordNotFound
)

type IStudentService interface {
	Create(ctx context.Context, s domain.Student) error
	FindById(ctx context.Context, sId int64) (domain.Student, error)
	Update(ctx context.Context, s domain.Student) error
	Delete(ctx context.Context, sId int64) error
}

type StudentService struct {
	repo repository.IStudentRepository
}

func NewStudentService(repo repository.IStudentRepository) IStudentService {
	return &StudentService{
		repo: repo,
	}
}
func (svc *StudentService) Create(ctx context.Context, s domain.Student) error {
	return svc.repo.Create(ctx, s)
}

func (svc *StudentService) FindById(ctx context.Context, sId int64) (domain.Student, error) {
	return svc.repo.FindById(ctx, sId)
}

func (svc *StudentService) Update(ctx context.Context, s domain.Student) error {
	return svc.repo.UpdateById(ctx, s)
}

func (svc *StudentService) Delete(ctx context.Context, sId int64) error {
	return svc.repo.DeleteById(ctx, sId)
}
