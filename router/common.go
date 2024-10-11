package router

import (
	"fisco-cert-app/controller"
	"fisco-cert-app/middlewares"
	"github.com/gin-gonic/gin"
)

func CommonGroup(r *gin.RouterGroup) {
	// 水印， need login
	r.POST("/watermark", middlewares.JwtAuthMiddleware(), controller.WaterMarkHandler)
	// 图片水印添加
	r.POST("/img", middlewares.JwtAuthMiddleware(), controller.WaterMarkImgHandler)
	// 鉴权
	r.POST("/auth", middlewares.JwtAuthMiddleware(), controller.AuthHandler)

	r.POST("/imgAuth", middlewares.JwtAuthMiddleware(), controller.ImgAuthHandler)
	// 鉴权列表
	r.GET("getUserCount", controller.GetUserCountHandler)
	// 已添加的水印版权列表
	r.GET("/list", controller.GetEvidenceListHandler)

	// 版权详情
	r.GET("/evidence", controller.GetEvidenceDetailHandler)

	// 上传水印
	r.POST("/upwatermark", middlewares.JwtAuthMiddleware(), controller.UpWaterMarkHandler)
}
