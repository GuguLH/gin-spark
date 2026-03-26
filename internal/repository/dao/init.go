/*
@File    : init.go
@Author  : GuguLH
@Date    : 2026/3/26 10:04
@Desc    : 数据库迁移
*/

package dao

import "gorm.io/gorm"

func InitTables(db *gorm.DB) error {
	return db.AutoMigrate()
}
