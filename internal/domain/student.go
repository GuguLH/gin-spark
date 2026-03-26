/*
@File    : student.go
@Author  : GuguLH
@Date    : 2026/3/26 10:58
@Desc    : Student领域对象
*/

package domain

import "time"

type Student struct {
	Id        int64
	Name      string
	Phone     string
	CreatedAt time.Time
	UpdateAt  time.Time
}
