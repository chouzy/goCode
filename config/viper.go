package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitViper(globalConf *GlobalConfig, path string) {
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %v \n", err))
	}
	// 解析本地配置到全局变量
	if err = viper.Unmarshal(&globalConf); err != nil {
		panic(fmt.Errorf("Unmarshal conf failed, err:%s \n", err))
	}
	// 监控配置文件变化
	viper.WatchConfig()
	// 配置文件发生变化后同步到全局
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件已被修改...")
		if err = viper.Unmarshal(&globalConf); err != nil {
			panic(fmt.Errorf("Unmarshal conf failed, err:%s \n", err))
		}
	})
}
