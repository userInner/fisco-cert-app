package controller

import (
	"errors"
	"fisco-cert-app/dao/mysql"
	"fisco-cert-app/logic"
	"fisco-cert-app/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SignUpHandler 注册请求处理
// @Summary 用户注册
// @Description 用户注册接口
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer JWT"
// @Param object body models.ParamSignUp true "参数"
// @Success 200 {object} controller.ResponseData
// @Router /signup [post]
func SignUpHandler(c *gin.Context) {
	p := new(models.ParamSignUp)
	// 校验参数
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("signUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 实现注册逻辑
	hexPrivateKey, err := logic.SignUp(p)
	if err != nil {
		zap.L().Error("logic.signup failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, hexPrivateKey)
}

// LoginHandler 登陆
// @Summary 用户登陆
// @Description 用户登陆接口
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer JWT"
// @Param object body models.ParamLogin true "查询参数"
// @Success 200 {object} controller.ResponseData
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	p := new(models.ParamLogin)
	// 校验参数
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invailed param", zap.Error(err))
		errs, ok := err.(*validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 实现登陆逻辑
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.login failed", zap.Error(err))
		// 密码错误
		if errors.Is(err, mysql.ErrorInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, gin.H{
		"user_id":       user.UserID,
		"user_name":     user.UserName,
		"access_token":  user.AccessToken,
		"refresh_token": user.RefreshToken,
	})
}

// UserInfoHandler 获取用户详情信息
// @Summary 获取用户详情信息
// @Description 通过Header中Authorization查询用户信息
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer JWT"
// @Success 200 {object} controller.ResponseData
// @Router /getUserInfo [post]
func UserInfoHandler(c *gin.Context) {
	uid, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeInvalidToken)
		return
	}
	user, err := logic.GetUserInfo(uid)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	ResponseSuccess(c, gin.H{
		"user_id":   user.UserID,
		"user_name": user.UserName,
		"address":   user.Address,
		"cert_sum":  user.CertSum,
	})
}

// RefreshToken 获取新的AccessToken
// @Summary 获取新AccessToken
// @Description 用户获取新的AccessToken接口
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer Access_Token"
// @Param RefreshToken header string true "Bearer Refresh_Token"
// @Success 200 {object} string "access_token"
// @Router /refresh [get]
func RefreshToken(c *gin.Context) {
	AToken, err := getAccessToken(c)
	if err != nil {
		ResponseError(c, CodeInvalidToken)
		return
	}
	ResponseSuccess(c, gin.H{
		"access_token": AToken,
	})
}
