package config

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
	"path/filepath"
)

// 定义数据库配置结构体
type DatabaseConfig struct {
	DB_DRVICE   string `ini:"DB_DRVICE"`
	DB_NAME     string `ini:"DB_NAME"`
	DB_USER     string `ini:"DB_USER"`
	DB_PASSWORD string `ini:"DB_PASSWORD"`
	DB_HOST     string `ini:"DB_HOST"`
	DB_PORT     int    `ini:"DB_PORT"`
}

// 定义应用程序配置结构体
type AppConfig struct {
	VERSION    string `ini:"VERSION"`
	PORT       string `ini:"PORT"`
	CPUNUM     string `ini:"CPUNUM"`
	RUNLOGTYPE string `ini:"RUNLOGTYPE"`
	ENV        string `ini:"ENV"`
}

type LogConfig struct {
	LEVEL       string `ini:"LEVEL"`
	ROOT_DIR    string `ini:"ROOT_DIR"`
	FILENAME    string `ini:"FILENAME"`
	FORMAT      string `ini:"FORMAT"`
	SHOW_LINE   bool   `ini:"SHOW_LINE"`
	MAX_BACKUPS int    `ini:"MAX_BACKUPS"`
	MAX_SIZE    int    `ini:"MAX_SIZE"`
	MAX_AGE     int    `ini:"MAX_AGE"`
	COMPRESS    bool   `ini:"COMPRESS"`
}

type CacheConfig struct {
	CACHE_TYPE     string `ini:"CACHE_TYPE"`
	CACHE_HOST     string `ini:"CACHE_HOST"`
	CACHE_PASSWPRD string `ini:"CACHE_PASSWPRD"`
	CACHE_DB       int    `ini:"CACHE_DB"`
}

func Inconfig(ConfigPath string) {
	cfg := ini.Empty()

	// 添加 DATABASE 部分
	databaseSection, err := cfg.NewSection("DATABASE")
	if err != nil {
		log.Fatalf("无法创建 DATABASE 部分: %v", err)
	}
	databaseSection.Key("DB_DRVICE").SetValue("mysql")
	databaseSection.Key("DB_NAME").SetValue("mems_db")
	databaseSection.Key("DB_USER").SetValue("root")
	databaseSection.Key("DB_PASSWORD").SetValue("qq080808")
	databaseSection.Key("DB_HOST").SetValue("localhost")
	databaseSection.Key("DB_PORT").SetValue("3306")

	// 添加 APP 部分
	appSection, err := cfg.NewSection("APP")
	if err != nil {
		log.Fatalf("无法创建 APP 部分: %v", err)
	}
	appSection.Key("VERSION").SetValue("1.3.0")
	appSection.Key("PORT").SetValue("8000")
	appSection.Key("CPUNUM").SetValue("3")
	appSection.Key("RUNLOGTYPE").SetValue("debug")
	appSection.Key("ENV").SetValue("dev")

	// 添加 CACHE 缓存部分
	CacheSection, err := cfg.NewSection("CACHE")
	if err != nil {
		log.Fatalf("无法创建 CACHE 部分: %v", err)
	}
	CacheSection.Key("CACHE_TYPE").SetValue("bigcache")
	CacheSection.Key("CACHE_HOST").SetValue("localhost:6379")
	CacheSection.Key("CACHE_PASSWPRD").SetValue("qq080808")
	CacheSection.Key("CACHE_DB").SetValue("0")

	logSection, err := cfg.NewSection("LOG")
	if err != nil {
		log.Fatalf("无法创建 LOG 部分: %v", err)
	}

	logSection.Key("LEVEL").SetValue("info")
	logSection.Key("ROOT_DIR").SetValue("./runtime/logs")
	logSection.Key("FILENAME").SetValue("app.log")
	logSection.Key("SHOW_LINE").SetValue("json")
	logSection.Key("MAX_BACKUPS").SetValue("3")
	logSection.Key("MAX_SIZE").SetValue("500")
	logSection.Key("MAX_AGE").SetValue("28")
	logSection.Key("COMPRESS").SetValue("true")
	// 将配置写入到 config.ini 文件
	err = cfg.SaveTo(ConfigPath)
	if err != nil {
		log.Fatalf("无法保存配置到文件: %v", err)
	}

	log.Println("配置已成功写入到 config/config.ini")
}

func OpenFile() string {
	//StatmPath := utils.GetStatmPath()
	cwd, err := os.Getwd()
	if err != nil {
		log.Println("获取当前工作目录失败:", err)
		return ""
	}

	// 构建配置文件的完整路径
	configFilePath := filepath.Join(cwd, "config", "config.ini")

	// 检查配置文件是否存在
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		// 如果文件不存在，调用 Inconfig 创建它
		Inconfig(configFilePath)
	} else if err != nil {
		// 如果其他错误，打印错误信息
		log.Println("检查配置文件失败:", err)

		return ""
	}
	return configFilePath

}
