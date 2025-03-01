package middleware

import (
	"fmt"
	"go_admin_api/global"
	"go_admin_api/internal/database"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func CustomRecovery() gin.HandlerFunc {
	//加载配置
	conf := global.App.LogConfig
	tiemstr := time.Now().Format("200601/02")
	return gin.RecoveryWithWriter(
		&lumberjack.Logger{
			Filename:   conf.ROOT_DIR + "/" + tiemstr + "_err.log",
			MaxSize:    conf.MAX_SIZE,
			MaxBackups: conf.MAX_BACKUPS,
			MaxAge:     conf.MAX_AGE,
			Compress:   conf.COMPRESS,
		},
		ServerError)
}

type Response struct {
	Code      int         `json:"code"`
	ErrorCode int         `json:"error_code"`
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
}

func ServerError(c *gin.Context, err interface{}) {
	conf := global.App.AppConfig
	msg := "内部服务器错误"
	if os.Getenv(gin.EnvGinMode) != gin.ReleaseMode && reflect.TypeOf(err).Name() == "string" {
		msg = err.(string)
	} else {
		if conf.ENV != "pro" && os.Getenv(gin.EnvGinMode) != gin.ReleaseMode {
			if _, ok := err.(error); ok {
				msg = err.(error).Error()
			}
		} else {
			str := fmt.Sprintf("内部服务器错误： %s\n", err.(error).Error()) //拼接字符串
			global.App.Log.Error(str)
		}
	}
	//判断错误类型
	if res := strings.Contains(msg, "token is expired by"); res { //token超时
		c.JSON(200, Response{
			401,
			http.StatusInternalServerError,
			nil,
			msg,
		})
	} else if res := strings.Contains(msg, "invalid memory address or nil pointer dereference"); res { //数据库链接失败
		database.InitDatabase() //重连数据库-初始化数据
		c.JSON(http.StatusInternalServerError, Response{1,
			http.StatusInternalServerError,
			"可能是数据库链接失败，请查看数据库链接是否正常",
			msg + "，可能是数据库链接失败，请查看数据库配置及是否启动，再刷新试试！",
		})
	} else {
		c.JSON(http.StatusInternalServerError, Response{1,
			http.StatusInternalServerError,
			nil,
			msg,
		})
	}
	c.Abort()
}
