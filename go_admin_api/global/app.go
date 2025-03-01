package global

import (
	"context"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go_admin_api/config"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
	"log"
	"time"
)

type Application struct {
	ConfigViper    *viper.Viper
	DatabaseConfig config.DatabaseConfig
	AppConfig      config.AppConfig
	LogConfig      config.LogConfig
	CacheConfig    config.CacheConfig
	Log            *zap.Logger
	DB             *gorm.DB
	Cache          Cache
}

var App = new(Application)

// InifConfig 初始化配置
func InifConfig() {
	configFilePath := config.OpenFile()
	if configFilePath == "" {
		log.Fatalf("Config file path is empty")
	}

	// 加载 INI 文件
	cfg, err := ini.Load(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config file: %v", err)
	}

	// 映射配置到结构体
	mapConfigSection(cfg, "DATABASE", &App.DatabaseConfig)
	mapConfigSection(cfg, "APP", &App.AppConfig)
	mapConfigSection(cfg, "LOG", &App.LogConfig)
	mapConfigSection(cfg, "CACHE", &App.CacheConfig)

}

// mapConfigSection 将 INI 文件的指定 section 映射到结构体
func mapConfigSection(cfg *ini.File, sectionName string, configStruct interface{}) {
	section, err := cfg.GetSection(sectionName)
	if err != nil {
		log.Fatalf("Failed to get section %s: %v", sectionName, err)
	}

	if err := section.MapTo(configStruct); err != nil {
		log.Fatalf("Failed to map section %s to struct: %v", sectionName, err)
	}
}

type Cache interface {
	// Get 获取缓存值
	Get(ctx context.Context, key string) ([]byte, error)
	// SetEX 设置超时时间缓存值
	Set(ctx context.Context, key string, value []byte, ttl time.Duration) error
	// SetEX 设置缓存值
	//Set(ctx context.Context, key string, value []byte) error
	// Delete 删除缓存值
	Delete(ctx context.Context, key string) error
	// Close 关闭缓存连接
	Close() error
}
