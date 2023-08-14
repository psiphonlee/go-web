package initialize

import (
	"gomap/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	systemRouter := router.RouterGroupApp.System

	PublicGroup := Router.Group("")
	{
		// 健康检测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}

	{
		// 注册基础路由 不做鉴权
		systemRouter.InitBaseRouter(PublicGroup)
	}
	return Router
}
