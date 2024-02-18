package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yanblue/controller"
	"yanblue/logger"
	"yanblue/middlewares"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")
	// 注册路由
	v1.POST("/signup", controller.SignUpHandler)

	// 注册登陆
	v1.POST("/login", controller.LoginHandler)

	v1.GET("/community", controller.CommunityHandler)

	v1.GET("/community/:id", controller.CommunityDetailHandler)

	v1.GET("/post/:id", controller.GetPostDetailHandler)

	v1.GET("/posts", controller.GetPostListHandler)

	v1.Use(middlewares.JWTAuthMiddleware()) // jwt auth

	{
		v1.POST("/post", controller.CreatePostHandler)

	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
