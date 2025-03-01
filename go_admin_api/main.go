package main

import (
	"go_admin_api/global"
	"go_admin_api/router"
	"go_admin_api/utils"
	"runtime"
	"strconv"
)

func main() {

	//初始化
	global.InifConfig()

	//初始化日志
	global.App.Log = utils.InitializeLog()

	cpu_num, _ := strconv.Atoi(global.App.AppConfig.CPUNUM)
	mycpu := runtime.NumCPU()
	if cpu_num > mycpu { //如果配置cpu核数大于当前计算机核数，则等当前计算机核数
		cpu_num = mycpu
	}
	if cpu_num > 0 {
		runtime.GOMAXPROCS(cpu_num)
	} else {
		runtime.GOMAXPROCS(mycpu)
	}

	////端口以及服务
	router.RunServer()

}
