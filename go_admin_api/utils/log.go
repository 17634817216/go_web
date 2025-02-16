package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go_admin_api/global"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	level   zapcore.Level // zap 日志等级
	options []zap.Option  // zap 配置项
)

type zapWriter struct {
	logger *zap.Logger
}

func (w *zapWriter) Write(p []byte) (n int, err error) {
	w.logger.Info(string(p)) // 将 log 内容通过 zap 输出
	return len(p), nil
}

func InitializeLog() *zap.Logger {
	// 创建根目录
	createRootDir()

	// 设置日志等级
	setLogLevel()
	if global.App.LogConfig.SHOW_LINE {
		options = append(options, zap.AddCaller())
	}
	logger := zap.New(getZapCore(), options...)
	log.SetOutput(&zapWriter{logger: logger})
	// 初始化 zap
	return logger
}

func getZapCore() zapcore.Core {
	var encoder zapcore.Encoder

	// 调整编码器默认配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("[" + "2006-01-02 15:04:05.000" + "]"))
	}
	encoderConfig.EncodeLevel = func(l zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(global.App.AppConfig.ENV + "." + l.String())
	}

	// 设置编码器
	if global.App.LogConfig.FORMAT == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	return zapcore.NewCore(encoder, getLogWriter(), level)
}

func getLogWriter() zapcore.WriteSyncer {
	tiemstr := time.Now().Format("200601/02")
	file := &lumberjack.Logger{
		Filename:   global.App.LogConfig.ROOT_DIR + "/" + tiemstr + ".log",
		MaxSize:    global.App.LogConfig.MAX_SIZE,
		MaxBackups: global.App.LogConfig.MAX_BACKUPS,
		MaxAge:     global.App.LogConfig.MAX_AGE,
		Compress:   global.App.LogConfig.COMPRESS,
	}
	fileWriter := zapcore.AddSync(file)

	// 写入终端
	consoleWriter := zapcore.AddSync(os.Stdout)

	return zapcore.NewMultiWriteSyncer(fileWriter, consoleWriter)
}

func createRootDir() {
	StatmPath := GetStatmPath()
	var filename = filepath.Join(StatmPath, global.App.LogConfig.ROOT_DIR)
	if ok, _ := PathExists(filename); !ok {
		_ = os.Mkdir(filename, os.ModePerm)
	}
}

func setLogLevel() {
	switch global.App.LogConfig.LEVEL {
	case "debug":
		level = zap.DebugLevel
		options = append(options, zap.AddStacktrace(level))
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
		options = append(options, zap.AddStacktrace(level))
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
}
