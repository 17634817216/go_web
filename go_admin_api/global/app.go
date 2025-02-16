package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go_admin_api/Config"
	"gopkg.in/ini.v1"
	"log"
)

type Application struct {
	ConfigViper    *viper.Viper
	DatabaseConfig Config.DatabaseConfig
	AppConfig      Config.AppConfig
	LogConfig      Config.LogConfig
	Log            *zap.Logger
}

var App = new(Application)

func InifConfig() {
	configFilePath := Config.OpenFile()
	// 加载 INI 文件
	cfg, err := ini.Load(configFilePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// 将 INI 文件中的配置映射到结构体
	err = cfg.Section("DATABASE").MapTo(&App.DatabaseConfig)
	if err != nil {
		log.Fatalf("Failed to map DATABASE section to struct: %v", err)
	}
	err = cfg.Section("APP").MapTo(&App.AppConfig)
	if err != nil {
		log.Fatalf("Failed to map APP section to struct: %v", err)
	}
	err = cfg.Section("LOG").MapTo(&App.LogConfig)
	if err != nil {
		log.Fatalf("Failed to map LOG section to struct: %v", err)
	}

}
