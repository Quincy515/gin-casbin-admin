package viper

import (
	"fmt"
	"github.com/custer-go/gin-casbin-admin/global"
	"github.com/custer-go/gin-casbin-admin/pkg/utils"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	v := viper.New()
	v.SetConfigFile(utils.ConfigFile)
	err := v.ReadInConfig() // 读取配置文件
	if err != nil {
		panic(fmt.Errorf("读取配置文件 config.yaml 失败：%s\n", err))
	}
	v.WatchConfig() // 监控配置文件变化

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件发生更改：", e.Name)
		if err := v.Unmarshal(&global.SYS_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&global.SYS_CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
