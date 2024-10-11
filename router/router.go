package router

import (
	"fisco-cert-app/logger"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()

	// 注册自己的中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(false))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})
	v1 := r.Group("/api/v1")
	v1.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 用户api
	UserGroup(v1)
	// 水印相关
	CommonGroup(v1)
	// 区块链相关
	BlockChainRouter(v1)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
