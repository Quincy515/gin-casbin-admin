package global

import (
	"github.com/custer-go/gin-casbin-admin/global/model"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	SYS_CONFIG model.Server
	SYS_VIPER  *viper.Viper
	SYS_LOG    *zap.Logger
	SYS_DB     *gorm.DB
	SYS_REDIS  *redis.Client
)
