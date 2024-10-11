package router

import (
	"fisco-cert-app/controller"
	"fisco-cert-app/middlewares"

	"github.com/gin-gonic/gin"
)

func UserGroup(r *gin.RouterGroup) {
	r.POST("/signup", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)
	r.GET("/getUserInfo", middlewares.JwtAuthMiddleware(), controller.UserInfoHandler)
	r.GET("/refresh", middlewares.JwtAuthByRefreshMiddleware(), controller.RefreshToken)
	r.GET("/getUserSum", controller.GetUserCountHandler)
}
