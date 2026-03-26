/*
@File    : db.go
@Author  : GuguLH
@Date    : 2026/3/26 9:59
@Desc    : 数据库连接
*/

package ioc

import (
	"github.com/GuguLH/gin-spark/internal/repository/dao"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	type Config struct {
		DSN string `yaml:"dsn"`
	}
	var cfg Config
	err := viper.UnmarshalKey("db", &cfg)
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 数据库迁移
	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}
