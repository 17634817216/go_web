package router

import (
	"context"
	"fmt"
	"go_admin_api/global"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunServer() {
	//加载路由
	r := InitRouter()
	srv := &http.Server{
		Addr:    ":" + global.App.AppConfig.PORT,
		Handler: r,
	}

	global.App.Log.Info("启动端口：" + global.App.AppConfig.PORT)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			str := fmt.Sprintf("listen: %s\n", err) //拼接字符串
			global.App.Log.Error(str)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.App.Log.Info("关闭服务器 ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		str := fmt.Sprintf("服务器关闭： %s\n", err) //拼接字符串
		global.App.Log.Error(str)
	}
	global.App.Log.Info("服务器正在退出 ...")
}
