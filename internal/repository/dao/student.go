/*
@File    : student.go
@Author  : GuguLH
@Date    : 2026/3/26 10:30
@Desc    :
*/

package dao

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

var (
	ErrDuplicatePhone = errors.New("手机号冲突")
	ErrRecordNotFound = errors.New("未找到相关记录")
)

type IStudentDAO interface {
	Insert(ctx context.Context, s Student) error
	FindById(ctx context.Context, sId int64) (Student, error)
	UpdateById(ctx context.Context, s Student) error
	DeleteById(ctx context.Context, sId int64) error
}

type GormStudentDAO struct {
	db *gorm.DB
}

func NewGormStudentDAO(db *gorm.DB) IStudentDAO {
	return &GormStudentDAO{
		db: db,
	}
}

func (dao *GormStudentDAO) Insert(ctx context.Context, s Student) error {
	now := time.Now().UnixMilli()
	s.CreatedAt = now
	s.UpdatedAt = now
	err := dao.db.WithContext(ctx).Create(&s).Error
	if err != nil && strings.Contains(err.Error(), "UNIQUE constraint failed") {
		return ErrDuplicatePhone
	}
	return err
}

func (dao *GormStudentDAO) FindById(ctx context.Context, sId int64) (Student, error) {
	var ret Student
	err := dao.db.WithContext(ctx).Where("id = ?", sId).First(&ret).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ret, ErrRecordNotFound
	}
	return ret, err
}

func (dao *GormStudentDAO) UpdateById(ctx context.Context, s Student) error {
	ret := dao.db.WithContext(ctx).Model(&s).Where("id = ?", s.Id).
		Updates(map[string]any{
			"name":      s.Name,
			"phone":     s.Phone,
			"update_at": time.Now().UnixMilli(),
		})
	if ret.Error != nil {
		return ret.Error
	}
	if ret.RowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

func (dao *GormStudentDAO) DeleteById(ctx context.Context, sId int64) error {
	ret := dao.db.WithContext(ctx).Delete(&Student{}, sId)
	if ret.Error != nil {
		return ret.Error
	}
	if ret.RowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

type Student struct {
	Id        int64          `gorm:"primaryKey,autoIncrement"`
	Name      string         `gorm:"type=varchar(64)"`
	Phone     sql.NullString `gorm:"unique"`
	CreatedAt int64
	UpdatedAt int64
}
