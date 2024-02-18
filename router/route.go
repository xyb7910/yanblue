package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yanblue/controller"
	"yanblue/logger"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册路由
	r.POST("/api/v1/login", controller.SignUpHandler)

	// 注册登陆
	r.POST("/api/v1/login", controller.LoginHandler)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
