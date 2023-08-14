package core

import (
	"fmt"
	"gomap/global"
	"gomap/initialize"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	// 初始化redis服务
	initialize.Redis()

	// 初始化路由
	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.EWA_CONFIG.App.Port)
	s := initServer(address, Router)

	global.EWA_LOG.Info("server run success on:", zap.String("address", address))

	// 保证文本顺序输出
	time.Sleep(10 * time.Microsecond)

	fmt.Println(`address`, address)

	global.EWA_LOG.Error(s.ListenAndServe().Error())
}

func initServer(address string, router *gin.Engine) server {
	// 使用endless库创建一个HTTP服务器，监听地址为address，router是HTTP请求路由器
	s := endless.NewServer(address, router)

	// 设置HTTP请求头的读取超时时间为20秒，如果在20秒内未读取到请求头，会返回一个超时错误
	s.ReadHeaderTimeout = 20 * time.Second

	// 设置HTTP响应体的写入超时时间为20秒，如果在20秒内未将响应体写入完成，会返回一个超时错误
	s.WriteTimeout = 20 * time.Second

	// 设置HTTP请求头的最大字节数为1MB，如果超过1MB，会返回一个错误
	s.MaxHeaderBytes = 1 << 20

	return s
}
