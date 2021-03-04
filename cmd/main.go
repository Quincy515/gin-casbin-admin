package main

import (
	"github.com/custer-go/gin-casbin-admin/global"
	"github.com/custer-go/gin-casbin-admin/global/gorm"
	"github.com/custer-go/gin-casbin-admin/global/logger"
	"github.com/custer-go/gin-casbin-admin/global/viper"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.SYS_VIPER = viper.Viper() // 初始化 viper
	global.SYS_LOG = logger.Zap()    // 初始化 zap 日志库
	global.SYS_DB = gorm.Gorm()      // gorm 连接数据库
}
