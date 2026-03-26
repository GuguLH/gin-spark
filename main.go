/*
@File    : main.go
@Author  : GuguLH
@Date    : 2026/3/26 10:06
@Desc    :
*/

package main

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	// 1 配置读取
	initViper()
	// 2 web服务获取
	server := InitWebServer()
	server.Run(":9090")
}

func initViper() {
	cFile := pflag.String("config", "config/dev.yaml", "配置文件路径")
	pflag.Parse()
	viper.SetConfigType("yaml")
	viper.SetConfigFile(*cFile)
	// 读取配置
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
