/*
@File    : student.go
@Author  : GuguLH
@Date    : 2026/3/26 10:59
@Desc    :
*/

package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/GuguLH/gin-spark/internal/domain"
	"github.com/GuguLH/gin-spark/internal/repository/dao"
)

var (
	ErrDuplicatePhone = dao.ErrDuplicatePhone
	ErrRecordNotFound = dao.ErrRecordNotFound
)

type IStudentRepository interface {
	Create(ctx context.Context, s domain.Student) error
	FindById(ctx context.Context, sId int64) (domain.Student, error)
	UpdateById(ctx context.Context, s domain.Student) error
	DeleteById(ctx context.Context, sId int64) error
}

type StudentRepository struct {
	dao dao.IStudentDAO
}

func NewStudentRepository(dao dao.IStudentDAO) IStudentRepository {
	return &StudentRepository{
		dao: dao,
	}
}

func (repo *StudentRepository) Create(ctx context.Context, s domain.Student) error {
	return repo.dao.Insert(ctx, repo.toEntity(s))
}

func (repo *StudentRepository) FindById(ctx context.Context, sId int64) (domain.Student, error) {
	s, err := repo.dao.FindById(ctx, sId)
	if err != nil {
		return domain.Student{}, err
	}
	return repo.toDomain(s), nil
}

func (repo *StudentRepository) UpdateById(ctx context.Context, s domain.Student) error {
	return repo.dao.UpdateById(ctx, repo.toEntity(s))
}

func (repo *StudentRepository) DeleteById(ctx context.Context, sId int64) error {
	return repo.dao.DeleteById(ctx, sId)
}

func (repo *StudentRepository) toDomain(s dao.Student) domain.Student {
	return domain.Student{
		Id:        s.Id,
		Name:      s.Name,
		Phone:     s.Phone.String,
		UpdateAt:  time.UnixMilli(s.UpdatedAt),
		CreatedAt: time.UnixMilli(s.CreatedAt),
	}
}

func (repo *StudentRepository) toEntity(s domain.Student) dao.Student {
	return dao.Student{
		Id:   s.Id,
		Name: s.Name,
		Phone: sql.NullString{
			String: s.Phone,
			Valid:  s.Phone != "",
		},
		UpdatedAt: s.UpdateAt.UnixMilli(),
		CreatedAt: s.CreatedAt.UnixMilli(),
	}
}
