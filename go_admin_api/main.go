package main

import (
	"go_admin_api/global"
	"go_admin_api/router"
	"go_admin_api/utils"
)

func main() {
	//初始化文件
	global.InifConfig()
	//初始化日志
	global.App.Log = utils.InitializeLog()
	//端口以及服务
	router.RunServer()

}
