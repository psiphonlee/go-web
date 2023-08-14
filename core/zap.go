package core

import (
	"fmt"
	"gomap/core/internal"
	"gomap/global"
	"gomap/utils"
	"os"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// initZap 获取 zap.Logger 实例指针
func InitZap() (logger *zap.Logger) {
	// 检查存放日志的文件夹是否存在
	if ok, _ := utils.PathExists(global.EWA_CONFIG.Zap.LogDir); !ok {
		fmt.Printf("创建日志文件夹:%s\n", global.EWA_CONFIG.Zap.LogDir)
		os.Mkdir(global.EWA_CONFIG.Zap.LogDir, os.ModePerm)
	}
	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.EWA_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	fmt.Println("====2-zap====: zap log init success")
	return logger
}
