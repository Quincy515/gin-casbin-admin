package gorm

import (
	"fmt"
	"github.com/custer-go/gin-casbin-admin/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// Gorm 初始化数据库连接，并产生数据库全局变量
func Gorm() *gorm.DB {
	switch global.SYS_CONFIG.System.DbType {
	case "mysql":
		return GormMySQL()
	default:
		return GormMySQL()
	}
}

// GormMySQL 初始化 MySQL 数据库
func GormMySQL() *gorm.DB {
	m := global.SYS_CONFIG.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", m.Username, m.Password, m.Path, m.Dbname, m.Config)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(m.LogMode)); err != nil {
		global.SYS_LOG.Error("MySQL启动失败", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// gormConfig 根据配置决定是否开启日志
func gormConfig(mod bool) *gorm.Config {
	var config = &gorm.Config{}
	switch global.SYS_CONFIG.Mysql.LogZap {
	case "silent", "Silent":
		config.Logger = NewGormLogger(logger.Silent)
	case "error", "Error":
		config.Logger = NewGormLogger(logger.Error)
	case "warn", "Warn":
		config.Logger = NewGormLogger(logger.Warn)
	case "info", "Info":
		config.Logger = NewGormLogger(logger.Info)
	case "zap", "Zap":
		config.Logger = NewGormLogger(logger.Info)
	default:
		if mod {
			config.Logger = NewGormLogger(logger.Info)
			break
		}
		config.Logger = NewGormLogger(logger.Silent)
	}
	return config
}

// NewGormLogger 默认 gorm logger new 一个实例
func NewGormLogger(level logger.LogLevel) logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      level,       // Log level
			Colorful:      true,        // 黑白打印
		},
	)
}
