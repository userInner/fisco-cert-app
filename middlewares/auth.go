package middlewares

import (
	"fisco-cert-app/controller"
	"fisco-cert-app/pkg/jwt"
	"strings"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// JwtAuthMiddleware 基于JWT认证中间件
func JwtAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}
		// Authorization: Bearer fdjfjjjjj.....
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidToken)
			zap.L().Error("Authoriztion failed,err:", zap.Error(err))
			c.Abort()
			return
		}
		c.Set(controller.CtxUserIDKey, mc.UserID)
		c.Next()
	}
}

// JwtAuthByRefreshMiddleware 用于刷新AccessToken的中间件
func JwtAuthByRefreshMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		AToken := c.Request.Header.Get("Authorization")
		if AToken == "" {
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}
		// Authorization: Bearer fdjfjjjjj.....
		parts := strings.SplitN(AToken, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		RToken := c.Request.Header.Get("RefreshToken")
		if RToken == "" {
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}
		AToken, RToken, err := jwt.ParseRefreshToken(parts[1], RToken)
		if err != nil {
			zap.L().Error("refresh failed", zap.Error(err))
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}
		c.Set(controller.CtxUserAccessTokenKey, AToken)
		c.Next()
	}
}
