package system

import (
	"gomap/model/system"
	"gomap/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")

	{
		baseRouter.POST("login", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "login")
		})

		baseRouter.POST("register", func(ctx *gin.Context) {
			var form system.Register
			if err := ctx.ShouldBindJSON(&form); err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"error": utils.GetErrorMsg(form, err),
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		})
	}
	return baseRouter
}
