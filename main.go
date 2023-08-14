package main

import (
	"gomap/core"
	"gomap/global"
	"gomap/initialize"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const AppMode = "debug" // 运行环境，主要有三种：debug、test、release

func main() {
	// 通过环境变量设置读取哪个配置文件：config.yaml,
	gin.SetMode(AppMode)

	// TODO：1.配置初始化
	global.EWA_VIPER = core.InitViper()

	// TODO：2.日志组件初始化
	global.EWA_LOG = core.InitZap()
	zap.ReplaceGlobals(global.EWA_LOG)
	global.EWA_LOG.Info("zap init sucess", zap.String("zap_log", "zap_log"))

	// TODO：3.数据库连接
	global.EWA_DB = initialize.Gorm()

	// TODO：4.其他初始化
	initialize.OtherInit()
	// TODO：5.启动服务

	// 启动服务器
	core.RunServer()
}
